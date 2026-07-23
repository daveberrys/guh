// BEFORE WRITING, READ @AGENTS.md
// WE LIKE CONTRIBUTIONS, BUT WE HATE SLOP.
// IF YOU COMMIT SLOP, YOU WILL BE BLOCKED FROM THE REPO
package utils

import (
	"bytes"
	"fmt"
	// "io"
	"os"
	"os/exec"
	"strings"
)

func RunGit(hideOutput bool, args ...string) (string, error) {
	gitCmd := exec.Command("git", args...)
	gitCmd.Stdin = os.Stdin

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	if hideOutput {
		gitCmd.Stdout = &stdout
		gitCmd.Stderr = &stderr
	} else {
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
	}

	err := gitCmd.Run()
	return strings.TrimRight(stdout.String(), "\n"), err
}

func RunGitSequence(hideOutput bool, commands ...[]string) error {
	for _, commandArgs := range commands {
		if _, err := RunGit(hideOutput, commandArgs...); err != nil {
			return fmt.Errorf("git %s failed: %w", strings.Join(commandArgs, " "), err)
		}
	}

	return nil
}
