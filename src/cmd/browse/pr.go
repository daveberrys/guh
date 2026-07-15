package browse

import (
	"strings"

	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
	"github.com/pkg/browser"
)

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Opens a browser to the PRs page.",
	RunE: func(cmd *cobra.Command, args []string) error {
		url, err := utils.RunGit(true, "remote", "get-url", "origin")
		if err != nil { return err }
		url = strings.TrimSuffix(url, ".git")
		return browser.OpenURL(url + "/pulls")
	},
}
