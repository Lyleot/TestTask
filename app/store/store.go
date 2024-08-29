package store

import (
	"context"
	"database/sql"
)

// Store предоставляет интерфейс для работы с базой данных и управления репозиториями.
type Store interface {
	// SqlDB возвращает объект *sql.DB для низкоуровневого доступа к базе данных.
	SqlDB() *sql.DB

	// DBStats возвращает статистику соединений с базой данных.
	DBStats() sql.DBStats

	// Begin начинает новую транзакцию и возвращает новый Store для управления транзакцией.
	Begin() (Store, error)

	// Commit завершает транзакцию, делая все изменения постоянными.
	Commit() error

	// Rollback откатывает транзакцию, отменяя все изменения.
	Rollback() error

	// BeginTxx начинает новую транзакцию с заданными опциями и возвращает новый Store.
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (Store, error)

	// Healthcheck возвращает репозиторий для проверки состояния системы.
	Healthcheck() HealthcheckRepository

	// Role возвращает репозиторий для работы с ролями.
	Role() RoleRepository

	// User возвращает репозиторий для работы с пользователями.
	User() UserRepository
}
