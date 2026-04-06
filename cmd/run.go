package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a script in multiple directories",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
