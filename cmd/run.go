package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a script in multiple directories",
	Args:  cobra.ExactArgs(1),
	Run:   runCommand,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runCommand(cmd *cobra.Command, args []string) {
	command := prepareCommand(args)

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}

func prepareCommand(args []string) *exec.Cmd {
	script := args[0]
	if _, err := os.Stat(script); err == nil {
		return exec.Command(script)
	}
	return exec.Command("sh", "-c", script)
}
