package server

import (
	"TestTask/app/conf"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"

	migrate "github.com/rubenv/sql-migrate"
)

const MigrationsDir = "migrations" // Директория для миграций

// ParseParams обрабатывает флаги командной строки для выполнения миграций.
func ParseParams(config conf.Config) {

	var (
		migrateCmd string
	)

	// Определение флага командной строки для миграций
	flag.StringVar(&migrateCmd, "migrate", "", "up - выполнить все миграции Up\n"+
		"down - выполнить все миграции Down\n"+
		"number - количество шагов для миграции (если > 0 - выполнить number шагов Up, если < 0 - выполнить number шагов Down)\n"+
		"status - показать список применённых миграций\n",
	)

	flag.Parse()

	if len(migrateCmd) > 0 {
		ExecMigrate(config, migrateCmd) // Выполнение миграции в зависимости от команды
	}

}

// ExecMigrate выполняет миграции в зависимости от указанной команды.
func ExecMigrate(config conf.Config, migrateCmd string) {

	var steps int

	log := config.LOG().WithField("module", "migrate")

	switch migrateCmd {
	case "status":
		// Получение и вывод статуса миграций
		recs, err := migrate.GetMigrationRecords(config.DB().SqlDB(), "postgres")
		if err != nil {
			log.WithError(err).Fatal("failed to get migration records")
		}
		for _, rec := range recs {
			fmt.Println(rec.AppliedAt.Format("2006-01-02 15:04 "), rec.Id)
		}
		os.Exit(0)
	case "up":
		steps = 999 // Выполнить все миграции Up
	case "down":
		steps = -999 // Выполнить все миграции Down
	default:
		if regexp.MustCompile(`^-?[0-9]+$`).Match([]byte(migrateCmd)) { // Проверка, что параметр - это целое число
			var err error
			steps, err = strconv.Atoi(migrateCmd)
			if err != nil {
				log.WithError(err).Fatal("failed to convert migrate argument to digit")
			}
		} else {
			log.Fatal("unknown command") // Неизвестная команда
		}
	}

	// Выполнение миграций
	if steps != 0 {

		migrations := &migrate.FileMigrationSource{
			Dir: MigrationsDir,
		}

		var direction migrate.MigrationDirection

		if steps > 0 {
			direction = migrate.Up // Направление миграций Up
		} else if steps < 0 {
			direction = migrate.Down // Направление миграций Down
			steps = -steps
		}

		n, err := migrate.ExecMax(config.DB().SqlDB(), "postgres", migrations, direction, steps)
		if err != nil {
			log.WithError(err).Fatal("failed to execute migrations")
		}

		if direction == migrate.Down {
			n = -n
		}

		log.Infof("%d steps done", n) // Логирование количества выполненных шагов

		os.Exit(0)
	}
}
