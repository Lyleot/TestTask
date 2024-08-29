package models

import "time"

// User представляет модель пользователя в базе данных.
type User struct {
	ID         int       `db:"user_id"`     // Идентификатор пользователя
	Login      string    `db:"login"`       // Логин пользователя
	FirstName  string    `db:"first_name"`  // Имя пользователя
	SecondName string    `db:"second_name"` // Фамилия пользователя
	RoleID     int       `db:"role_id"`     // Идентификатор роли пользователя
	CreatedAt  time.Time `db:"created_at"`  // Дата и время создания пользователя
	UpdatedAt  time.Time `db:"updated_at"`  // Дата и время последнего обновления пользователя
}
