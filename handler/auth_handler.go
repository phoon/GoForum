package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/phoon/go-forum/common/helper"
	"github.com/phoon/go-forum/repository"
	"github.com/phoon/go-forum/repository/model"
)

//SignIn use user's account and password for certification.
//A account can be a username or email address.
func SignIn(ctx *gin.Context) {
	sess := sessions.Default(ctx)

	//get and verify the form values
	var loginForm model.UserLoginInput
	if err := ctx.ShouldBind(&loginForm); err != nil {
		helper.CtxAbort(ctx, http.StatusUnauthorized, "parameters not met: "+err.Error())
		return
	}

	user := &model.User{}
	//get user in different way by check the account type
	if helper.IsEmailValid(loginForm.Account) {
		user = repository.UserRepo.FindByEmail(loginForm.Account)
	} else {
		user = repository.UserRepo.FindByUserName(loginForm.Account)
	}
	if user == nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "invalid username")
		return
	}

	//the user who has been banned for some reasons can not allowed to sign in.
	if user.Status == model.UserStatusBanned {
		helper.CtxAbort(ctx, http.StatusForbidden, "sorry, your account has been banned.")
		return
	}

	if !helper.IsUserPassValid(user, loginForm.Password) {
		helper.CtxAbort(ctx, http.StatusBadRequest, "password verification failed")
		return
	}
	//authentication success, store information to the session
	sess.Set("login_info", model.LoginInfo{
		UserID:        user.ID,
		UserName:      user.Name,
		Role:          user.Role,
		Authenticated: true,
	})

	if err := sess.Save(); err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helper.CtxJSON(ctx, "signin success", nil)
}

//SignUp will create a new user, and user's password should always be bcrypted before it got stored.
func SignUp(ctx *gin.Context) {
	var signupForm model.User
	if err := ctx.ShouldBind(&signupForm); err != nil {
		helper.CtxAbort(ctx, http.StatusBadRequest, "parameters not met: "+err.Error())
		return
	}

	//create user
	//use bcrypt to process the password
	signupForm.Password = helper.GenHashedPass(signupForm.Password)
	if err := repository.UserRepo.Add(&signupForm); err != nil {
		//registeration fail
		helper.CtxAbort(ctx, http.StatusInternalServerError, "fail to register: "+err.Error())
		return
	}

	helper.CtxJSON(ctx, "register success", nil)
}

//SignOut will delete the session information of the user.
func SignOut(ctx *gin.Context) {
	sess := sessions.Default(ctx)

	sess.Clear()
	sess.Options(sessions.Options{
		Path:   "/",
		MaxAge: 0,
	})

	if err := sess.Save(); err != nil {
		helper.CtxAbort(ctx, http.StatusInternalServerError, "logout failed: "+err.Error())
	}

	helper.CtxJSON(ctx, "sigin out successfully, see you next time", nil)
}
