package ui

import (
	"github.com/duladissa/architerm/internal/theme"
	"github.com/charmbracelet/lipgloss"
)

// GetColors returns colors from the current theme
func GetColors() (
	primary, secondary, accent, warning, errorColor, muted,
	text, dim, background, border, highlight,
	prompt, command, separator, exitOK, exitFail, duration, output lipgloss.Color,
) {
	t := theme.CurrentTheme
	return t.GetPrimary(), t.GetSecondary(), t.GetAccent(), t.GetWarning(),
		t.GetError(), t.GetMuted(), t.GetForeground(), t.GetMuted(),
		t.GetBackground(), t.GetBorder(), t.GetSuggestionMatch(),
		t.GetPrompt(), t.GetCommand(), t.GetSeparator(), t.GetSuccess(),
		t.GetError(), t.GetSecondary(), t.GetOutput()
}

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

// DefaultStyles returns the default application styles using the current theme
func DefaultStyles() *Styles {
	t := theme.CurrentTheme
	s := &Styles{}

	// Base border style
	baseBorder := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(t.GetBorder())

	// App container
	s.App = lipgloss.NewStyle().Padding(0)

	// Header
	s.Header = lipgloss.NewStyle().
		Foreground(t.GetForeground()).
		Background(t.GetPrimary()).
		Bold(true).
		Padding(0, 2).
		MarginBottom(0)

	s.HeaderTitle = lipgloss.NewStyle().Foreground(t.GetForeground()).Bold(true)
	s.HeaderHelp = lipgloss.NewStyle().Foreground(t.GetMuted()).Italic(true)

	// Input panel
	s.InputPanel = baseBorder.Copy().BorderForeground(t.GetPrimary()).Padding(0, 1)
	s.InputPanelTitle = lipgloss.NewStyle().Foreground(t.GetPrimary()).Bold(true)
	s.InputPrompt = lipgloss.NewStyle().Foreground(t.GetAccent()).Bold(true)
	s.InputText = lipgloss.NewStyle().Foreground(t.GetForeground())
	s.InputGhost = lipgloss.NewStyle().Foreground(t.GetMuted()).Italic(true)
	s.InputCursor = lipgloss.NewStyle().Foreground(t.GetForeground()).Background(t.GetSuggestionMatch())

	// Suggestions panel
	s.SuggestionsPanel = baseBorder.Copy().BorderForeground(t.GetSecondary()).Padding(0, 1)
	s.SuggestionsPanelTitle = lipgloss.NewStyle().Foreground(t.GetSecondary()).Bold(true)
	s.SuggestionItem = lipgloss.NewStyle().Foreground(t.GetForeground()).Padding(0, 1)
	s.SuggestionSelected = lipgloss.NewStyle().Foreground(t.GetForeground()).Background(t.GetSuggestionMatch()).Bold(true).Padding(0, 1)
	s.SuggestionCommand = lipgloss.NewStyle().Foreground(t.GetForeground())
	s.SuggestionDesc = lipgloss.NewStyle().Foreground(t.GetMuted()).Italic(true)
	s.SuggestionCategory = lipgloss.NewStyle().Foreground(t.GetSecondary()).Bold(true)

	// Output panel
	s.OutputPanel = baseBorder.Copy().BorderForeground(t.GetAccent()).Padding(0, 1)
	s.OutputPanelTitle = lipgloss.NewStyle().Foreground(t.GetAccent()).Bold(true)
	s.OutputText = lipgloss.NewStyle().Foreground(t.GetOutput())
	s.OutputError = lipgloss.NewStyle().Foreground(t.GetError())
	s.OutputSuccess = lipgloss.NewStyle().Foreground(t.GetAccent())
	s.OutputPrompt = lipgloss.NewStyle().Foreground(t.GetPrompt()).Bold(true)
	s.OutputCommand = lipgloss.NewStyle().Foreground(t.GetCommand()).Bold(true)
	s.OutputSeparator = lipgloss.NewStyle().Foreground(t.GetSeparator())
	s.OutputDuration = lipgloss.NewStyle().Foreground(t.GetSecondary())
	s.OutputExitOK = lipgloss.NewStyle().Foreground(t.GetSuccess()).Bold(true)
	s.OutputExitFail = lipgloss.NewStyle().Foreground(t.GetError()).Bold(true)

	// Status bar
	s.StatusBar = lipgloss.NewStyle().Foreground(t.GetMuted()).Padding(0, 1)
	s.StatusText = lipgloss.NewStyle().Foreground(t.GetMuted())
	s.StatusKeyHint = lipgloss.NewStyle().Foreground(t.GetSecondary()).Bold(true)

	return s
}

// PanelTitle creates a styled panel title
func (s *Styles) PanelTitle(title string, style lipgloss.Style) string {
	return style.Render(" " + title + " ")
}
