package config

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func LoadEnv() (env *EnvironmentVariable, err error) {
	log.Info().Msg("Load Env Here")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	SetDefaultConfig(viper.GetViper())

	err = viper.ReadInConfig()
	if err != nil {
		log.Error().Err(err).Msg("viper error read config")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Error().Err(err).Msg("viper error unmarshal config")
	}

	return
}

func SetDefaultConfig(v *viper.Viper) {
	v.SetDefault("DATABASE.TIMEOUT.PING", "1s")
	v.SetDefault("DATABASE.TIMEOUT.READ", "5s")
	v.SetDefault("DATABASE.TIMEOUT.WRITE", "5s")

	v.SetDefault("ASYNQ.PROCESS_TIMEOUT", "30s")
	v.SetDefault("ASYNQ.MAX_RETRY", 5)

}

type EnvironmentVariable struct {
	App struct {
		Host  string `mapstucture:"HOST"`
		Port  int    `mapstructure:"PORT"`
		Debug bool   `mapstructure:"DEBUG"`
	} `mapstructure:"APP"`
	Token struct {
		SecretKey string `mapstructure:"SECRET_KEY"`
	} `mapstructure:"TOKEN"`
	Database struct {
		Postgres struct {
			UseMigration   bool   `mapstructure:"USE_MIGRATION"`
			Scheme         string `mapstructure:"SCHEME"`
			Host           string `mapstructure:"HOST"`
			Port           string `mapstructure:"PORT"`
			User           string `mapstructure:"USER"`
			Password       string `mapstructure:"PASSWORD"`
			Name           string `mapstructure:"NAME"`
			MaxConnections int    `mapstructure:"MAX_CONNECTIONS"`
			MaxIdleTime    int    `mapstructure:"MAX_IDLE_TIME"`
		} `mapstructure:"POSTGRES"`
	} `mapstructure:"DATABASE"`
	Swagger struct {
		Host     string `mapstructure:"HOST"`
		Protocol string `mapstructure:"PROTOCOL"`
	} `mapstructure:"SWAGGER"`
}

func (e *EnvironmentVariable) GetDBDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", e.Database.Postgres.Host, e.Database.Postgres.Port, e.Database.Postgres.User, e.Database.Postgres.Password, e.Database.Postgres.Name)
}

func (e *EnvironmentVariable) GetDBUrl() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", e.Database.Postgres.Scheme, e.Database.Postgres.User, e.Database.Postgres.Password, e.Database.Postgres.Host, e.Database.Postgres.Port, e.Database.Postgres.Name)
}
