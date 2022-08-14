package utils

import "github.com/gin-gonic/gin"

func GinResponseWrapper(ctx *gin.Context, data interface{}, code int) {
	var res gin.H
	if code == 0 {
		code = 400
	}

	if data != nil {
		res = gin.H{
			"data": data,
		}
	}
	ctx.JSON(code, res)
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
		"meta": gin.H{
			"message": message,
		},
	})
}
