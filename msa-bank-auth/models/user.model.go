package models

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Login      string    `gorm:"type:varchar(255)"`
	Password   string    `gorm:"not null"`
	Position   string    `gorm:"type:varchar(255)"`
	Email      string    `gorm:"type:varchar(255)"`
	LastName   string    `gorm:"type:varchar(255)"`
	FirstName  string    `gorm:"type:varchar(255)"`
	MiddleName string    `gorm:"type:varchar(255)"`
	Role       string    `gorm:"type:varchar(255)"`
}

type SignUpInput struct {
	Login           string `json:"login" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	Position        string `json:"position" binding:"required"`
	Email           string `json:"email" binding:"required"`
	LastName        string `json:"lastName" binding:"required"`
	FirstName       string `json:"firstName" binding:"required"`
	MiddleName      string `json:"middleName" binding:"required"`
	Role            string `json:"role" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
}
