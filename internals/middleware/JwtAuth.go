package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ishanshre/GoRestApiExample/internals/helper"
)

func JwtAccessAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// get the token string
		bearerToken := ctx.GetHeader("Authorization")
		if bearerToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			ctx.Abort()
			return
		}
		tokenString := strings.Split(bearerToken, " ")
		if len(tokenString) != 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token format",
			})
			ctx.Abort()
			return
		}
		if tokenString[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token format",
			})
			ctx.Abort()
			return
		}
		// verify and parse the JWT token
		token, err := helper.VerifyTokenWithClaims(tokenString[1], "access_token")
		if err != nil {
			log.Println("I am heree")
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
		log.Println("After token verification")
		ctx.Set("token", token)

		ctx.Next()
	}
}
