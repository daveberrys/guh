package switchcmd

import (
	"github.com/daveberrys/guh/src/utils"
)

func performBranchSwitch(branchName string) error {
	return utils.RunGitSequence(false, []string{"switch", branchName})
}
