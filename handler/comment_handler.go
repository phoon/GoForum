package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phoon/go-forum/common/helper"
	"github.com/phoon/go-forum/repository"
	"github.com/phoon/go-forum/repository/model"
)

//GetAllComments gets all the comments, no need to generate fields for output.
func GetAllComments(ctx *gin.Context) {
	limit, offset, err := helper.CalLimitOffset(ctx)
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	comments := repository.CommentRepo.GetAll(limit, offset)
	helper.CtxJSON(ctx, "", comments)
}

//GetCommentByID get the comment by specified id.
func GetCommentByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	comment := repository.CommentRepo.FindByID(uint(id))
	if comment == nil {
		helper.CtxAbort(ctx, http.StatusNotFound, "comment does not exists")
		return
	}

	helper.CtxJSON(ctx, "", comment)
}

//CreateComment create a new comment.user is should be obtained from the `claims`.
func CreateComment(ctx *gin.Context) {
	var createCommentForm model.Comment
	if err := ctx.ShouldBind(&createCommentForm); err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "parameters not met: "+err.Error())
		return
	}

	//get the login information
	loginInfo, _ := helper.GetLoginInfoFromSession(ctx)

	//set comment owner
	createCommentForm.UserID = loginInfo.UserID
	if err := repository.CommentRepo.Add(&createCommentForm); err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, "fail to create the comment: "+err.Error())
		return
	}

	helper.CtxJSON(ctx, "create the comment successfully", createCommentForm)
}

//UpdateComment update a comment, only the owner or super user can do this.
func UpdateComment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	loginInfo, _ := helper.GetLoginInfoFromSession(ctx)
	//get the comment
	comment := repository.CommentRepo.FindByID(uint(id))
	if comment == nil {
		helper.CtxAbort(ctx, http.StatusNotFound, "comment does not exists")
		return
	}

	//check current user's permission
	if loginInfo.UserID != comment.UserID && loginInfo.Role != model.UserRoleAdmin {
		helper.CtxAbort(ctx, http.StatusUnauthorized, "you don's have permission to do that")
		return
	}

	//check the input form
	var updateCommentFrom model.Comment
	if err = ctx.ShouldBind(&updateCommentFrom); err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "parameters not met: "+err.Error())
		return
	}

	//set owner
	updateCommentFrom.UserID = loginInfo.UserID
	if err = repository.CommentRepo.Update(&updateCommentFrom); err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, "fail to update the comment: "+err.Error())
		return
	}

	helper.CtxJSON(ctx, "update the comment successfully", updateCommentFrom)
}

//DeleteComment deletes the comment by specified id.
//Only performance with comment owner or super user's permission.
func DeleteComment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	loginInfo, _ := helper.GetLoginInfoFromSession(ctx)
	//get the given comment
	comment := repository.CommentRepo.FindByID(uint(id))
	if comment == nil {
		helper.CtxAbort(ctx, http.StatusNotFound, "comment does not exists")
		return
	}

	//check the permission
	if comment.UserID != loginInfo.UserID && loginInfo.Role != model.UserRoleAdmin {
		helper.CtxAbort(ctx, http.StatusUnauthorized, "you don't have permission to do that")
		return
	}

	//do the delete
	if err := repository.CommentRepo.Delete(uint(id)); err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, "fail to delete the comment: "+err.Error())
		return
	}

	helper.CtxJSON(ctx, "comment was removed successfully", nil)
}
