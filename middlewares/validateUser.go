package middlewares

import (
	"net/http"

	"github.com/dileofrancoj/blog-app/utils"
	"github.com/gin-gonic/gin"
)

func ValidateJWT () gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		
		_,_,_, err := utils.ValidateJWT(header)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		} else {
			ctx.Next()
		}
		

	}
}