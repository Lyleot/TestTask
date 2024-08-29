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

// DB содержит параметры подключения к базе данных PostgreSQL.
type DB struct {
	Name     string `env:"PG_DB_NAME,required"`          // Имя базы данных
	Host     string `env:"PG_HOST,required"`             // Хост базы данных
	Port     int    `env:"PG_PORT" envDefault:"5432"`    // Порт базы данных
	User     string `env:"PG_USER,required"`             // Пользователь базы данных
	Password string `env:"PG_PASSWORD,required"`         // Пароль пользователя
	SSL      string `env:"PG_SSL" envDefault:"disable"`  // Режим SSL
	TimeZone string `env:"PG_TIMEZONE" envDefault:"UTC"` // Часовой пояс

	MaxOpenConns    int           `env:"PG_MAX_OPEN_CONNS" envDefault:"20"`    // Максимум открытых соединений
	MaxIdleConns    int           `env:"PG_MAX_IDLE_CONNS" envDefault:"10"`    // Максимум неактивных соединений
	ConnMaxLifetime time.Duration `env:"PG_CONN_MAX_LIFETIME" envDefault:"5m"` // Максимальное время жизни соединения
}

// URL формирует строку подключения к базе данных.
func (d *DB) URL() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s TimeZone=%s", d.Host, d.Port, d.User, d.Password, d.Name, d.SSL, d.TimeZone)
}

// connect устанавливает подключение к базе данных PostgreSQL.
func (d *DB) connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", d.URL()) // Подключается к базе данных с использованием pgx драйвера

	if err != nil {
		return nil, err // Возвращает ошибку, если подключение не удалось
	}

	// Пытается подключиться к базе данных с повторными попытками
	for i := 0; i < 5; i++ {
		if err := db.Ping(); err == nil {
			return db, nil // Возвращает подключение при успешной проверке
		} else {
			fmt.Println("can't connect to DB retry after 2 seconds") // Логирует ошибку и делает паузу
			time.Sleep(2 * time.Second)
		}
	}

	return nil, err // Возвращает ошибку, если подключение не удалось
}

// DB возвращает экземпляр хранилища, инициализируя его при необходимости.
func (c *config) DB() store.Store {
	if c.db != nil {
		return c.db // Возвращает уже инициализированное хранилище
	}

	c.Lock()
	defer c.Unlock()

	if c.db != nil {
		return c.db // Возвращает уже инициализированное хранилище после блокировки
	}

	dbConf := &DB{}

	if err := env.Parse(dbConf); err != nil {
		panic(err) // Паника при ошибке загрузки переменных окружения
	}

	var err error

	client, err := dbConf.connect()
	if err != nil {
		panic(err) // Паника при ошибке подключения к базе данных
	}

	client.DB.SetMaxOpenConns(dbConf.MaxOpenConns)       // Устанавливает максимальное количество открытых соединений
	client.DB.SetConnMaxLifetime(dbConf.ConnMaxLifetime) // Устанавливает максимальное время жизни соединения
	client.DB.SetMaxIdleConns(dbConf.MaxIdleConns)       // Устанавливает максимальное количество неактивных соединений

	c.db = pgstore.New(client) // Инициализирует хранилище и сохраняет его в конфигурации

	return c.db
}
