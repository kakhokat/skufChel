package migrate

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrat(host string) error {
	m, err := migrate.New(
		"file://migrations", fmt.Sprintf(
			"postgres://root:root@%s:5433/root?sslmode=disable", host))
	if err != nil {
		return err
	}

	err = m.Up()
	//todo выглядит как бред, но оставлю
	if errors.Is(err, migrate.ErrNoChange) {
		err := m.Down()
		if err != nil {
			return err
		}

		return m.Up()
	}

	return err
}
