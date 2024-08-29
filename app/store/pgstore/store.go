package pgstore

import (
	"TestTask/app/store"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db IDB // Интерфейс для работы с базой данных

	healthcheckRepository *HealthcheckRepository // Репозиторий для проверки состояния
	roleRepository        *RoleRepository        // Репозиторий для ролей
	userRepository        *UserRepository        // Репозиторий для пользователей
}

// New создает новый экземпляр Store с переданной базой данных
func New(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

// Healthcheck возвращает репозиторий проверки состояния, создавая его при необходимости
func (s *Store) Healthcheck() store.HealthcheckRepository {
	if s.healthcheckRepository != nil {
		return s.healthcheckRepository
	}

	s.healthcheckRepository = &HealthcheckRepository{
		store: s,
	}

	return s.healthcheckRepository
}

// Role возвращает репозиторий ролей, создавая его при необходимости
func (s *Store) Role() store.RoleRepository {
	if s.roleRepository != nil {
		return s.roleRepository
	}

	s.roleRepository = &RoleRepository{
		store: s,
	}

	return s.roleRepository
}

// User возвращает репозиторий пользователей, создавая его при необходимости
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
