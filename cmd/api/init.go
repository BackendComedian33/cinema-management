package api

import (
	"technical-test/config"
	"technical-test/database"
	"technical-test/handler"
	"technical-test/middleware"
	"technical-test/repository"
	"technical-test/router"
	"technical-test/service"

	"github.com/gin-gonic/gin"
)

type Setup struct {
	Router  *gin.Engine
	Handler Handler
	WrapDB  *database.WrapDB
}

func Init(env *config.EnvironmentVariable) (*Setup, error) {
	wrapDB := database.InitDB(env)
	userRepo := repository.NewUserRepository(wrapDB, env)
	showTimeRepo := repository.NewShowtimeRepository(wrapDB, env)

	userService := service.NewUserService(userRepo, showTimeRepo, wrapDB, env)

	userHandler := handler.NewUserHandler(env, userService)

	// middleware
	middleware := middleware.NewMiddleware(env)
	router := router.NewRouter(router.Handler{
		Env:         env,
		Middleware:  middleware,
		UserHandler: userHandler,
	})
	return &Setup{
		Router: router,
		WrapDB: wrapDB,
	}, nil
}
