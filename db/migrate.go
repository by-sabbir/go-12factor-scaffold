package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func (d *DataBase) MigrateDB() error {

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create pg driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)

	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if err.Error() == migrate.ErrNoChange.Error() {
		} else {
			log.Warn("migrating down: ", err)
			if err := m.Steps(-1); err != nil {
				logrus.Error("failed to migrate to previous version")
			}

			log.Error("could not migrate: ", err)
			return err
		}
	}
	log.Info("Successfully migrated database")
	return nil
}
