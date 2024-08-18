package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/syoi-org/judge.syoi.org/logger"
	"go.uber.org/fx"
)

type ServerConfig struct {
	fx.Out
	Logger *logger.Config `mapstructure:"logger" yaml:"logger" validate:"required"`
}

func (a *ServerConfig) Validate() error {
	validate := validator.New()
	logger.RegisterLogLevelValidation(validate)
	return validate.Struct(a)
}

func NewServerConfig() (ServerConfig, error) {
	var config ServerConfig
	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("fail to unmarshal config: %w", err)
	}
	if err := config.Validate(); err != nil {
		return config, fmt.Errorf("fail to validate config: %w", err)
	}
	return config, nil
}

var ServerModule = fx.Module("config", fx.Provide(NewServerConfig))
