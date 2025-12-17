package api

import (
	"technical-test/config"
	"technical-test/database"
	"technical-test/repository"
)

type Repository struct {
	UserRepository     repository.UserRepository
	ShowtimeRepository repository.ShowtimeRepository
}

func Newrepository(
	wrapDB *database.WrapDB,
	env *config.EnvironmentVariable,
) Repository {
	return Repository{
		ShowtimeRepository: repository.NewShowtimeRepository(wrapDB, env),
		UserRepository:     repository.NewUserRepository(wrapDB, env),
	}
}
