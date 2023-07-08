package dbrepo

import (
	"context"
	"time"

	"github.com/ishanshre/GoRestApiExample/internals/models"
)

const timeout = 3 * time.Second

func (s *postgresDBRepo) Save(video *models.Video) (*models.Video, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	statement := `INSERT INTO videos (title, description, url, author_id) values
		($1,$2,$3,$4) RETURNING id, title, description, url, author_id
	`
	row := s.DB.QueryRowContext(
		ctx,
		statement,
		video.Title,
		video.Description,
		video.URL,
		video.Author.ID,
	)
	var v *models.Video
	err := row.Scan(
		&v.ID,
		&v.Title,
		&v.Description,
		&v.URL,
		&v.Author.ID,
	)

	return v, err
}

func (s *postgresDBRepo) FindAll() ([]*models.Video, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	query := `SELECT * FROM videos`
	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	videos := []*models.Video{}
	for rows.Next() {
		video := new(models.Video)
		if err := rows.Scan(
			&video.ID,
			&video.Title,
			&video.Description,
			&video.URL,
			&video.Author.ID,
		); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func (s *postgresDBRepo) FindOneWithID(id int) (*models.Video, error) {
	return nil, nil
}
