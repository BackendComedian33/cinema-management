package repository

import (
	"context"
	"errors"
	"technical-test/config"
	"technical-test/database"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

type UserRepository interface {
	Login(ctx context.Context, tx pgx.Tx, email string) (userID int, password string, err error)
}

type UserRepositoryImpl struct {
	WrapDB *database.WrapDB
	Env    *config.EnvironmentVariable
}

func NewUserRepository(
	wrapDB *database.WrapDB,
	env *config.EnvironmentVariable,
) UserRepository {
	return &UserRepositoryImpl{
		WrapDB: wrapDB,
		Env:    env,
	}
}

func (r *UserRepositoryImpl) Login(ctx context.Context, tx pgx.Tx, email string) (userID int, password string, err error) {
	query := `SELECT id,password FROM users WHERE LOWER(email) = LOWER($1)`
	if tx != nil {
		err = tx.QueryRow(ctx, query, email).Scan(&userID, &password)
		if errors.Is(err, pgx.ErrNoRows) {
			// expected no rows
			err = nil
			return 0, password, nil
		} else if err != nil {
			log.Error().Err(err).Msg("Failed to get user ")
			return
		}
	} else {
		err = r.WrapDB.Postgres.QueryRow(ctx, query, email).Scan(&userID, &password)
		if errors.Is(err, pgx.ErrNoRows) {
			// expected no rows
			err = nil
			return 0, password, nil
		} else if err != nil {
			log.Error().Err(err).Msg("Failed to get user ")
			return
		}

	}

	return
}
