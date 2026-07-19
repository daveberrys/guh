package branch

import (
	"fmt"

	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename [name]",
	Short: "Rename the current branch",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		utils.RunGitSequence(false, []string{"branch", "-m", args[0]})
		fmt.Println("Successfully renamed branch to", args[0])
		return nil
	},
}
