package pgstore

import (
	"TestTask/app/store"
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	ErrRecursiveBegin = errors.New("recursive calling restricted") // Ошибка рекурсивного вызова
	ErrNotTx          = errors.New("transaction did not start")    // Ошибка: транзакция не началась
)

// IDB интерфейс для работы с базой данных
type IDB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)             // Выполнить запрос
	DriverName() string                                                     // Название драйвера
	Rebind(query string) string                                             // Переменные для драйвера
	BindNamed(query string, arg interface{}) (string, []interface{}, error) // Привязка именованных переменных
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)           // Именованный запрос
	NamedExec(query string, arg interface{}) (sql.Result, error)            // Именованное выполнение запроса
	Select(dest interface{}, query string, args ...interface{}) error       // Выборка данных
	Get(dest interface{}, query string, args ...interface{}) error          // Получение одной записи
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)           // Запрос с возвратом строк
	QueryRowx(query string, args ...interface{}) *sqlx.Row                  // Запрос с одной строкой
	MustExec(query string, args ...interface{}) sql.Result                  // Обязательное выполнение запроса
	Preparex(query string) (*sqlx.Stmt, error)                              // Подготовка запроса
	PrepareNamed(query string) (*sqlx.NamedStmt, error)                     // Подготовка именованного запроса
}

// Begin начинает новую транзакцию
func (s *Store) Begin() (store.Store, error) {

	if _, yes := s.db.(*sqlx.Tx); yes {
		return nil, ErrRecursiveBegin // Ошибка рекурсивного вызова
	}

	tx, err := s.db.(*sqlx.DB).Beginx()
	if err != nil {
		return nil, err
	}

	clone := &Store{
		db: tx,
	}

	return clone, nil
}

// DBStats возвращает статистику подключения
func (s *Store) DBStats() sql.DBStats {
	return s.SqlDB().Stats()
}

// Commit подтверждает транзакцию
func (s *Store) Commit() error {
	if tx, yes := s.db.(*sqlx.Tx); yes {
		return tx.Commit()
	}
	return ErrNotTx // Ошибка: транзакция не началась
}

// Rollback откатывает транзакцию
func (s *Store) Rollback() error {
	if tx, yes := s.db.(*sqlx.Tx); yes {
		return tx.Rollback()
	}
	return ErrNotTx // Ошибка: транзакция не началась
}

// SqlDB возвращает *sql.DB из *sqlx.DB
func (s *Store) SqlDB() *sql.DB {
	if d, yes := s.db.(*sqlx.DB); yes {
		return d.DB
	}
	return nil
}

// BeginTxx начинает транзакцию с контекстом и опциями
func (s *Store) BeginTxx(ctx context.Context, opts *sql.TxOptions) (store.Store, error) {

	if _, yes := s.db.(*sqlx.Tx); yes {
		return nil, ErrRecursiveBegin // Ошибка рекурсивного вызова
	}

	tx, err := s.db.(*sqlx.DB).BeginTxx(ctx, opts)
	if err != nil {
		return nil, err
	}

	clone := &Store{
		db: tx,
	}

	return clone, nil
}
