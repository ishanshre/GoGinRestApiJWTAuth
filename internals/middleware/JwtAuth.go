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
		tokenClaims, err := helper.VerifyTokenWithClaims(tokenString[1], "access_token")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
		if err := redisClient.Exists(ctx, tokenClaims.TokenID).Err(); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid token or token does not exists in cache",
			})
			ctx.Abort()
			return
		}
		log.Println(redisClient.Get(ctx, tokenClaims.TokenID))
		ctx.Set("tokenID", tokenClaims.TokenID)
		ctx.Set("userID", tokenClaims.UserID)
		ctx.Set("username", tokenClaims.Username)
		ctx.Next()
	}
}
