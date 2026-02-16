package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/duladissa/architerm/internal/app"
	"github.com/duladissa/architerm/internal/theme"
	"github.com/spf13/cobra"
)

var (
	configPath  string
	themeName   string
	version     = "0.1.0"
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
  • Multiple color themes (dark, light, dracula, nord, gruvbox)
  • Cross-platform (Windows, Linux, macOS)`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set theme before starting app
		if themeName != "" {
			theme.SetTheme(themeName)
		}
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

var themesCmd = &cobra.Command{
	Use:   "themes",
	Short: "List available themes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available themes:")
		fmt.Println()
		for _, t := range theme.GetAvailableThemes() {
			th := theme.GetTheme(t)
			fmt.Printf("  • %s\n", th.Name)
		}
		fmt.Println()
		fmt.Println("Usage: architerm --theme <name>")
		fmt.Println("Example: architerm --theme dracula")
	},
}

var themePreviewCmd = &cobra.Command{
	Use:   "theme-preview [name]",
	Short: "Preview a theme's colors",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := "dark"
		if len(args) > 0 {
			name = args[0]
		}
		t := theme.GetTheme(name)
		fmt.Printf("Theme: %s\n", t.Name)
		fmt.Println(strings.Repeat("─", 40))
		fmt.Printf("Primary:    %s\n", t.Colors.Primary)
		fmt.Printf("Secondary:  %s\n", t.Colors.Secondary)
		fmt.Printf("Accent:     %s\n", t.Colors.Accent)
		fmt.Printf("Background: %s\n", t.Colors.Background)
		fmt.Printf("Foreground: %s\n", t.Colors.Foreground)
		fmt.Printf("Prompt:     %s\n", t.Colors.Prompt)
		fmt.Printf("Command:    %s\n", t.Colors.Command)
		fmt.Printf("Success:    %s\n", t.Colors.Success)
		fmt.Printf("Error:      %s\n", t.Colors.Error)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "path to custom config file (YAML or JSON)")
	rootCmd.PersistentFlags().StringVarP(&themeName, "theme", "t", "dark", "color theme (dark, light, dracula, nord, gruvbox)")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(themesCmd)
	rootCmd.AddCommand(themePreviewCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
