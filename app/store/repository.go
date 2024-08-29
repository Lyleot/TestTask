package store

import "TestTask/app/store/pgstore/models"

type HealthcheckRepository interface {
	Check() error
}

// RoleRepository ...
type RoleRepository interface {
	Create(m *models.Role) error
	Find(id int) (*models.Role, error)
	Delete(id int) error
	FindAll() ([]models.Role, error)
	Save(m *models.Role) error
}

// UserRepository ...
type UserRepository interface {
	Create(m *models.User) error
	Find(id int) (*models.User, error)
	Delete(id int) error
	FindAll() ([]models.User, error)
	Save(m *models.User) error
}
