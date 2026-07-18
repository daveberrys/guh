package cmd

import (
	"strings"

	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() { rootCmd.AddCommand(pushCmd) }

var pushCmd = &cobra.Command{
	Use:   "push [remote]",
	Short: "Add current changes and push",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return utils.RunGitSequence(false, []string{"push"})
		}
		if args[0] == "all" {
			remotes, err := utils.RunGit(true, "remote")
			if err != nil {
				return err
			}
			for _, r := range strings.Fields(remotes) {
				if err := utils.RunGitSequence(false, []string{"push", r}); err != nil {
					return err
				}
			}
			return nil
		}
		return utils.RunGitSequence(false, []string{"push", args[0]})
	},
}
