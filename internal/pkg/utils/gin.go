package utils

import "github.com/gin-gonic/gin"

func GinResponseWrapper(ctx *gin.Context, data interface{}, code int) {
	if code == 0 {
		code = 400
	}
	ctx.JSON(code, data)
}

func GinErrorWrapper(ctx *gin.Context, err interface{}, code int) {
	if code == 0 {
		code = 400
	}

	message := ""
	switch e := err.(type) {
	case error:
		message = e.Error()
	}

	ctx.JSON(code, gin.H{
		"message": message,
	})
}
