package db

import (
	"context"
	"fmt"

	"github.com/syoi-org/judy/ent"
	"go.uber.org/fx"
)

func NewDbClient(config *Config) (*ent.Client, error) {
	switch config.Type {
	case "sqlite":
		return NewSQLiteClient(config.SQLite)
	case "memory":
		return NewInMemoryClient()
	default:
		return nil, fmt.Errorf("unknown db type: %s", config.Type)
	}
}

func runDbClient(lifecycle fx.Lifecycle, config *Config, client *ent.Client) error {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if config.AutoMigrate {
				return client.Schema.Create(ctx)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})
	return nil
}
