package pgstore

import (
	"TestTask/app/store/pgstore/models"
	"database/sql"
	"errors"
	"fmt"
)

const tabRoles = "roles" // Название таблицы для ролей

var (
	ErrRecordNotFound = errors.New("record not found") // Ошибка, если запись не найдена
)

// RoleRepository обеспечивает доступ к данным ролей в базе данных.
type RoleRepository struct {
	store *Store // Хранилище данных, использующее интерфейс IDB
}

// Create добавляет новую роль в базу данных и возвращает её идентификатор.
// При успешном выполнении возвращает nil, в противном случае - ошибку.
func (r *RoleRepository) Create(m *models.Role) error {
	sqlQuery := fmt.Sprintf("INSERT INTO %s (name, description, created_at, updated_at) "+
		"VALUES (:name, :description, now(), now()) RETURNING role_id", tabRoles)

	stmt, err := r.store.db.PrepareNamed(sqlQuery)
	if err != nil {
		return err // Возвращает ошибку подготовки запроса
	}

	return stmt.Get(&m.ID, m) // Выполняет запрос и получает идентификатор созданной роли
}

// Find находит роль по её идентификатору и возвращает её.
// Если роль не найдена, возвращает ErrRecordNotFound, в противном случае возвращает роль и nil.
func (r *RoleRepository) Find(id int) (*models.Role, error) {
	role := &models.Role{}

	sqlQuery := fmt.Sprintf("SELECT * FROM %s WHERE role_id = $1", tabRoles)

	if err := r.store.db.Get(role, sqlQuery, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound // Роль не найдена
		}
		return nil, err // Возвращает ошибку выполнения запроса
	}

	return role, nil // Возвращает найденную роль
}

// Delete удаляет роль по её идентификатору.
// При успешном удалении возвращает nil, в противном случае - ошибку.
func (r *RoleRepository) Delete(id int) error {
	sqlQuery := fmt.Sprintf("DELETE FROM %s WHERE role_id = $1", tabRoles)

	_, err := r.store.db.Exec(sqlQuery, id)
	if err != nil {
		return err // Возвращает ошибку выполнения запроса
	}

	return nil // Возвращает nil при успешном удалении
}

// FindAll возвращает все роли из базы данных.
// Если ролей нет, возвращает пустой срез. При ошибке возвращает nil и ошибку.
func (r *RoleRepository) FindAll() ([]models.Role, error) {
	var roles []models.Role

	sqlQuery := fmt.Sprintf("SELECT * FROM %s", tabRoles)
	sqlQuery = r.store.db.Rebind(sqlQuery) // Преобразует запрос для поддержки специфичных форматов драйвера

	if err := r.store.db.Select(&roles, sqlQuery); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return roles, nil // Возвращает пустой список, если нет записей
		}
		return nil, err // Возвращает ошибку выполнения запроса
	}

	return roles, nil // Возвращает список ролей
}

// Save обновляет существующую роль по её идентификатору.
// Если роль не найдена, возвращает ошибку, в противном случае выполняет обновление и возвращает nil.
func (r *RoleRepository) Save(m *models.Role) error {
	// Проверка существования роли перед обновлением
	_, err := r.Find(m.ID)
	if err != nil {
		return errors.New("the role you save does not exist") // Роль не существует
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
		return err // Возвращает ошибку выполнения запроса
	}

	return nil // Возвращает nil при успешном обновлении
}
