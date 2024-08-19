package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/syoi-org/judy/config"
	"github.com/syoi-org/judy/db"
	"go.uber.org/zap"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database schema",
	Long: `This command is used to migrate database schema. Note that currently migration
is not versioned. Please use carefully.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.NewMigrateConfig()
		if err != nil {
			zap.S().Fatalf("fail to read config: %v", err)
			return
		}
		if err := db.Migrate(cmd.Context(), &config.Db); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Migrate success")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
