package db

type Config struct {
	Type        string        `mapstructure:"type" yaml:"type" validate:"required"`
	AutoMigrate bool          `mapstructure:"auto_migrate" yaml:"auto_migrate"`
	SQLite      *SQLiteConfig `mapstructure:"sqlite" yaml:"sqlite"`
}
