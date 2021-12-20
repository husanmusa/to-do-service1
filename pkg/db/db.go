package db

import (
	"fmt"

	"github.com/husanmusa/to-do-service/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres drivers
)

func ConnectToDB(cfg config.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDB, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, err
	}

	return connDB, nil
}

func ConnectToDBForSuite(cfg config.Config) (*sqlx.DB, func()) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)
	connDB, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		panic(err)
	}

	cleanUpFunc := func() {
		connDB.Close()
	}

	return connDB, cleanUpFunc
}
