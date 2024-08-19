package cmd

import (
	"github.com/spf13/cobra"
	"github.com/syoi-org/judy/app"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start worker",
	Long:  `This command is used to start worker.`,
	Run: func(cmd *cobra.Command, args []string) {
		workerApp := app.NewWorker()
		workerApp.Run()
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
