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
		fmt.Println("Welcome to guh! Use --help to see available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
