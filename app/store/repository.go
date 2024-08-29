package store

import "TestTask/app/store/pgstore/models"

// HealthcheckRepository определяет метод для проверки состояния системы или базы данных.
type HealthcheckRepository interface {
	Check() error
}

// RoleRepository определяет методы для управления ролями в базе данных.
type RoleRepository interface {
	Create(m *models.Role) error       // Добавляет новую роль
	Find(id int) (*models.Role, error) // Ищет роль по идентификатору
	Delete(id int) error               // Удаляет роль по идентификатору
	FindAll() ([]models.Role, error)   // Возвращает все роли
	Save(m *models.Role) error         // Обновляет существующую роль
}

// UserRepository определяет методы для управления пользователями в базе данных.
type UserRepository interface {
	Create(m *models.User) error       // Добавляет нового пользователя
	Find(id int) (*models.User, error) // Ищет пользователя по идентификатору
	Delete(id int) error               // Удаляет пользователя по идентификатору
	FindAll() ([]models.User, error)   // Возвращает всех пользователей
	Save(m *models.User) error         // Обновляет существующего пользователя
}
