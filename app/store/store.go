package store

import (
	"context"
	"database/sql"
)

type Store interface {
	SqlDB() *sql.DB

	DBStats() sql.DBStats
	Begin() (Store, error)
	Commit() error
	Rollback() error
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (Store, error)

	Healthcheck() HealthcheckRepository
	Role() RoleRepository
	User() UserRepository
}
