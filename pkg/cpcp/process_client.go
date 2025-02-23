package cpcp

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"sync"
)

// ProcessClient manages a single plugin process.
type ProcessClient struct {
	cmd      *exec.Cmd
	stdin    chan string
	stdout   chan string
	errCh    chan error
	exitCode chan int
	done     chan struct{}
	mu       sync.Mutex
	started  bool
	logger   *slog.Logger
}

// NewProcessClient initializes a PluginHost instance.
func NewProcessClient(logger *slog.Logger, pluginPath string, args ...string) *ProcessClient {
	return &ProcessClient{
		cmd:      exec.Command(pluginPath, args...),
		stdin:    make(chan string),
		stdout:   make(chan string),
		errCh:    make(chan error, 1),
		exitCode: make(chan int, 1),
		done:     make(chan struct{}),
		started:  false,
		logger:   logger,
	}
}

var _ DuplexClient = (*ProcessClient)(nil)

// Start launches the plugin process.
func (h *ProcessClient) Start() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.started {
		return fmt.Errorf("plugin already started")
	}

	stdinPipe, err := h.cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdin pipe: %w", err)
	}

	stdoutPipe, err := h.cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	stderrPipe, err := h.cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			h.logger.Debug(
				scanner.Text(),
				slog.String("source", "stderr"),
				slog.String("plugin", h.cmd.Path),
				slog.Any("args", h.cmd.Args),
			)
		}
	}()

	// Start reading stdout
	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			h.stdout <- scanner.Text()
		}
		close(h.stdout)
	}()

	// Start writing to stdin
	go func() {
		writer := bufio.NewWriter(stdinPipe)
		for line := range h.stdin {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				h.errCh <- err
				break
			}
			writer.Flush()
		}
	}()

	// Start process
	if err := h.cmd.Start(); err != nil {
		return fmt.Errorf("failed to start plugin: %w", err)
	}

	// Monitor process exit
	go func() {
		err := h.cmd.Wait()
		if err != nil {
			h.errCh <- err
		}
		if exitErr, ok := err.(*exec.ExitError); ok {
			h.exitCode <- exitErr.ExitCode()
		} else {
			h.exitCode <- 0 // Normal exit
		}
		h.done <- struct{}{}
		close(h.exitCode)
	}()

	h.started = true
	return nil
}

// Send sends a line to the plugin.
func (h *ProcessClient) Send(line string) {
	h.stdin <- line
}

// Receive returns the stdout channel.
func (h *ProcessClient) Receive() <-chan string {
	return h.stdout
}

// Errors returns the error channel.
func (h *ProcessClient) Errors() <-chan error {
	return h.errCh
}

// ExitCode returns the exit code channel.
func (h *ProcessClient) ExitCode() <-chan int {
	return h.exitCode
}

// Stop terminates the plugin process.
func (h *ProcessClient) Stop() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if !h.started {
		return fmt.Errorf("plugin is not running")
	}

	close(h.stdin)

	err := h.cmd.Process.Signal(os.Interrupt) //TODO: is this the correct signal?
	<-h.done                                  //TODO: add process kill timeout
	h.started = false
	return err
}
