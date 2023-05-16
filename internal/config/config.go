/*
Package config represents application's configurations.
*/
package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Config is a struct which holds typed config options.
type Config struct {
	LogLevel  string `mapstructure:"LogLevel" validate:"required"`
	ServeHost string `mapstructure:"ServeHost" validate:"required"`
	ServePort string `mapstructure:"ServePort" validate:"required"`
}

// New reads the environment variables. Then it marshals
// config options, validates them and ruterns config.
func New() (*Config, error) {
	v := viper.New()

	v.BindEnv("ServeHost", "SERVE_HOST")
	v.SetDefault("ServeHost", "0.0.0.0")

	v.BindEnv("ServePort", "SERVE_PORT")
	v.SetDefault("ServePort", "8080")

	v.BindEnv("LogLevel", "LOG_LEVEL")
	v.SetDefault("LogLevel", "info")

	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		return cfg, err
	}

	val := validator.New()
	if err := val.Struct(cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
