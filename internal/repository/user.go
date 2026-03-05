package repository

import (
	"Into/internal/models"
	"context"
	"database/sql"

)

type UserRepository interface {
	CreateUser(ctx context.Context, req *models.User) (*models.User, error)
	CreateUserProfile(ctx context.Context, req *models.UserProfile) (*models.UserProfile, error)
	UpdateUserProfile(ctx context.Context, req *models.UserProfile) (*models.UserProfile, error)
	DeleteUser(ctx context.Context, ID uint) error
	GetByID(ctx context.Context, ID uint) (models.User, error)
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

func (r *userRepository) CreateUserProfile(ctx context.Context, req *models.UserProfile) (*models.UserProfile, error) {
	query := `
	INSERT INTO user_profiles (user_id, first_name, last_name, phone, birth_date)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, user_id, first_name, last_name, phone, birth_date
	`

	var profile models.UserProfile

	err := r.db.QueryRowContext(ctx,
		query,
		req.UserID,
		req.FirstName,
		req.LastName,
		req.Phone,
		req.BirthDate).
		Scan(
			&profile.ID,
			&profile.UserID,
			&profile.FirstName,
			&profile.LastName,
			&profile.Phone,
			&profile.BirthDate,
		)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *userRepository) UpdateUserProfile(ctx context.Context, req *models.UserProfile) (*models.UserProfile, error) {
	query := `
	UPDATE user_profiles
	SET first_name = $1,
	    last_name = $2,
	    phone = $3,
	    birth_date = $4,
	    updated_at = NOW()
	WHERE user_id = $5
	RETURNING id, user_id, first_name, last_name, phone, birth_date
	`

	var profile models.UserProfile

	err := r.db.QueryRowContext(ctx,
		query,
		req.FirstName,
		req.LastName,
		req.Phone,
		req.BirthDate,
		req.UserID).
		Scan(
			&profile.ID,
			&profile.UserID,
			&profile.FirstName,
			&profile.LastName,
			&profile.Phone,
			&profile.BirthDate,
		)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, ID uint) error{

	_, err := r.db.ExecContext(ctx, `DELETE FROM users WHERE id = $1`, ID)

	return err
}


func (r *userRepository) GetByID(ctx context.Context, ID uint) (models.User, error) {

	var user models.User

	err := r.db.QueryRowContext(ctx,`
	SELECT username, email
	FROM users
	WHERE id = $1`, ID).
	Scan(
		&user.Username,
		&user.Email,
	)

	return user, err
	 
}