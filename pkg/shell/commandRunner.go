package shell

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"al.essio.dev/pkg/shellescape"
)

type CommandRunner struct {
	EnvContext EnvironmentContext
}

func NewCommandRunner() *CommandRunner {
	envContext, err := NewEnvironmentContext()
	if err != nil {
		panic(err)
	}
	return &CommandRunner{EnvContext: envContext}
}

func (c *CommandRunner) Run(command string, args ...string) (string, error) {
	var escapedArgs []string
	for _, arg := range args {
		escapedArgs = append(escapedArgs, shellescape.Quote(arg))
	}
	cmdStr := command + " " + strings.Join(escapedArgs, " ")

	return c.RunString(cmdStr)
}

func (c *CommandRunner) RunString(command string) (string, error) {
	shell := c.EnvContext.Shell

	if shell == "" {
		shell = "/bin/sh"
	}

	args := []string{
		"-c",
		command,
	}

	if c.EnvContext.IsWindowsStyleFlags {
		args = []string{
			"/c",
			command,
		}
	}

	cmd := exec.Command(shell, args...)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("failed to get stdout pipe: %v", err)
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return "", fmt.Errorf("failed to get stderr pipe: %v", err)
	}
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("failed to start command: %v", err)
	}

	outputBuilder := &strings.Builder{}
	stdoutScanner := bufio.NewScanner(stdoutPipe)
	stderrScanner := bufio.NewScanner(stderrPipe)

	stdoutDone := make(chan bool)
	stderrDone := make(chan bool)

	go func() {
		for stdoutScanner.Scan() {
			outputBuilder.WriteString(stdoutScanner.Text() + "\n")
		}
		stdoutDone <- true
	}()

	go func() {
		for stderrScanner.Scan() {
			outputBuilder.WriteString(stderrScanner.Text() + "\n")
		}
		stderrDone <- true
	}()

	<-stdoutDone
	<-stderrDone

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("failed to execute command: %v, output: %s", err, outputBuilder.String())
	}

	combinedOutput := outputBuilder.String()

	return combinedOutput, nil
}
