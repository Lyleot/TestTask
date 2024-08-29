package models

import "time"

// User model
type User struct {
	ID         int       `db:"user_id"`
	Login      string    `db:"login"`
	FirstName  string    `db:"first_name"`
	SecondName string    `db:"second_name"`
	RoleID     int       `db:"role_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
