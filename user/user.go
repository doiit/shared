package user

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Email     string    `json:"email" binding:"required,email,lowercase"`
	Roles     []Role    `json:"roles" db:"-"`
	FirstName string    `json:"first_name" binding:"required,lowercase"`
	LastName  string    `json:"last_name" binding:"required,lowercase"`
	Password  string    `json:"password" binding:"required,min=8"`
	// metadata
	UserStatus string `json:"user_status" db:"user_status" binding:"required,oneof=active inactive suspended pending_verification"`
	// Timestamps
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty" db:"last_login_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`

	// Audit
	CreatedBy *uuid.UUID `json:"created_by,omitempty" db:"created_by"`
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty" db:"updated_by"`
}
type Role struct {
	ID          uuid.UUID  `json:"id" db:"id" binding:"required,uuid"`
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description" binding:"required"`
	Value       uint16     `json:"value" binding:"required"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type UserResponse struct {
	ID          uuid.UUID  `json:"id" binding:"required,uuid"`
	Email       string     `json:"email" binding:"required,email"`
	FirstName   string     `json:"first_name" binding:"required"`
	LastName    string     `json:"last_name" binding:"required"`
	UserStatus  string     `json:"user_status" binding:"required"`
	Roles       []Role     `json:"roles" binding:"dive,required"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
}
