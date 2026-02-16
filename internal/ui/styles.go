package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// Colors - Modern terminal color palette with Unix/Linux style
var (
	ColorPrimary    = lipgloss.Color("#7C3AED") // Purple
	ColorSecondary  = lipgloss.Color("#06B6D4") // Cyan
	ColorAccent     = lipgloss.Color("#10B981") // Green
	ColorWarning    = lipgloss.Color("#F59E0B") // Amber
	ColorError      = lipgloss.Color("#EF4444") // Red
	ColorMuted      = lipgloss.Color("#6B7280") // Gray
	ColorText       = lipgloss.Color("#F9FAFB") // White
	ColorDim        = lipgloss.Color("#4B5563") // Dim gray
	ColorBackground = lipgloss.Color("#1F2937") // Dark background
	ColorBorder     = lipgloss.Color("#374151") // Border gray
	ColorHighlight  = lipgloss.Color("#3B82F6") // Blue highlight

	// Unix/Linux terminal colors
	ColorPrompt     = lipgloss.Color("#22C55E") // Bright green (like PS1 prompt)
	ColorCommand    = lipgloss.Color("#FBBF24") // Yellow/gold (command text)
	ColorSeparator  = lipgloss.Color("#8B5CF6") // Purple (separator lines)
	ColorExitOK     = lipgloss.Color("#22C55E") // Green (success)
	ColorExitFail   = lipgloss.Color("#EF4444") // Red (failure)
	ColorDuration   = lipgloss.Color("#60A5FA") // Light blue (timing info)
	ColorOutput     = lipgloss.Color("#E5E7EB") // Light gray (command output)
)

// Styles holds all the application styles
type Styles struct {
	// App container
	App lipgloss.Style

	// Header
	Header      lipgloss.Style
	HeaderTitle lipgloss.Style
	HeaderHelp  lipgloss.Style

	// Input panel
	InputPanel       lipgloss.Style
	InputPanelTitle  lipgloss.Style
	InputPrompt      lipgloss.Style
	InputText        lipgloss.Style
	InputGhost       lipgloss.Style
	InputCursor      lipgloss.Style

	// Suggestions panel
	SuggestionsPanel      lipgloss.Style
	SuggestionsPanelTitle lipgloss.Style
	SuggestionItem        lipgloss.Style
	SuggestionSelected    lipgloss.Style
	SuggestionCommand     lipgloss.Style
	SuggestionDesc        lipgloss.Style
	SuggestionCategory    lipgloss.Style

	// Output panel
	OutputPanel      lipgloss.Style
	OutputPanelTitle lipgloss.Style
	OutputText       lipgloss.Style
	OutputError      lipgloss.Style
	OutputSuccess    lipgloss.Style
	OutputPrompt     lipgloss.Style
	OutputCommand    lipgloss.Style
	OutputSeparator  lipgloss.Style
	OutputDuration   lipgloss.Style
	OutputExitOK     lipgloss.Style
	OutputExitFail   lipgloss.Style

	// Status bar
	StatusBar     lipgloss.Style
	StatusText    lipgloss.Style
	StatusKeyHint lipgloss.Style

	// General
	Border lipgloss.Style
}

// DefaultStyles returns the default application styles
func DefaultStyles() *Styles {
	s := &Styles{}

	// Base border style
	baseBorder := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(ColorBorder)

	// App container
	s.App = lipgloss.NewStyle().
		Padding(0)

	// Header
	s.Header = lipgloss.NewStyle().
		Foreground(ColorText).
		Background(ColorPrimary).
		Bold(true).
		Padding(0, 2).
		MarginBottom(0)

	s.HeaderTitle = lipgloss.NewStyle().
		Foreground(ColorText).
		Bold(true)

	s.HeaderHelp = lipgloss.NewStyle().
		Foreground(ColorMuted).
		Italic(true)

	// Input panel
	s.InputPanel = baseBorder.Copy().
		BorderForeground(ColorPrimary).
		Padding(0, 1)

	s.InputPanelTitle = lipgloss.NewStyle().
		Foreground(ColorPrimary).
		Bold(true)

	s.InputPrompt = lipgloss.NewStyle().
		Foreground(ColorAccent).
		Bold(true)

	s.InputText = lipgloss.NewStyle().
		Foreground(ColorText)

	s.InputGhost = lipgloss.NewStyle().
		Foreground(ColorDim).
		Italic(true)

	s.InputCursor = lipgloss.NewStyle().
		Foreground(ColorText).
		Background(ColorHighlight)

	// Suggestions panel
	s.SuggestionsPanel = baseBorder.Copy().
		BorderForeground(ColorSecondary).
		Padding(0, 1)

	s.SuggestionsPanelTitle = lipgloss.NewStyle().
		Foreground(ColorSecondary).
		Bold(true)

	s.SuggestionItem = lipgloss.NewStyle().
		Foreground(ColorText).
		Padding(0, 1)

	s.SuggestionSelected = lipgloss.NewStyle().
		Foreground(ColorText).
		Background(ColorHighlight).
		Bold(true).
		Padding(0, 1)

	s.SuggestionCommand = lipgloss.NewStyle().
		Foreground(ColorText)

	s.SuggestionDesc = lipgloss.NewStyle().
		Foreground(ColorMuted).
		Italic(true)

	s.SuggestionCategory = lipgloss.NewStyle().
		Foreground(ColorSecondary).
		Bold(true)

	// Output panel
	s.OutputPanel = baseBorder.Copy().
		BorderForeground(ColorAccent).
		Padding(0, 1)

	s.OutputPanelTitle = lipgloss.NewStyle().
		Foreground(ColorAccent).
		Bold(true)

	s.OutputText = lipgloss.NewStyle().
		Foreground(ColorOutput)

	s.OutputError = lipgloss.NewStyle().
		Foreground(ColorError)

	s.OutputSuccess = lipgloss.NewStyle().
		Foreground(ColorAccent)

	s.OutputPrompt = lipgloss.NewStyle().
		Foreground(ColorPrompt).
		Bold(true)

	s.OutputCommand = lipgloss.NewStyle().
		Foreground(ColorCommand).
		Bold(true)

	s.OutputSeparator = lipgloss.NewStyle().
		Foreground(ColorSeparator)

	s.OutputDuration = lipgloss.NewStyle().
		Foreground(ColorDuration)

	s.OutputExitOK = lipgloss.NewStyle().
		Foreground(ColorExitOK).
		Bold(true)

	s.OutputExitFail = lipgloss.NewStyle().
		Foreground(ColorExitFail).
		Bold(true)

	// Status bar
	s.StatusBar = lipgloss.NewStyle().
		Foreground(ColorMuted).
		Padding(0, 1)

	s.StatusText = lipgloss.NewStyle().
		Foreground(ColorMuted)

	s.StatusKeyHint = lipgloss.NewStyle().
		Foreground(ColorSecondary).
		Bold(true)

	return s
}

// PanelTitle creates a styled panel title
func (s *Styles) PanelTitle(title string, style lipgloss.Style) string {
	return style.Render(" " + title + " ")
}
