package models

type Role string

const (
	RoleUser   Role = "user"
	RoleSeller Role = "seller"
	RoleAdmin  Role = "admin"
)

type User struct {
	Base
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Role         Role   `json:"role"`
}
