package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pullCmd)
}

var pullCmd = &cobra.Command{
	Use:   "pull [source]",
	Short: "Pull updates from a source",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		source := "default"
		if len(args) > 0 {
			source = args[0]
		}
		fmt.Printf("Pulling data from: %s...\n", source)
	},
}
