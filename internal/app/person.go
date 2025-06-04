package app

import (
	"time"
)

type Person struct {
	ID        int64     `json:"id" db:"id"`
	Email     string    `json:"email" db:"email" validate:"required"`
	Phone     string    `json:"phone" db:"phone"`
	FirstName string    `json:"firstName" db:"first_name" validate:"required"`
	LastName  string    `json:"lastName" db:"last_name"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
