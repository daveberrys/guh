package diff

import (
	"github.com/daveberrys/guh/src/cmd"
	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() { cmd.RootCmd.AddCommand(diffCmd) }

var diffCmd = &cobra.Command{
	Use:   "diff [file:opt]",
	Short: "Show working tree changes",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cobraCmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			return utils.RunGitSequence(false, []string{"diff", args[0]})
		}
		return utils.RunGitSequence(false, []string{"diff"})
	},
}
