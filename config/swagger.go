package config

import "github.com/swaggo/swag/example/basic/docs"

const (
	ReleaseVersion = "0.0.1"
)

func InitSwagger(env *EnvironmentVariable) {
	if env.App.Debug {
		docs.SwaggerInfo.Host = env.Swagger.Host
		docs.SwaggerInfo.Version = ReleaseVersion
		docs.SwaggerInfo.Host = env.Swagger.Host
		docs.SwaggerInfo.BasePath = "/api/v1"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	}
}
