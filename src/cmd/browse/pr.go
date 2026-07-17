package browse

import (
	"strings"
	"fmt"

	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
	"github.com/pkg/browser"
)

var prCmd = &cobra.Command{
	Use:   "pr [print]",
	Short: "Opens a browser to the PRs page. If 'print' is true, prints the URL instead.",
	RunE: func(cmd *cobra.Command, args []string) error {
		url, err := utils.RunGit(true, "remote", "get-url", "origin")
		if err != nil { return err }

		url = strings.TrimSuffix(url, ".git")
		if len(args) > 0 && args[0] == "print" {
			fmt.Println(url + "/pulls")
			return nil
		}

		return browser.OpenURL(url + "/pulls")
	},
}
