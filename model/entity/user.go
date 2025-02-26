package entity

import "time"

type User struct {
	ID            string    `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Email         string    `json:"email"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	EmailVerified bool      `json:"email_verified"`
	SID           string    `json:"sid"`
}

func NewUserFromAuth(id, firstName, lastName, email string, emailVerified bool) *User {
	return &User{
		ID:            id,
		FirstName:     firstName,
		LastName:      lastName,
		Email:         email,
		EmailVerified: emailVerified,
	}
}
