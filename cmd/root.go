package cmd

import (
	"fmt"
	"os"

	"github.com/architerm/architerm/internal/app"
	"github.com/spf13/cobra"
)

var (
	configPath string
	version    = "0.1.0"
)

var rootCmd = &cobra.Command{
	Use:   "architerm",
	Short: "archiTerm - A smart terminal for architects and developers",
	Long: `archiTerm is a lightweight TUI application that provides smart autocomplete
for common infrastructure commands like Docker, Kubernetes, gcloud, Azure, curl, and git.

Features:
  • Gmail-style autocomplete with Tab completion
  • Three-panel UI (input, suggestions, output)
  • Support for custom commands via YAML/JSON config
  • Cross-platform (Windows, Linux, macOS)`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := app.Run(configPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("archiTerm v%s\n", version)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "path to custom config file (YAML or JSON)")
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
