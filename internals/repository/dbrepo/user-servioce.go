package dbrepo

import (
	"context"
	"fmt"

	"github.com/ishanshre/GoRestApiExample/internals/models"
)

func (s *postgresDBRepo) CreateUser(u *models.CreateUser) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	stmt := `
		INSERT INTO users (first_name, last_name, username, email, password, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, first_name, last_name, username, email, password, created_at
	`
	row := s.DB.QueryRowContext(
		ctx,
		stmt,
		u.FirstName,
		u.LastName,
		u.Username,
		u.Email,
		u.Password,
		u.CreatedAt,
	)
	user := &models.User{}
	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *postgresDBRepo) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	query := `SELECT * FROM users WHERE id=$1`
	row := s.DB.QueryRowContext(ctx, query, id)
	u := &models.User{}
	if err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Username,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
	); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *postgresDBRepo) UserExists(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	query := `
		SELECT COUNT(*) FROM users
		WHERE username=$1
	`
	var count int
	if err := s.DB.QueryRowContext(ctx, query, username).Scan(&count); err != nil {
		return false, fmt.Errorf("failed to execute query: %w", err)
	}
	return count > 0, nil
}

func (s *postgresDBRepo) EmailExists(email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	query := `
		SELECT COUNT(*) FROM users 
		WHERE email = $1
	`
	var count int

	if err := s.DB.QueryRowContext(ctx, query, email).Scan(&count); err != nil {
		return false, fmt.Errorf("failed to execute query: %w", err)
	}
	return count > 0, nil
}
