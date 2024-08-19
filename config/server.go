package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/syoi-org/judy/db"
	"github.com/syoi-org/judy/logger"
	"github.com/syoi-org/judy/transport"
	"go.uber.org/fx"
)

type ServerConfig struct {
	fx.Out
	Logger *logger.Config    `mapstructure:"logger" yaml:"logger" validate:"required"`
	Http   *transport.Config `mapstructure:"http" yaml:"http" validate:"required"`
	Db     *db.Config        `mapstructure:"db" yaml:"db" validate:"required"`
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
