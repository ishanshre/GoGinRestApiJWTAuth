package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ishanshre/GoRestApiExample/internals/handler"
	"github.com/ishanshre/GoRestApiExample/internals/repository/dbrepo"
	"github.com/ishanshre/GoRestApiExample/internals/router"
	"github.com/joho/godotenv"
)

const port = ":8000"

func main() {
	handler := run()

	r := router.SetupRouter(handler)

	r.Run(port)
}

func run() handler.VideoHandler {

	// Loading environment files
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Error in loading env files")
	}
	// setup log file
	setupLoggerOutput()

	// connecting to database repository
	videoService := dbrepo.NewVideoService()

	// connecting the handler
	handler := handler.NewRepo(videoService)
	return handler
}

func setupLoggerOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
