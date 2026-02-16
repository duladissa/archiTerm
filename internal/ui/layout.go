package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Layout manages the three-panel layout (side-by-side)
// Left side: Input + Suggestions (stacked)
// Right side: Output
type Layout struct {
	Width  int
	Height int
	styles *Styles
}

// NewLayout creates a new layout manager
func NewLayout(styles *Styles) *Layout {
	return &Layout{
		Width:  80,
		Height: 24,
		styles: styles,
	}
}

// SetSize sets the terminal size
func (l *Layout) SetSize(width, height int) {
	l.Width = width
	l.Height = height
}

// GetLeftPanelWidth returns the width for the left panel (input + suggestions)
func (l *Layout) GetLeftPanelWidth() int {
	// Left panel takes 45% of width
	w := int(float64(l.Width) * 0.45)
	if w < 40 {
		w = 40
	}
	return w
}

// GetRightPanelWidth returns the width for the right panel (output)
func (l *Layout) GetRightPanelWidth() int {
	// Right panel takes remaining width
	return l.Width - l.GetLeftPanelWidth() - 3 // 3 for gap
}

// GetInputHeight returns the height for the input panel
func (l *Layout) GetInputHeight() int {
	return 4
}

// GetSuggestionsHeight returns the height for the suggestions panel
func (l *Layout) GetSuggestionsHeight() int {
	// Fill remaining height on left side, minus categories panel
	remaining := l.Height - l.GetInputHeight() - l.GetCategoriesHeight() - 6 // header + status + gaps
	if remaining < 6 {
		remaining = 6
	}
	return remaining
}

// GetCategoriesHeight returns the height for the categories panel
func (l *Layout) GetCategoriesHeight() int {
	return 4
}

// GetOutputHeight returns the height for the output panel
func (l *Layout) GetOutputHeight() int {
	// Output panel fills the right side
	return l.Height - 4 // header + status + gaps
}

// GetContentWidth returns the width for content panels (legacy, use specific widths)
func (l *Layout) GetContentWidth() int {
	return l.Width - 2
}

// RenderHeader renders the application header
func (l *Layout) RenderHeader() string {
	title := " 🏛️  archiTerm "
	help := " Tab: complete │ ↑↓: navigate │ Enter: execute │ Ctrl+C: exit "
	
	titleStyle := l.styles.Header.Copy().Width(l.Width)
	
	// Calculate spacing
	spacing := l.Width - lipgloss.Width(title) - lipgloss.Width(help)
	if spacing < 0 {
		spacing = 0
		help = ""
	}
	
	header := title + strings.Repeat(" ", spacing) + l.styles.HeaderHelp.Render(help)
	
	return titleStyle.Render(header)
}

// RenderStatusBar renders the status bar at the bottom
func (l *Layout) RenderStatusBar(status string) string {
	statusStyle := l.styles.StatusBar.Copy().Width(l.Width)
	
	keyHints := []string{
		l.styles.StatusKeyHint.Render("Tab") + " complete",
		l.styles.StatusKeyHint.Render("Enter") + " run",
		l.styles.StatusKeyHint.Render("Ctrl+Y") + " copy out",
		l.styles.StatusKeyHint.Render("Ctrl+B") + " copy cmd",
		l.styles.StatusKeyHint.Render("Ctrl+L") + " clear",
		l.styles.StatusKeyHint.Render("Ctrl+C") + " exit",
	}
	
	hintsStr := strings.Join(keyHints, " │ ")
	
	if status != "" {
		status = " " + status + " │ "
	}
	
	return statusStyle.Render(status + hintsStr)
}

// Render combines all panels into the final side-by-side layout
func (l *Layout) Render(header, input, suggestions, categories, output, statusBar string) string {
	// Stack input, suggestions, and categories vertically on the left
	leftPanel := lipgloss.JoinVertical(lipgloss.Left, input, suggestions, categories)
	
	// Join left and right panels horizontally
	mainContent := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, "  ", output)
	
	// Combine header, main content, and status bar
	return lipgloss.JoinVertical(lipgloss.Left, header, mainContent, statusBar)
}
