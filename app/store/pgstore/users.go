package pgstore

import (
	"TestTask/app/store/pgstore/models"
	"database/sql"
	"errors"
	"fmt"
)

const tabUsers = "users"

// UserRepository предоставляет методы для работы с пользователями в базе данных
type UserRepository struct {
	store *Store // Хранилище данных, которое содержит подключение к базе данных
}

// Create вставляет нового пользователя в таблицу `users` в базе данных.
// Метод использует SQL-запрос с именованными параметрами и возвращает идентификатор созданного пользователя.
func (r *UserRepository) Create(m *models.User) error {
	sqlQuery := fmt.Sprintf("INSERT INTO %s (login, first_name, second_name, role_id, created_at, updated_at) "+
		"VALUES (:login, :first_name, :second_name, :role_id, now(), now()) RETURNING user_id", tabUsers)

	stmt, err := r.store.db.PrepareNamed(sqlQuery)
	if err != nil {
		return err // Возвращает ошибку, если не удалось подготовить запрос
	}

	return stmt.Get(&m.ID, m) // Выполняет запрос и получает идентификатор нового пользователя
}

// Find ищет пользователя по его идентификатору `id`.
// Если пользователь найден, возвращает указатель на структуру `User`. В противном случае возвращает ошибку `ErrRecordNotFound`.
func (r *UserRepository) Find(id int) (*models.User, error) {
	user := &models.User{}

	sqlQuery := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", tabUsers)

	if err := r.store.db.Get(user, sqlQuery, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound // Возвращает ошибку, если пользователь не найден
		}
		return nil, err // Возвращает ошибку, если произошла другая проблема
	}

	return user, nil // Возвращает найденного пользователя
}

// Delete удаляет пользователя по его идентификатору `id`.
// Если удаление прошло успешно, возвращает `nil`. В противном случае возвращает ошибку.
func (r *UserRepository) Delete(id int) error {
	sqlQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", tabUsers)

	_, err := r.store.db.Exec(sqlQuery, id)
	if err != nil {
		return err // Возвращает ошибку, если не удалось выполнить запрос
	}

	return nil // Возвращает `nil`, если удаление прошло успешно
}

// FindAll возвращает список всех пользователей из таблицы `users`.
// Если пользователей нет, возвращает пустой срез. В случае ошибки возвращает `nil` и ошибку.
func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User

	sqlQuery := fmt.Sprintf("SELECT * FROM %s", tabUsers)
	sqlQuery = r.store.db.Rebind(sqlQuery) // Преобразует запрос в формат, поддерживаемый драйвером

	if err := r.store.db.Select(&users, sqlQuery); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return users, nil // Возвращает пустой срез, если пользователей нет
		}
		return nil, err // Возвращает ошибку, если произошла проблема с выполнением запроса
	}

	return users, nil // Возвращает список найденных пользователей
}

// Save обновляет данные пользователя в базе данных по его идентификатору.
// Если пользователь не существует, возвращает ошибку. В противном случае обновляет данные пользователя и возвращает `nil`.
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
		return err // Возвращает ошибку, если не удалось выполнить запрос
	}

	return nil // Возвращает `nil`, если обновление прошло успешно
}
