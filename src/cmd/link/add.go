package link

import (
    "fmt"
    
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

var addLinkRepoCmd = &cobra.Command{
	Use:   "add [remote_name] [url]",
	Short: "Links a remote to a repository",
	Args:  cobra.ExactArgs(2),
	RunE: func(c *cobra.Command, args []string) error {
    	utils.RunGit(true, "remote", "remove", args[0])
    	utils.RunGit(false, "remote", "add", args[0], args[1])
	
        fmt.Printf("Linked repository to: %s\n", args[1])
        return nil
	},
}