package pgstore

import (
	"TestTask/app/store"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db IDB

	healthcheckRepository *HealthcheckRepository

	roleRepository *RoleRepository
	userRepository *UserRepository
}

// New ...
func New(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

// Healthcheck ...
func (s *Store) Healthcheck() store.HealthcheckRepository {
	if s.healthcheckRepository != nil {
		return s.healthcheckRepository
	}

	s.healthcheckRepository = &HealthcheckRepository{
		store: s,
	}

	return s.healthcheckRepository
}

// Role ...
func (s *Store) Role() store.RoleRepository {
	if s.roleRepository != nil {
		return s.roleRepository
	}

	s.roleRepository = &RoleRepository{
		store: s,
	}

	return s.roleRepository
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
