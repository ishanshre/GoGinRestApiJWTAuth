package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ishanshre/GoRestApiExample/internals/driver"
	"github.com/ishanshre/GoRestApiExample/internals/handler"
	"github.com/ishanshre/GoRestApiExample/internals/middleware"
	"github.com/ishanshre/GoRestApiExample/internals/repository/dbrepo"
	"github.com/ishanshre/GoRestApiExample/internals/router"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

const port = ":8000"

const redisHost = "127.0.0.1:6379"

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

	db, err := driver.NewDatabase("postgres", os.Getenv("postgres"))
	if err != nil {
		log.Printf("Error in connecting to database: %s\n", err)
	}
	// connecting to redis
	redisPool := redis.NewClient(
		&redis.Options{
			Addr:         redisHost,
			Password:     "",
			DB:           0,
			MaxIdleConns: 10,
		},
	)

	// connecting to database repository
	videoService := dbrepo.NewPostgresRepo(db)

	// connecting the handler
	handler := handler.NewRepo(videoService, redisPool)
	middleware.NewClient(redisPool)
	return handler
}

func setupLoggerOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
