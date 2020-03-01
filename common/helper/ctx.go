package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phoon/go-forum/repository/model"
)

//CtxAbort abort the context, you should use the return statement immediately after.
func CtxAbort(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, model.Response{
		Code: code,
		Msg:  msg,
	})
	ctx.Abort()
}

//CtxJSON for normal output.
func CtxJSON(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, model.Response{
		Code: http.StatusOK,
		Msg:  msg,
		Data: data,
	})
}
