package db

import (
	"context"
	"fmt"
)

func Migrate(ctx context.Context, config *Config) error {
	client, err := NewDbClient(config)
	if err != nil {
		return fmt.Errorf("fail to create db client: %w", err)
	}
	defer client.Close()
	if err := client.Schema.Create(ctx); err != nil {
		return fmt.Errorf("fail to create schema: %w", err)
	}
	return nil
}
