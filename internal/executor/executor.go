package executor

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Result represents the result of a command execution
type Result struct {
	Command  string
	Output   string
	Error    string
	ExitCode int
	Duration time.Duration
}

// Executor handles command execution
type Executor struct {
	mu         sync.Mutex
	cancelFunc context.CancelFunc
	isRunning  bool
}

// NewExecutor creates a new command executor
func NewExecutor() *Executor {
	return &Executor{}
}

// Execute runs a command and returns the result
func (e *Executor) Execute(command string) *Result {
	e.mu.Lock()
	ctx, cancel := context.WithCancel(context.Background())
	e.cancelFunc = cancel
	e.isRunning = true
	e.mu.Unlock()

	defer func() {
		e.mu.Lock()
		e.isRunning = false
		e.cancelFunc = nil
		e.mu.Unlock()
	}()

	startTime := time.Now()
	result := &Result{
		Command: command,
	}

	// Determine shell based on OS
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.CommandContext(ctx, "cmd", "/C", command)
	} else {
		cmd = exec.CommandContext(ctx, "sh", "-c", command)
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	result.Duration = time.Since(startTime)

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitErr.ExitCode()
		} else {
			result.ExitCode = 1
		}
		result.Error = stderr.String()
		if result.Error == "" {
			result.Error = err.Error()
		}
	} else {
		result.ExitCode = 0
	}

	result.Output = stdout.String()

	// If stdout is empty but we have stderr output, include it
	if result.Output == "" && result.Error != "" {
		result.Output = result.Error
	}

	return result
}

// ExecuteAsync runs a command asynchronously and returns results via channel
func (e *Executor) ExecuteAsync(command string) <-chan *Result {
	resultChan := make(chan *Result, 1)
	go func() {
		result := e.Execute(command)
		resultChan <- result
		close(resultChan)
	}()
	return resultChan
}

// Cancel cancels the currently running command
func (e *Executor) Cancel() {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.cancelFunc != nil {
		e.cancelFunc()
	}
}

// IsRunning returns true if a command is currently executing
func (e *Executor) IsRunning() bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.isRunning
}

// CommandNotFoundInfo contains information about missing commands
type CommandNotFoundInfo struct {
	Command     string
	InstallHint string
}

// knownCommands maps commands to their installation hints
var knownCommands = map[string]CommandNotFoundInfo{
	"gcloud": {
		Command:     "gcloud",
		InstallHint: "Install Google Cloud SDK: https://cloud.google.com/sdk/docs/install",
	},
	"kubectl": {
		Command:     "kubectl",
		InstallHint: "Install kubectl: https://kubernetes.io/docs/tasks/tools/",
	},
	"docker": {
		Command:     "docker",
		InstallHint: "Install Docker: https://docs.docker.com/get-docker/",
	},
	"docker-compose": {
		Command:     "docker-compose",
		InstallHint: "Install Docker Compose: https://docs.docker.com/compose/install/",
	},
	"az": {
		Command:     "az",
		InstallHint: "Install Azure CLI: https://docs.microsoft.com/en-us/cli/azure/install-azure-cli",
	},
	"git": {
		Command:     "git",
		InstallHint: "Install Git: https://git-scm.com/downloads",
	},
	"curl": {
		Command:     "curl",
		InstallHint: "Install curl: apt install curl (Linux) | brew install curl (macOS)",
	},
	"terraform": {
		Command:     "terraform",
		InstallHint: "Install Terraform: https://www.terraform.io/downloads",
	},
	"helm": {
		Command:     "helm",
		InstallHint: "Install Helm: https://helm.sh/docs/intro/install/",
	},
	"aws": {
		Command:     "aws",
		InstallHint: "Install AWS CLI: https://aws.amazon.com/cli/",
	},
}

// isCommandNotFound checks if the error indicates a missing command
func isCommandNotFound(output string) bool {
	patterns := []string{
		"command not found",
		"not found",
		"not recognized as an internal or external command",
		"is not recognized",
		"No such file or directory",
		"not found in PATH",
		"executable file not found",
	}
	lowerOutput := strings.ToLower(output)
	for _, pattern := range patterns {
		if strings.Contains(lowerOutput, strings.ToLower(pattern)) {
			return true
		}
	}
	return false
}

// extractCommandName extracts the main command from a command string
func extractCommandName(command string) string {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return ""
	}
	return parts[0]
}

// FormatResult formats the result for display
func FormatResult(r *Result) string {
	var sb strings.Builder

	// Top separator for visual distinction between commands
	sb.WriteString("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	
	// Command line with prompt
	sb.WriteString(fmt.Sprintf("$ %s\n", r.Command))
	sb.WriteString("────────────────────────────────────────\n")

	// Check if command was not found
	if r.ExitCode != 0 && isCommandNotFound(r.Output) {
		cmdName := extractCommandName(r.Command)
		
		sb.WriteString("\n")
		sb.WriteString("⚠️  COMMAND NOT FOUND\n")
		sb.WriteString("────────────────────────────────────────\n")
		sb.WriteString(fmt.Sprintf("The command '%s' is not installed or not in PATH.\n\n", cmdName))
		
		// Check if we have install hints for this command
		if info, ok := knownCommands[cmdName]; ok {
			sb.WriteString("📦 HOW TO INSTALL:\n")
			sb.WriteString(fmt.Sprintf("   %s\n\n", info.InstallHint))
		} else {
			sb.WriteString("💡 SUGGESTIONS:\n")
			sb.WriteString(fmt.Sprintf("   • Check if '%s' is installed: which %s\n", cmdName, cmdName))
			sb.WriteString(fmt.Sprintf("   • Add '%s' to your PATH environment variable\n", cmdName))
			sb.WriteString("   • Install the required package using your package manager\n\n")
		}
		
		sb.WriteString("🔍 TROUBLESHOOTING:\n")
		sb.WriteString("   • Verify installation: which <command>\n")
		sb.WriteString("   • Check PATH: echo $PATH\n")
		sb.WriteString("   • Reload shell: source ~/.bashrc (or ~/.zshrc)\n")
		sb.WriteString("\n")
	} else if r.Output != "" {
		sb.WriteString(r.Output)
		if !strings.HasSuffix(r.Output, "\n") {
			sb.WriteString("\n")
		}
	}

	// Status line with exit code and duration
	if r.ExitCode != 0 {
		sb.WriteString(fmt.Sprintf("✗ [Exit code: %d] [Duration: %s]\n", r.ExitCode, r.Duration.Round(time.Millisecond)))
	} else {
		sb.WriteString(fmt.Sprintf("✓ [Exit code: 0] [Duration: %s]\n", r.Duration.Round(time.Millisecond)))
	}

	return sb.String()
}
