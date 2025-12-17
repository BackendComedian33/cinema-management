package api

import (
	"technical-test/config"
	"technical-test/handler"
)

type Handler struct {
	UserHandler handler.UserHandler
}

func NewHandler(
	service Service,
	env *config.EnvironmentVariable,
) Handler {
	return Handler{
		UserHandler: handler.NewUserHandler(env, service.UserService),
	}
}
