package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ishanshre/GoRestApiExample/internals/handler"
)

func SetupRouter(handler handler.VideoHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/videos", handler.FindAll)
	r.POST("/videos/create", handler.Save)
	return r
}
