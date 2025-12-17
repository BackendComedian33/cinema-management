package api

import (
	"technical-test/config"
	"technical-test/database"
	"technical-test/service"
)

type Service struct {
	UserService service.UserService
}

func NewService(
	env *config.EnvironmentVariable,
	repo Repository,
	db *database.WrapDB,
) Service {
	return Service{
		UserService: service.NewUserService(
			repo.UserRepository,
			repo.ShowtimeRepository,
			db,
			env,
		),
	}
}
