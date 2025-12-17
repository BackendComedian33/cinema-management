package postgres

import (
	"context"
	"errors"
	"fmt"
	"technical-test/config"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type WrapDatabase struct {
	Conn *pgxpool.Pool
}

const MIGRATION_LOCATIONS = "database/migrations"

func NewDBConnection(env *config.EnvironmentVariable) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.Database.Postgres.User,
		env.Database.Postgres.Password,
		env.Database.Postgres.Host,
		env.Database.Postgres.Port,
		env.Database.Postgres.Name,
	)

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatal().Err(err).Str("database", env.Database.Postgres.Name).Msg("[x] failed to parse connection config for postgres")
		panic(err)
	}

	config.MaxConns = int32(env.Database.Postgres.MaxConnections)
	config.MaxConnIdleTime = time.Minute * time.Duration(env.Database.Postgres.MaxIdleTime)

	log.Info().Msgf("Connecting to Postgres at %s", connStr)
	conn, err := pgxpool.New(ctx, config.ConnString())
	if err != nil {
		log.Fatal().Err(err).Str("database", env.Database.Postgres.Name).Msg("[x] failed to connect to postgres")
		panic(err)
	}
	log.Info().Msgf("[+] Successfully connected to Postgres at %s", connStr)

	log.Info().Msgf("Pinging Postgres at %s", connStr)
	if err := conn.Ping(ctx); err != nil {
		log.Fatal().Err(err).Str("database", env.Database.Postgres.Name).Msg("[x] failed to ping postgres")
		panic(err)
	}
	log.Info().Msgf("[+] Successfully pinged Postgres at %s", connStr)

	return conn
}

func InitMigrations(env *config.EnvironmentVariable) error {
	log.Info().Msg("Checking migrations")

	m, err := migrate.New(
		"file://"+MIGRATION_LOCATIONS, // add scheme
		env.GetDBUrl(),
	)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed load migrations")
		return err
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info().Msg("Database has not changed")
			return nil
		}
		log.Fatal().Err(err).Msg("Failed to run migration")
		return err
	}

	log.Info().Msg("migrations success")
	return nil
}
