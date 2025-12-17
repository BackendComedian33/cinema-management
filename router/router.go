package router

import (
	"technical-test/config"
	"technical-test/handler"
	"technical-test/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	Env         *config.EnvironmentVariable
	UserHandler handler.UserHandler
	Middleware  middleware.Middleware
}

func NewRouter(handler Handler) *gin.Engine {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	HelloWorld(router)

	var apiRouterGroupName = "/api"

	// Api router
	apiRouterGroup := router.Group(apiRouterGroupName)

	if handler.Env.App.Debug {
		SwaggerRouter(apiRouterGroup)
	}

	RouterApiV1(handler, apiRouterGroup)

	return router

}

func HelloWorld(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello World")
	})
}

func SwaggerRouter(router *gin.RouterGroup) {
	router.GET("/docs/*any", func(ctx *gin.Context) {
		ctx.Header("Content-Security-Policy",
			"default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data:; font-src 'self' data:; connect-src 'self' https:; object-src 'none'; base-uri 'self'")
		ctx.Next()
	}, ginSwagger.WrapHandler(swaggerFiles.Handler))
}
