package link

import (
    "fmt"
    
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

var removeLinkRepoCmd = &cobra.Command{
    Use:   "remove [remote_name]",
	Short: "Removes a remote from a repository",
	Args:  cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
       	utils.RunGit(true, "remote", "remove", args[0])
	
        fmt.Printf("Removed remote: %s\n", args[0])
        return nil
	},
}
