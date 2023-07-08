package main

import (
	"log"

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
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Error in loading env files")
	}
	videoService := dbrepo.NewVideoService()
	handler := handler.NewRepo(videoService)
	return handler
}
