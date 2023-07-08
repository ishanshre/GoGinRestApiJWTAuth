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
	r.GET("/videos", handler.FindAll)
	r.POST("/videos/create", handler.Save)
	return r
}
