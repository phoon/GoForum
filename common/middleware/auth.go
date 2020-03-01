package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phoon/go-forum/common/helper"
	"github.com/phoon/go-forum/repository/model"
)

//LoginNeed checks the session.
func LoginNeed() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if login, _ := helper.GetLoginInfoFromSession(ctx); login == nil {
			helper.CtxAbort(ctx, http.StatusUnauthorized, "please login first")
			return
		}

		ctx.Next()
	}
}

//Admin checks whether the request is authenticated as super user.
//!!!This middleware is used for super priviledge only routers.
//Others like `DeleteUser` should be process within itself.
func Admin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loginInfo, _ := helper.GetLoginInfoFromSession(ctx)
		if loginInfo == nil {
			helper.CtxAbort(ctx, http.StatusUnauthorized, "please login first")
			return
		}

		if loginInfo.Role != model.UserRoleAdmin {
			helper.CtxAbort(ctx, http.StatusUnauthorized, "only super user can do this")
			return
		}

		ctx.Next()
	}
}
