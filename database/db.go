package database

import (
	"technical-test/config"
	"technical-test/database/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

type WrapDB struct {
	Postgres *pgxpool.Pool
}

func InitDB(env *config.EnvironmentVariable) *WrapDB {
	postgresDB := postgres.NewDBConnection(env)

	if env.Database.Postgres.UseMigration {
		// Init migrations
		err := postgres.InitMigrations(env)
		if err != nil {
			return nil
		}
	}

	return &WrapDB{
		Postgres: postgresDB,
	}
}
