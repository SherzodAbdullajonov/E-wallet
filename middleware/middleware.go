package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		client := ctx.Request.Header.Get("user")
		if client == ""{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintln("No Authorization header provided")})
			ctx.Abort()
			return
		}
	}
}