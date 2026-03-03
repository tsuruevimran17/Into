package repository

import (
	"Into/internal/models"
	"context"
	"database/sql"
)

type UserRepository interface {
	CreateUser(ctx context.Context, req *models.User) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, req *models.User) (*models.User, error) {
	query := `
	INSERT INTO users (username, email, role, password_hash)
	VALUES ($1, $2, $3, $4)
	RETURNING id, username, email, role
	`

	var user models.User

	err := r.db.QueryRowContext(ctx,
	query,
	req.Username,
	req.Email,
	req.Role,
	req.PasswordHash).
	Scan(&user.ID,
		&user.Username,
		&user.Email,
		&user.Role)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
