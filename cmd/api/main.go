package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ishanshre/GoRestApiExample/internals/handler"
	"github.com/ishanshre/GoRestApiExample/internals/repository/dbrepo"
	"github.com/joho/godotenv"
)

const port = ":8000"

func main() {
	handler := run()
	r := gin.Default()
	r.GET("/videos", handler.FindAll)
	r.POST("/videos/create", handler.Save)
	r.Run(port)
}

func run() handler.VideoHandler {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Error in loading env files")
	}
	videoService := dbrepo.NewVideoService()
	handler := handler.NewRepo(videoService)
	return handler
}
