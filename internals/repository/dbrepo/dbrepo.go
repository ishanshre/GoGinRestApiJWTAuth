package dbrepo

import (
	"github.com/ishanshre/GoRestApiExample/internals/models"
	"github.com/ishanshre/GoRestApiExample/internals/repository"
)

type videoService struct {
	videos []models.Video
}

func NewVideoService() repository.VideoService {
	return &videoService{}
}
