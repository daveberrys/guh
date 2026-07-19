package logs

import (
	"github.com/daveberrys/guh/src/cmd"
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() { cmd.RootCmd.AddCommand(logCmd) }

var logCmd = &cobra.Command{
	Use:   "logs [commits]",
	Short: "Show commits logs.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cobraCmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return utils.RunGitSequence(false, []string{"log"})
		}
		return utils.RunGitSequence(false, []string{"log", "-" + args[0]})
	},
}
