package conf

import (
	"TestTask/app/store"
	"TestTask/app/store/pgstore"
	"fmt"
	"time"

	"github.com/caarlos0/env"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Name     string `env:"PG_DB_NAME,required"`
	Host     string `env:"PG_HOST,required"`
	Port     int    `env:"PG_PORT" envDefault:"5432"`
	User     string `env:"PG_USER,required"`
	Password string `env:"PG_PASSWORD,required"`
	SSL      string `env:"PG_SSL" envDefault:"disable"`
	TimeZone string `env:"PG_TIMEZONE" envDefault:"UTC"`

	MaxOpenConns    int           `env:"PG_MAX_OPEN_CONNS" envDefault:"20"`
	MaxIdleConns    int           `env:"PG_MAX_IDLE_CONNS" envDefault:"10"`
	ConnMaxLifetime time.Duration `env:"PG_CONN_MAX_LIFETIME" envDefault:"5m"`
}

func (d *DB) URL() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s TimeZone=%s", d.Host, d.Port, d.User, d.Password, d.Name, d.SSL, d.TimeZone)
}

// ConnectPostgreSQL setup postgres
func (d *DB) connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", d.URL())

	if err != nil {
		return nil, err
	}

	for i := 0; i < 5; i++ {
		if err := db.Ping(); err == nil {
			return db, nil
		} else {
			fmt.Println("can't connect to DB retry after 2 seconds")
			time.Sleep(2 * time.Second)
		}
	}

	return nil, err
}

func (c *config) DB() store.Store {

	if c.db != nil {
		return c.db
	}

	c.Lock()
	defer c.Unlock()

	if c.db != nil {
		return c.db
	}

	dbConf := &DB{}

	if err := env.Parse(dbConf); err != nil {
		panic(err)
	}

	var err error

	client, err := dbConf.connect()
	if err != nil {
		panic(err)
	}

	client.DB.SetMaxOpenConns(dbConf.MaxOpenConns)
	client.DB.SetConnMaxLifetime(dbConf.ConnMaxLifetime)
	client.DB.SetMaxIdleConns(dbConf.MaxIdleConns)

	c.db = pgstore.New(client)

	return c.db
}
