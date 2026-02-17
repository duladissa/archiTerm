package theme

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
)

// Theme represents a color theme for the application
type Theme struct {
	Name        string      `json:"name"`
	Colors      ThemeColors `json:"colors"`
}

// ThemeColors contains all customizable colors
type ThemeColors struct {
	// Primary UI colors
	Primary     string `json:"primary"`
	Secondary   string `json:"secondary"`
	Accent      string `json:"accent"`
	Background  string `json:"background"`
	Foreground  string `json:"foreground"`
	Border      string `json:"border"`
	
	// Status colors
	Success     string `json:"success"`
	Warning     string `json:"warning"`
	Error       string `json:"error"`
	Muted       string `json:"muted"`
	
	// Output colors (Unix style)
	Prompt      string `json:"prompt"`
	Command     string `json:"command"`
	Output      string `json:"output"`
	Separator   string `json:"separator"`
	
	// Suggestion colors
	Suggestion      string `json:"suggestion"`
	SuggestionHint  string `json:"suggestion_hint"`
	SuggestionMatch string `json:"suggestion_match"`
	
	// Selection
	SelectionBg string `json:"selection_bg"`
	SelectionFg string `json:"selection_fg"`
}

// CurrentTheme is the active theme
var CurrentTheme = DarkTheme()

// DarkTheme returns the default dark theme
func DarkTheme() *Theme {
	return &Theme{
		Name: "dark",
		Colors: ThemeColors{
			// Primary UI colors
			Primary:     "#7C3AED", // Purple
			Secondary:   "#06B6D4", // Cyan
			Accent:      "#10B981", // Green
			Background:  "#1F2937", // Dark gray
			Foreground:  "#F9FAFB", // White
			Border:      "#374151", // Gray
			
			// Status colors
			Success:     "#10B981", // Green
			Warning:     "#F59E0B", // Amber
			Error:       "#EF4444", // Red
			Muted:       "#6B7280", // Gray
			
			// Output colors (Unix style)
			Prompt:      "#22C55E", // Bright green
			Command:     "#FBBF24", // Yellow/Gold
			Output:      "#D1D5DB", // Light gray
			Separator:   "#8B5CF6", // Purple
			
			// Suggestion colors
			Suggestion:      "#F9FAFB", // White
			SuggestionHint:  "#6B7280", // Muted
			SuggestionMatch: "#3B82F6", // Blue
			
			// Selection
			SelectionBg: "#3B82F6", // Blue
			SelectionFg: "#FFFFFF", // White
		},
	}
}

// DraculaTheme returns a Dracula-inspired theme
func DraculaTheme() *Theme {
	return &Theme{
		Name: "dracula",
		Colors: ThemeColors{
			// Primary UI colors
			Primary:     "#BD93F9", // Purple
			Secondary:   "#8BE9FD", // Cyan
			Accent:      "#50FA7B", // Green
			Background:  "#282A36", // Background
			Foreground:  "#F8F8F2", // Foreground
			Border:      "#44475A", // Current line
			
			// Status colors
			Success:     "#50FA7B", // Green
			Warning:     "#FFB86C", // Orange
			Error:       "#FF5555", // Red
			Muted:       "#6272A4", // Comment
			
			// Output colors
			Prompt:      "#50FA7B", // Green
			Command:     "#F1FA8C", // Yellow
			Output:      "#F8F8F2", // Foreground
			Separator:   "#BD93F9", // Purple
			
			// Suggestion colors
			Suggestion:      "#F8F8F2", // Foreground
			SuggestionHint:  "#6272A4", // Comment
			SuggestionMatch: "#8BE9FD", // Cyan
			
			// Selection
			SelectionBg: "#44475A", // Current line
			SelectionFg: "#F8F8F2", // Foreground
		},
	}
}

// NordTheme returns a Nord-inspired theme
func NordTheme() *Theme {
	return &Theme{
		Name: "nord",
		Colors: ThemeColors{
			// Primary UI colors
			Primary:     "#5E81AC", // Nord10
			Secondary:   "#88C0D0", // Nord8
			Accent:      "#A3BE8C", // Nord14
			Background:  "#2E3440", // Nord0
			Foreground:  "#ECEFF4", // Nord6
			Border:      "#4C566A", // Nord3
			
			// Status colors
			Success:     "#A3BE8C", // Nord14
			Warning:     "#EBCB8B", // Nord13
			Error:       "#BF616A", // Nord11
			Muted:       "#4C566A", // Nord3
			
			// Output colors
			Prompt:      "#A3BE8C", // Nord14
			Command:     "#EBCB8B", // Nord13
			Output:      "#D8DEE9", // Nord4
			Separator:   "#5E81AC", // Nord10
			
			// Suggestion colors
			Suggestion:      "#ECEFF4", // Nord6
			SuggestionHint:  "#4C566A", // Nord3
			SuggestionMatch: "#88C0D0", // Nord8
			
			// Selection
			SelectionBg: "#5E81AC", // Nord10
			SelectionFg: "#ECEFF4", // Nord6
		},
	}
}

// GruvboxTheme returns a Gruvbox dark theme
func GruvboxTheme() *Theme {
	return &Theme{
		Name: "gruvbox",
		Colors: ThemeColors{
			// Primary UI colors
			Primary:     "#D79921", // Yellow
			Secondary:   "#458588", // Blue
			Accent:      "#98971A", // Green
			Background:  "#282828", // bg0
			Foreground:  "#EBDBB2", // fg1
			Border:      "#504945", // bg2
			
			// Status colors
			Success:     "#98971A", // Green
			Warning:     "#D79921", // Yellow
			Error:       "#CC241D", // Red
			Muted:       "#928374", // gray
			
			// Output colors
			Prompt:      "#B8BB26", // Bright green
			Command:     "#FABD2F", // Bright yellow
			Output:      "#EBDBB2", // fg1
			Separator:   "#D3869B", // Purple
			
			// Suggestion colors
			Suggestion:      "#EBDBB2", // fg1
			SuggestionHint:  "#928374", // gray
			SuggestionMatch: "#83A598", // Bright blue
			
			// Selection
			SelectionBg: "#504945", // bg2
			SelectionFg: "#EBDBB2", // fg1
		},
	}
}

// GetTheme returns a theme by name
func GetTheme(name string) *Theme {
	switch name {
	case "dracula":
		return DraculaTheme()
	case "nord":
		return NordTheme()
	case "gruvbox":
		return GruvboxTheme()
	case "dark":
		fallthrough
	default:
		return DarkTheme()
	}
}

// GetAvailableThemes returns list of available theme names
func GetAvailableThemes() []string {
	return []string{"dark", "dracula", "nord", "gruvbox"}
}

// SetTheme sets the current theme by name
func SetTheme(name string) {
	CurrentTheme = GetTheme(name)
}

// CycleTheme cycles to the next available theme and returns its name
func CycleTheme() string {
	themes := GetAvailableThemes()
	currentName := CurrentTheme.Name
	
	// Find current theme index
	currentIndex := 0
	for i, name := range themes {
		if name == currentName {
			currentIndex = i
			break
		}
	}
	
	// Move to next theme (wrap around)
	nextIndex := (currentIndex + 1) % len(themes)
	nextTheme := themes[nextIndex]
	
	SetTheme(nextTheme)
	return nextTheme
}

// LoadThemeFromFile loads a custom theme from a JSON file
func LoadThemeFromFile(path string) (*Theme, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	
	var theme Theme
	if err := json.Unmarshal(data, &theme); err != nil {
		return nil, err
	}
	
	return &theme, nil
}

// SaveThemeToFile saves a theme to a JSON file
func SaveThemeToFile(theme *Theme, path string) error {
	data, err := json.MarshalIndent(theme, "", "  ")
	if err != nil {
		return err
	}
	
	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	
	return os.WriteFile(path, data, 0644)
}

// Color helper to convert theme color string to lipgloss.Color
func (t *Theme) Color(c string) lipgloss.Color {
	return lipgloss.Color(c)
}

// GetPrimary returns the primary color
func (t *Theme) GetPrimary() lipgloss.Color {
	return lipgloss.Color(t.Colors.Primary)
}

// GetSecondary returns the secondary color
func (t *Theme) GetSecondary() lipgloss.Color {
	return lipgloss.Color(t.Colors.Secondary)
}

// GetAccent returns the accent color
func (t *Theme) GetAccent() lipgloss.Color {
	return lipgloss.Color(t.Colors.Accent)
}

// GetBackground returns the background color
func (t *Theme) GetBackground() lipgloss.Color {
	return lipgloss.Color(t.Colors.Background)
}

// GetForeground returns the foreground color
func (t *Theme) GetForeground() lipgloss.Color {
	return lipgloss.Color(t.Colors.Foreground)
}

// GetBorder returns the border color
func (t *Theme) GetBorder() lipgloss.Color {
	return lipgloss.Color(t.Colors.Border)
}

// GetSuccess returns the success color
func (t *Theme) GetSuccess() lipgloss.Color {
	return lipgloss.Color(t.Colors.Success)
}

// GetWarning returns the warning color
func (t *Theme) GetWarning() lipgloss.Color {
	return lipgloss.Color(t.Colors.Warning)
}

// GetError returns the error color
func (t *Theme) GetError() lipgloss.Color {
	return lipgloss.Color(t.Colors.Error)
}

// GetMuted returns the muted color
func (t *Theme) GetMuted() lipgloss.Color {
	return lipgloss.Color(t.Colors.Muted)
}

// GetPrompt returns the prompt color
func (t *Theme) GetPrompt() lipgloss.Color {
	return lipgloss.Color(t.Colors.Prompt)
}

// GetCommand returns the command color
func (t *Theme) GetCommand() lipgloss.Color {
	return lipgloss.Color(t.Colors.Command)
}

// GetOutput returns the output color
func (t *Theme) GetOutput() lipgloss.Color {
	return lipgloss.Color(t.Colors.Output)
}

// GetSeparator returns the separator color
func (t *Theme) GetSeparator() lipgloss.Color {
	return lipgloss.Color(t.Colors.Separator)
}

// GetSuggestion returns the suggestion color
func (t *Theme) GetSuggestion() lipgloss.Color {
	return lipgloss.Color(t.Colors.Suggestion)
}

// GetSuggestionHint returns the suggestion hint color
func (t *Theme) GetSuggestionHint() lipgloss.Color {
	return lipgloss.Color(t.Colors.SuggestionHint)
}

// GetSuggestionMatch returns the suggestion match color
func (t *Theme) GetSuggestionMatch() lipgloss.Color {
	return lipgloss.Color(t.Colors.SuggestionMatch)
}

// GetSelectionBg returns the selection background color
func (t *Theme) GetSelectionBg() lipgloss.Color {
	return lipgloss.Color(t.Colors.SelectionBg)
}

// GetSelectionFg returns the selection foreground color
func (t *Theme) GetSelectionFg() lipgloss.Color {
	return lipgloss.Color(t.Colors.SelectionFg)
}
