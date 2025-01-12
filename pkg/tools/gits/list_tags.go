package gits

import (
	"fmt"
	"os/exec"
)

// ListTags lists all tags in the repository and prints them to stdout.
// It runs the "git tag" command with no arguments.
func ListTags() error {
	// Prepare the command
	cmd := exec.Command("git", "tag")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run git tag: %w", err)
	}

	// Print the output to stdout
	fmt.Println(string(output))
	return nil
}