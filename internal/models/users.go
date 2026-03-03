package models

import "time"

type User struct {
	Base
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Role         Role   `json:"role"`
}

type UserProfile struct {
	Base
	UserID    uint      `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Phone     string    `json:"phone"`
	BirthDate time.Time `json:"birth_date"`
	AvatarURL string    `json:"avatar_url"`
}

type UserAddress struct {
	Base
	UserID      uint   `json:"user_id"`
	Label       string `json:"label"`
	Country     string `json:"country"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	AddressLine string `json:"address_line"`
	IsDefault   bool   `json:"is_default"`
}
