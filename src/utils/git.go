package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func RunGit(hideOutput bool, args ...string) error {
	gitCmd := exec.Command("git", args...)
	gitCmd.Stdin = os.Stdin
	if hideOutput {
		gitCmd.Stdout = io.Discard
		gitCmd.Stderr = io.Discard
	} else {
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
	}

	return gitCmd.Run()
}

func RunGitSequence(hideOutput bool, commands ...[]string) error {
	for _, commandArgs := range commands {
		if err := RunGit(hideOutput, commandArgs...); err != nil {
			return fmt.Errorf("git %s failed: %w", strings.Join(commandArgs, " "), err)
		}
	}

	return nil
}
