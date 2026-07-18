package createcmd

import (
	"fmt"

	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

var linkRepoCmd = &cobra.Command{
	Use:   "repo [url]",
	Short: "Link a repository",
	Args:  cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		switch args[0] {
		    case "":
		        return fmt.Errorf("repository URL cannot be empty")
		    case "what":
		        utils.RunGit(false, "remote", "get-url", "origin")
		        return nil
		    default:
				utils.RunGit(true, "remote", "remove", "origin")
				utils.RunGit(false, "remote", "add", "origin", args[0])
				
		        fmt.Printf("Linked repository to: %s\n", args[0])
		        return nil
		}
	},
}