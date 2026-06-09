package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use:   "push [destination]",
	Short: "Push changes to a destination",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dest := "main"
		if len(args) > 0 { dest = args[0] }
		fmt.Printf("Pushing changes to: %s...\n", dest)
		
		out, err := exec.Command("git", "push", "origin", dest).CombinedOutput()
		if err != nil { fmt.Fprintf(os.Stderr, "error: %s\n", err) }
		fmt.Print(string(out))
	},
}
