package browse

import (
	"strings"
	"fmt"

	"github.com/daveberrys/guh/src/cmd"
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
	"github.com/pkg/browser"
)

func init() {
	cmd.RootCmd.AddCommand(createCmd)
	createCmd.AddCommand(prCmd)
	createCmd.AddCommand(issuesCmd)
}

var createCmd = &cobra.Command{
	Use:   "browse [print]",
	Short: "Opens a browser to the repository. If 'print' is true, prints the URL instead.",
	RunE: func(cmd *cobra.Command, args []string) error {
		url, err := utils.RunGit(true, "remote", "get-url", "origin")
		if err != nil { return err }

		url = strings.TrimSuffix(url, ".git")
		if len(args) > 0 && args[0] == "print" {
			fmt.Println(url)
			return nil
		}

		return browser.OpenURL(url)
	},
}