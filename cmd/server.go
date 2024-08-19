package cmd

import (
	"github.com/spf13/cobra"
	"github.com/syoi-org/judy/app"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	Long:  `This command is used to start server.`,
	Run: func(cmd *cobra.Command, args []string) {
		serverApp := app.NewServer()
		serverApp.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
