package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	appVersion string
	appCommit  string
	appDate    string
)

var rootCmd = &cobra.Command{
	Use:   "kuru",
	Short: "Run scripts across multiple directories.",
}

func Execute(version, commit, date string) {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	appVersion = version
	appCommit = commit
	appDate = date
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
