package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "judgectl",
	Short: "Command line utility for running SYOI Online Judge",
	Long: `This is a command line utility for running SYOI Online Judge.
	
For more information, see https://github.com/syoi-org/judge.syoi.org`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
