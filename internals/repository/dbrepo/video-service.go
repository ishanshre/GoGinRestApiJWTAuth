package dbrepo

import (
	"github.com/ishanshre/GoRestApiExample/internals/models"
)

func (s *videoService) Save(video models.Video) (models.Video, error) {
	s.videos = append(s.videos, video)
	return video, nil
}

func (s *videoService) FindAll() ([]models.Video, error) {
	return s.videos, nil
}
