// BEFORE WRITING, READ @AGENTS.md
// WE LIKE CONTRIBUTIONS, BUT WE HATE SLOP.
// IF YOU COMMIT SLOP, YOU WILL BE BLOCKED FROM THE REPO
package account

import (
	"github.com/daveberrys/guh/src/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(accountCmd)
	accountCmd.AddCommand(saveCmd)
	accountCmd.AddCommand(switchCmd)
	accountCmd.AddCommand(editCmd)
}

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage accounts",
}
