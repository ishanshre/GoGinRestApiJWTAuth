package repository

import "github.com/ishanshre/GoRestApiExample/internals/models"

type VideoService interface {
	Save(models.Video) (models.Video, error)
	FindAll() ([]models.Video, error)
}
