package cmd

import (
	"fmt"
	"os/exec"
	"os"
	
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
    
    out, err := exec.Command("git", "fetch").CombinedOutput()
    if err != nil { fmt.Fprintf(os.Stderr, "error: %s\n", err) }

    out, err = exec.Command("git", "pull").CombinedOutput()
    if err != nil { fmt.Fprintf(os.Stderr, "error: %s\n", err) }
    
    fmt.Print(string(out))
},

}
