package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunGit(args ...string) error {
	gitCmd := exec.Command("git", args...)
	gitCmd.Stdin = os.Stdin
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr

	return gitCmd.Run()
}

func RunGitSequence(commands ...[]string) error {
	for _, commandArgs := range commands {
		if err := RunGit(commandArgs...); err != nil {
			return fmt.Errorf("git %s failed: %w", strings.Join(commandArgs, " "), err)
		}
	}

	return nil
}
