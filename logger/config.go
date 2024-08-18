package logger

import (
	"github.com/go-playground/validator/v10"
)

type Config struct {
	Path  string `mapstructure:"path" yaml:"path" validate:"required"`
	Level struct {
		File    LogLevel `mapstructure:"file" yaml:"file" validate:"required,loglevel"`
		Console LogLevel `mapstructure:"console" yaml:"console" validate:"required,loglevel"`
	} `mapstructure:"level" yaml:"level" validate:"required"`
}

func RegisterLogLevelValidation(validate *validator.Validate) (*validator.Validate, error) {
	if err := validate.RegisterValidation("loglevel", validateLogLevel); err != nil {
		return nil, err
	}
	return validate, nil
}

func validateLogLevel(fieldLevel validator.FieldLevel) bool {
	logLevel := fieldLevel.Field().String()
	_, ok := logLevelMap[LogLevel(logLevel)]
	return ok
}
