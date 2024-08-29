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

const MigrationsDir = "migrations"

func ParseParams(config conf.Config) {

	var (
		migrateCmd string
	)

	flag.StringVar(&migrateCmd, "migrate", "", "up - migrate all steps Up\n"+
		"down - migrate all steps Down\n"+
		"number - amount of steps to migrate (if > 0 - migrate number steps up, if < 0 migrate number steps down)\n"+
		"status - show list of applyed migrations\n",
	)

	flag.Parse()

	if len(migrateCmd) > 0 {
		ExecMigrate(config, migrateCmd)
	}

}

func ExecMigrate(config conf.Config, migrateCmd string) {

	var steps int

	log := config.LOG().WithField("module", "migrate")

	switch migrateCmd {
	case "status":
		recs, err := migrate.GetMigrationRecords(config.DB().SqlDB(), "postgres")
		if err != nil {
			log.WithError(err).Fatal("failed to get migration records")
		}
		for _, rec := range recs {
			fmt.Println(rec.AppliedAt.Format("2006-01-02 15:04 "), rec.Id)
		}
		os.Exit(0)
	case "up":
		steps = 999
	case "down":
		steps = -999
	default:
		if regexp.MustCompile(`^-?[0-9]+$`).Match([]byte(migrateCmd)) { // check if parameter is Integer
			var err error
			steps, err = strconv.Atoi(migrateCmd)
			if err != nil {
				log.WithError(err).Fatal("failed to convert migrate argument to digit")
			}
		} else {
			log.Fatal("unknown command")
		}
	}

	// Migrator
	if steps != 0 {

		migrations := &migrate.FileMigrationSource{
			Dir: MigrationsDir,
		}

		var direction migrate.MigrationDirection

		if steps > 0 {
			direction = migrate.Up
		} else if steps < 0 {
			direction = migrate.Down
			steps = -steps
		}

		n, err := migrate.ExecMax(config.DB().SqlDB(), "postgres", migrations, direction, steps)
		if err != nil {
			log.WithError(err).Fatal("failed to execute migrations")
		}

		if direction == migrate.Down {
			n = -n
		}

		log.Infof("%d steps done", n)

		os.Exit(0)
	}
}
