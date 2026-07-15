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
		utils.RunGitSequence(false, []string{"remote", "set-url", "origin", args[0]})
		fmt.Printf("Linked repository to: %s\n", args[0])
		return nil
	},
}