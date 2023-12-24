package dbrepo

import (
	"context"

	"github.com/ishanshre/GoRestApiExample/internals/models"
)

func (s *postgresDBRepo) CreateAuthor(a models.Author) (*models.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	stmt := `
		INSERT INTO authors (first_name, last_name, email)
		VALUES ($1,$2,$3) RETURNING id, first_name, last_name, email
	`
	row := s.DB.QueryRowContext(
		ctx,
		stmt,
		a.FirstName,
		a.LastName,
		a.Email,
	)
	author := &models.Author{}
	if err := row.Scan(
		&author.ID,
		&author.FirstName,
		&author.LastName,
		&author.Email,
	); err != nil {
		return nil, err
	}
	return author, nil
}

func (s *postgresDBRepo) GetAllAuthors() ([]*models.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	query := `SELECT * FROM authors`
	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	authors := []*models.Author{}
	for rows.Next() {
		a := &models.Author{}
		if err := rows.Scan(
			&a.ID,
			&a.FirstName,
			&a.LastName,
			&a.Email,
		); err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}
	return authors, err
}

func (s *postgresDBRepo) GetAuhtorByID(id int) (*models.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	query := `SELECT * FROM authors WHERE id=$1`
	row := s.DB.QueryRowContext(ctx, query, id)
	a := &models.Author{}
	if err := row.Scan(
		&a.ID,
		&a.FirstName,
		&a.LastName,
		&a.Email,
	); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *postgresDBRepo) DeleteAuthorByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	stmt := `DELETE FROM authors WHERE id=$1`
	_, err := s.DB.ExecContext(ctx, stmt, id)
	return err
}
