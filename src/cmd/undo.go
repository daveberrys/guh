package cmd
import (
	"github.com/spf13/cobra"
	"github.com/daveberrys/guh/src/utils"
)

func init() { RootCmd.AddCommand(undoCmd) }

var undoCmd = &cobra.Command{
	Use:   "undo [commits] [flavour]",
	Short: "Undo the last LOCAL commit. Flavour can be hard, mixed, or soft. It's optional.",
	Args: cobra.MinimumNArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		var flavour string
		switch args[1] {
		    case "hard":
			    flavour = "--hard"
		    case "mixed":
			    flavour = "--mixed"
		    default:
			    flavour = "--soft"
		}

		utils.RunGitSequence(false, []string{"reset", flavour, "HEAD~" + args[0]})
		return nil
	},
}
