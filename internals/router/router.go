package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ishanshre/GoRestApiExample/internals/handler"
	"github.com/ishanshre/GoRestApiExample/internals/middleware"
)

// setup gin and router
func SetupRouter(handler handler.VideoHandler) *gin.Engine {
	r := gin.New()

	// using middleware
	r.Use(gin.Recovery(), middleware.Logger())
	// r.Use(middleware.Basic_Auth())

	// In gin most specific router must be above
	r.DELETE("/videos/detail/:id/delete", handler.DeleteVideoByID)
	r.GET("/videos/detail/:id", handler.GetVideoByID)
	r.POST("/videos/create", middleware.JwtAccessAuthMiddleware(), handler.CreateVideo)
	r.GET("/videos", middleware.JwtAccessAuthMiddleware(), handler.GetAllVideos)

	// user router

	r.POST("/users/sign-up", handler.RegisterUser)
	r.POST("/users/login", handler.UserLogin)
	r.POST("/users/refresh", middleware.JwtAccessAuthMiddleware(), handler.RefreshToken)
	return r
}
