package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/syoi-org/judy/db"
)

type MigrateConfig struct {
	Db db.Config `mapstructure:"db" yaml:"db" validate:"required"`
}

func (a *MigrateConfig) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}

func NewMigrateConfig() (MigrateConfig, error) {
	var config MigrateConfig
	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("fail to unmarshal config: %w", err)
	}
	if err := config.Validate(); err != nil {
		return config, fmt.Errorf("fail to validate config: %w", err)
	}
	return config, nil
}
