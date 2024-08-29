package models

import "time"

// Role представляет модель роли в базе данных.
type Role struct {
	ID          int       `db:"role_id"`     // Идентификатор роли
	Name        string    `db:"name"`        // Название роли
	Description string    `db:"description"` // Описание роли
	CreatedAt   time.Time `db:"created_at"`  // Дата и время создания роли
	UpdatedAt   time.Time `db:"updated_at"`  // Дата и время последнего обновления роли
}
