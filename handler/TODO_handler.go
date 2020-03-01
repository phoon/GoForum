package handler

import (
	"github.com/gin-gonic/gin"
)

func TODO(ctx *gin.Context) {
	ctx.JSON(200, "a todo feature that unimplemented")
}
