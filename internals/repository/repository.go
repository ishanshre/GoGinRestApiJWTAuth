package repository

import "github.com/ishanshre/GoRestApiExample/internals/models"

type DatabaseRepo interface {

	// Video interface
	CreateVideo(*models.Video) (*models.Video, error)
	GetAllVideos() ([]*models.Video, error)
	GetVideoByID(id int) (*models.Video, error)
	DeleteVideoByID(id int) error

	// author interface
	CreateAuthor(a models.Author) (*models.Author, error)
	GetAllAuthors() ([]*models.Author, error)
	GetAuhtorByID(id int) (*models.Author, error)
	DeleteAuthorByID(id int) error

	//user interface
	CreateUser(u *models.CreateUser) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	UserExists(username string) (bool, error)
	EmailExists(email string) (bool, error)
}
