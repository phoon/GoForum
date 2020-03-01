package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/phoon/go-forum/common/helper"
	"github.com/phoon/go-forum/repository"
	"github.com/phoon/go-forum/repository/model"
)

//GetAllUsers gets the all users, paging is possible
func GetAllUsers(ctx *gin.Context) {
	limit, offset, err := helper.CalLimitOffset(ctx)
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}
	users := repository.UserRepo.GetAll(limit, offset)
	for _, user := range users {
		helper.GenUserRelatedFields(user)
	}

	helper.CtxJSON(ctx, "", users)
}

//GetUserByName output the user info by specified username
func GetUserByName(ctx *gin.Context) {
	name := ctx.Param("name")
	var user *model.User
	if user = repository.UserRepo.FindByUserName(name); user == nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid username")
		return
	}

	helper.GenUserRelatedFields(user)
	helper.CtxJSON(ctx, "", user)
}

//DeleteUser delete user by user self or super user
func DeleteUser(ctx *gin.Context) {
	name := ctx.Param("name")

	loginInfo, _ := helper.GetLoginInfoFromSession(ctx)

	//only user self or super user can do that
	if loginInfo.UserName != name && loginInfo.Role != model.UserRoleAdmin {
		helper.CtxAbort(ctx, http.StatusUnauthorized, "you don't have permission to do that")
		return
	}

	user := repository.UserRepo.FindByUserName(name)
	if user == nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, "invalid username")
		return
	}

	if err := repository.UserRepo.Delete(user.ID); err != nil {
		//fail to delete
		helper.CtxAbort(ctx, http.StatusInternalServerError, "fail to delete user ["+name+"]: "+err.Error())
		return
	}

	helper.CtxJSON(ctx, "delete user ["+name+"] success", nil)
}

//FollowOthers is become other's followers, or not to be.
func FollowOthers(ctx *gin.Context) {
	name := ctx.Param("name")

	loginInfo, _ := helper.GetLoginInfoFromSession(ctx)
	user := repository.UserRepo.FindByUserName(name)
	if user == nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid username")
		return
	}

	if loginInfo.UserName == name {
		//follow yourself??? come on bro
		helper.CtxAbort(ctx, http.StatusBadRequest, "follow yourself is not allowed")
		return
	}
	//processing for follow
	userSelf := repository.UserRepo.FindByID(loginInfo.UserID)
	err, follow := repository.UserRepo.FollowOthers(userSelf, user)
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "fail to follow user["+user.Name+"]: "+err.Error())
		return
	}

	if !follow {
		helper.CtxJSON(ctx, "unfollow success", nil)
		return
	}

	helper.CtxJSON(ctx, "follow success", nil)
}

//GetUserFollowers is get the user's associated followers
func GetUserFollowers(ctx *gin.Context) {
	name := ctx.Param("name")
	limit, offset, err := helper.CalLimitOffset(ctx)
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := repository.UserRepo.FindByUserName(name)
	if user == nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid username")
		return
	}

	followers := repository.UserRepo.GetUserFollowers(user, limit, offset)
	for _, follower := range followers {
		helper.GenUserRelatedFields(follower)
	}

	helper.CtxJSON(ctx, "", followers)
}

//GetUserFollowing is get the user's associated user that he follow with
func GetUserFollowing(ctx *gin.Context) {
	name := ctx.Param("name")
	limit, offset, err := helper.CalLimitOffset(ctx)
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := repository.UserRepo.FindByUserName(name)
	if user == nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid username")
		return
	}

	followings := repository.UserRepo.GetUserFollowing(user, limit, offset)
	for _, following := range followings {
		helper.GenUserRelatedFields(following)
	}

	helper.CtxJSON(ctx, "", followings)
}

//GetUserTopics finds the topics that a given user published
func GetUserTopics(ctx *gin.Context) {
	name := ctx.Param("name")
	limit, offset, err := helper.CalLimitOffset(ctx)
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := repository.UserRepo.FindByUserName(name)
	if user == nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid username")
		return
	}

	topics := repository.UserRepo.GetRelatedTopics(user, limit, offset)
	for _, topic := range topics {
		helper.GenTopicRelatedFields(topic)
	}

	helper.CtxJSON(ctx, "", topics)
}

//GetUserComments gets the comments that a given user left.
func GetUserComments(ctx *gin.Context) {
	name := ctx.Param("name")
	limit, offset, err := helper.CalLimitOffset(ctx)
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := repository.UserRepo.FindByUserName(name)
	if user == nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid username")
		return
	}

	comments := repository.UserRepo.GetRelatedComments(user, limit, offset)
	helper.CtxJSON(ctx, "", comments)
}

//UpdateUser updates the user's infomation.
//
//Only the users themselves and super user can do this.
//Update one's password will be ok with this function also?
func UpdateUser(ctx *gin.Context) {
	name := ctx.Param("name")

	loginInfo, _ := helper.GetLoginInfoFromSession(ctx)
	//check whether the user exists
	user := repository.UserRepo.FindByUserName(name)
	if user == nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "user does not exists, maybe you should register it first")
		return
	}

	//check current user's permission
	if loginInfo.UserName != name && loginInfo.Role != model.UserRoleAdmin {
		helper.CtxAbort(ctx, http.StatusUnauthorized, "you don't have permission to do that")
		return
	}

	//check the input form
	var userUpdateForm model.User
	if err := ctx.ShouldBind(&userUpdateForm); err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "parameters not met: "+err.Error())
		return
	}

	//avoid update password, it will cost many time
	//because it should be processed by bcrypt
	userUpdateForm.Password = ""
	userUpdateForm.ID = user.ID
	//update the user
	if err := repository.UserRepo.Update(&userUpdateForm); err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helper.CtxJSON(ctx, "update success", nil)
}

//UpdateUserPass update the user's password.
//Password should be bcrypted before it got stored.
//Only user self can do this.
func UpdateUserPass(ctx *gin.Context) {
	name := ctx.Param("name")
	password := ctx.PostForm("password")
	if len(strings.TrimSpace(password)) == 0 {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid password")
		return
	}
	loginInfo, _ := helper.GetLoginInfoFromSession(ctx)

	if loginInfo.UserName != name {
		helper.CtxAbort(ctx, http.StatusUnauthorized, "you don't have permission to do that")
		return
	}

	password = helper.GenHashedPass(password)

	if err := repository.UserRepo.UpdateColumn(loginInfo.UserID, "password", password); err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, "fail to update password: "+err.Error())
		return
	}

	helper.CtxJSON(ctx, "update password successfully", nil)
}

//UpdateUserEmail updates the user's email address.
func UpdateUserEmail(ctx *gin.Context) {
	name := ctx.Param("name")
	email := ctx.PostForm("email")
	if len(strings.TrimSpace(email)) == 0 || !helper.IsEmailValid(email) {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid email")
		return
	}
	loginInfo, _ := helper.GetLoginInfoFromSession(ctx)

	if loginInfo.UserName != name {
		helper.CtxAbort(ctx, http.StatusUnauthorized, "you don't have permission to do that")
		return
	}

	if err := repository.UserRepo.UpdateColumn(loginInfo.UserID, "email", email); err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, "fail to update email: "+err.Error())
		return
	}

	helper.CtxJSON(ctx, "update email successfully", nil)
}

//UpdateUserAvatar updates the user's avatar.
//Avatar should be a link, processed by front-end.
func UpdateUserAvatar(ctx *gin.Context) {
	name := ctx.Param("name")
	avatar := ctx.PostForm("avatar")
	loginInfo, _ := helper.GetLoginInfoFromSession(ctx)

	if loginInfo.UserName != name {
		helper.CtxAbort(ctx, http.StatusUnauthorized, "you don't have permission to do that")
		return
	}

	if err := repository.UserRepo.UpdateColumn(loginInfo.UserID, "avatar", avatar); err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, "fail to update avatar: "+err.Error())
		return
	}

	helper.CtxJSON(ctx, "update avatar successfully", nil)
}

//GetUserLikedTopics gets the topics that user liked, paging is support.
func GetUserLikedTopics(ctx *gin.Context) {
	name := ctx.Param("name")
	limit, offset, err := helper.CalLimitOffset(ctx)
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := repository.UserRepo.FindByUserName(name)
	if user == nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid username")
		return
	}

	likedTopics := repository.UserRepo.GetLikedTopics(user, limit, offset)
	for _, topic := range likedTopics {
		helper.GenTopicRelatedFields(topic)
	}

	helper.CtxJSON(ctx, "", likedTopics)
}

//UpdateUserStatus updates the status of user.
//Only super user can do this. Like to ban a user.
func UpdateUserStatus(ctx *gin.Context) {
	name := ctx.Param("name")
	status, err := strconv.Atoi(ctx.PostForm("status"))
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid status")
		return
	}

	user := repository.UserRepo.FindByUserName(name)
	if user == nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid user")
		return
	}

	if err := repository.UserRepo.UpdateColumn(user.ID, "status", status); err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, "fail to change the status: "+err.Error())
		return
	}

	helper.CtxJSON(ctx, "change the user's status successfully", nil)
}
