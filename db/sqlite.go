package db

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/syoi-org/judge.syoi.org/ent"
)

type SQLiteConfig struct {
	Dsn string `mapstructure:"dsn" yaml:"dsn" validate:"required"`
}

func NewSQLiteClient(config *SQLiteConfig) (*ent.Client, error) {
	client, err := ent.Open("sqlite3", config.Dsn)
	if err != nil {
		return nil, fmt.Errorf("fail to open sqlite client: %w", err)
	}
	return client, nil
}

func NewInMemoryClient() (*ent.Client, error) {
	client, err := ent.Open("sqlite3", ":memory:?_fk=1")
	if err != nil {
		return nil, fmt.Errorf("fail to open sqlite client: %w", err)
	}
	return client, nil
}
