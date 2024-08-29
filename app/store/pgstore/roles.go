package pgstore

import (
	"TestTask/app/store/pgstore/models"
	"database/sql"
	"errors"
	"fmt"
)

const tabRoles = "roles"

var (
	ErrRecordNotFound = errors.New("record not found")
)

type RoleRepository struct {
	store *Store
}

// Create ...
func (r *RoleRepository) Create(m *models.Role) error {
	sqlQuery := fmt.Sprintf("INSERT INTO %s (name, description, created_at, updated_at) "+
		"VALUES (:name, :description, now(), now()) RETURNING role_id", tabRoles)

	stmt, err := r.store.db.PrepareNamed(sqlQuery)
	if err != nil {
		return err
	}

	return stmt.Get(&m.ID, m)
}

// Find ...
func (r *RoleRepository) Find(id int) (*models.Role, error) {
	role := &models.Role{}

	sqlQuery := fmt.Sprintf("SELECT * FROM %s WHERE role_id = $1", tabRoles)

	if err := r.store.db.Get(role, sqlQuery, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return role, nil
}

// Delete ...
func (r *RoleRepository) Delete(id int) error {
	sqlQuery := fmt.Sprintf("DELETE FROM %s WHERE role_id = $1", tabRoles)

	_, err := r.store.db.Exec(sqlQuery, id)
	if err != nil {
		return err
	}

	return nil
}

// FindAll ...
func (r *RoleRepository) FindAll() ([]models.Role, error) {
	var roles []models.Role

	sqlQuery := fmt.Sprintf("SELECT * FROM %s", tabRoles)
	sqlQuery = r.store.db.Rebind(sqlQuery)

	if err := r.store.db.Select(&roles, sqlQuery); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return roles, nil
		}
		return nil, err
	}

	return roles, nil
}

// Save ...
func (r *RoleRepository) Save(m *models.Role) error {
	_, err := r.Find(m.ID)
	if err != nil {
		fmt.Println(err)
		return errors.New("the role you save does not exist")
	}

	sqlQuery := fmt.Sprintf(
		"UPDATE %s SET updated_at = now(),"+
			"name = :name, "+
			"description = :description "+
			"WHERE role_id = :role_id",
		tabRoles,
	)

	_, err = r.store.db.NamedExec(sqlQuery, m)
	if err != nil {
		return err
	}

	return nil
}
