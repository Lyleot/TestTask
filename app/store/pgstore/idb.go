package pgstore

import (
	"TestTask/app/store"
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	ErrRecursiveBegin = errors.New("recursive calling restricted")
	ErrNotTx          = errors.New("transaction did not start")
)

// IDB interface to *sqlx.DB and *sqlx.Tx
type IDB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)

	DriverName() string
	Rebind(query string) string
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	MustExec(query string, args ...interface{}) sql.Result
	Preparex(query string) (*sqlx.Stmt, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
}

func (s *Store) Begin() (store.Store, error) {

	if _, yes := s.db.(*sqlx.Tx); yes {
		return nil, ErrRecursiveBegin
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

func (s *Store) DBStats() sql.DBStats {
	return s.SqlDB().Stats()
}

func (s *Store) Commit() error {
	if tx, yes := s.db.(*sqlx.Tx); yes {
		return tx.Commit()
	}
	return ErrNotTx
}

func (s *Store) Rollback() error {
	if tx, yes := s.db.(*sqlx.Tx); yes {
		return tx.Rollback()
	}
	return ErrNotTx
}

func (s *Store) SqlDB() *sql.DB {
	if d, yes := s.db.(*sqlx.DB); yes {
		return d.DB
	}
	return nil
}

func (s *Store) BeginTxx(ctx context.Context, opts *sql.TxOptions) (store.Store, error) {

	if _, yes := s.db.(*sqlx.Tx); yes {
		return nil, ErrRecursiveBegin
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
