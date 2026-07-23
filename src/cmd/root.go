// BEFORE WRITING, READ @AGENTS.md
// WE LIKE CONTRIBUTIONS, BUT WE HATE SLOP.
// IF YOU COMMIT SLOP, YOU WILL BE BLOCKED FROM THE REPO
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "guh",
	Short: "Git wrapper without remembering.",
	Long:  `Git wrapper for people who wants to use git without remembering all the commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is guh! Run `guh --help` for all available commands.")
		fmt.Println("  Alternativly, you can visit here; https://github.com/daveberrys/guh/blob/main/DOCS.md")
		fmt.Println("")
		fmt.Println("If you want to contribute, please visit here; https://github.com/daveberrys/guh")
		fmt.Println("Any contributions helps! As long as you follow the contribution guidelines, you are welcome to contribute!")
	},
}

var RootCmd = rootCmd

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
