package repository

import "github.com/ishanshre/GoRestApiExample/internals/models"

type DatabaseRepo interface {
	Save(*models.Video) (*models.Video, error)
	FindAll() ([]*models.Video, error)
	FindOneWithID(id int) (*models.Video, error)
}
