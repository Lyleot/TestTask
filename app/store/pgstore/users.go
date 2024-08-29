package pgstore

import (
	"TestTask/app/store/pgstore/models"
	"database/sql"
	"errors"
	"fmt"
)

const tabUsers = "users"

type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(m *models.User) error {
	sqlQuery := fmt.Sprintf("INSERT INTO %s (login, first_name, second_name, role_id, created_at, updated_at) "+
		"VALUES (:login, :first_name, :second_name, :role_id, now(), now()) RETURNING user_id", tabUsers)

	stmt, err := r.store.db.PrepareNamed(sqlQuery)
	if err != nil {
		return err
	}

	return stmt.Get(&m.ID, m)
}

// Find ...
func (r *UserRepository) Find(id int) (*models.User, error) {
	user := &models.User{}

	sqlQuery := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", tabUsers)

	if err := r.store.db.Get(user, sqlQuery, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return user, nil
}

// Delete ...
func (r *UserRepository) Delete(id int) error {
	sqlQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", tabUsers)

	_, err := r.store.db.Exec(sqlQuery, id)
	if err != nil {
		return err
	}

	return nil
}

// FindAll ...
func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User

	sqlQuery := fmt.Sprintf("SELECT * FROM %s", tabUsers)
	sqlQuery = r.store.db.Rebind(sqlQuery)

	if err := r.store.db.Select(&users, sqlQuery); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return users, nil
		}
		return nil, err
	}

	return users, nil
}

// Save ...
func (r *UserRepository) Save(m *models.User) error {
	sqlQuery := fmt.Sprintf(
		"UPDATE %s SET updated_at = now(),"+
			"login = :login, "+
			"first_name = :first_name, "+
			"second_name = :second_name, "+
			"role_id = :role_id "+
			"WHERE user_id = :user_id",
		tabUsers,
	)

	_, err := r.store.db.NamedExec(sqlQuery, m)
	if err != nil {
		return err
	}

	return nil
}
