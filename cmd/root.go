package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/syoi-org/judy/config"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "judy",
	Short: "Command line utility for running SYOI Online Judge",
	Long: `This is a command line utility for running SYOI Online Judge.
	
For more information, see https://github.com/syoi-org/judy`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		config.InitConfig(configFile)
	})
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "path to config file")
}
