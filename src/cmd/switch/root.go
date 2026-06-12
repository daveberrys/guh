package switchcmd

import (
	"github.com/daveberrys/guh/src/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(switchCmd)
	switchCmd.AddCommand(switchBranchCmd)
	switchCmd.AddCommand(switchAccountCmd)
}

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch branches or accounts",
}

var switchBranchCmd = &cobra.Command{
	Use:   "branch [name]",
	Short: "Switch to a specific branch",
	Args:  cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		return performBranchSwitch(args[0])
	},
}

var switchAccountCmd = &cobra.Command{
	Use:   "account [username]",
	Short: "Switch to a specific user account",
	Args:  cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		return performAccountSwitch(args[0])
	},
}
