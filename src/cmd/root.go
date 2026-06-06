package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "guh",
	Short: "Guh is a modular CLI tool",
	Long:  `A comprehensive CLI tool built with Go and Cobra to demonstrate multi-command structures.`,
	Run: func(cmd *cobra.Command, args []string) {
		// This runs if no subcommand is provided
		fmt.Println("Welcome to guh! Use --help to see available commands.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
