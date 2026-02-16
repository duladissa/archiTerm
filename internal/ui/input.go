package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// InputPanel represents the command input area
type InputPanel struct {
	Value     string
	CursorPos int
	GhostText string
	Width     int
	Focused   bool
	styles    *Styles
}

// NewInputPanel creates a new input panel
func NewInputPanel(styles *Styles) *InputPanel {
	return &InputPanel{
		Value:     "",
		CursorPos: 0,
		GhostText: "",
		Width:     80,
		Focused:   true,
		styles:    styles,
	}
}

// SetValue sets the input value
func (p *InputPanel) SetValue(value string) {
	p.Value = value
	p.CursorPos = len(value)
}

// InsertChar inserts a character at cursor position
func (p *InputPanel) InsertChar(ch rune) {
	if p.CursorPos >= len(p.Value) {
		p.Value += string(ch)
	} else {
		p.Value = p.Value[:p.CursorPos] + string(ch) + p.Value[p.CursorPos:]
	}
	p.CursorPos++
}

// DeleteChar deletes character before cursor (backspace)
func (p *InputPanel) DeleteChar() {
	if p.CursorPos > 0 && len(p.Value) > 0 {
		p.Value = p.Value[:p.CursorPos-1] + p.Value[p.CursorPos:]
		p.CursorPos--
	}
}

// DeleteCharForward deletes character at cursor (delete key)
func (p *InputPanel) DeleteCharForward() {
	if p.CursorPos < len(p.Value) {
		p.Value = p.Value[:p.CursorPos] + p.Value[p.CursorPos+1:]
	}
}

// MoveCursorLeft moves cursor left
func (p *InputPanel) MoveCursorLeft() {
	if p.CursorPos > 0 {
		p.CursorPos--
	}
}

// MoveCursorRight moves cursor right
func (p *InputPanel) MoveCursorRight() {
	if p.CursorPos < len(p.Value) {
		p.CursorPos++
	}
}

// MoveCursorStart moves cursor to start
func (p *InputPanel) MoveCursorStart() {
	p.CursorPos = 0
}

// MoveCursorEnd moves cursor to end
func (p *InputPanel) MoveCursorEnd() {
	p.CursorPos = len(p.Value)
}

// Clear clears the input
func (p *InputPanel) Clear() {
	p.Value = ""
	p.CursorPos = 0
	p.GhostText = ""
}

// AcceptGhostText accepts the ghost text completion
func (p *InputPanel) AcceptGhostText() {
	if p.GhostText != "" {
		p.Value += p.GhostText
		p.CursorPos = len(p.Value)
		p.GhostText = ""
	}
}

// SetGhostText sets the ghost text for autocomplete
func (p *InputPanel) SetGhostText(ghost string) {
	p.GhostText = ghost
}

// SetWidth sets the panel width
func (p *InputPanel) SetWidth(width int) {
	p.Width = width
}

// SetStyles updates the styles for the input panel
func (p *InputPanel) SetStyles(styles *Styles) {
	p.styles = styles
}

// View renders the input panel
func (p *InputPanel) View() string {
	// Build input line with cursor
	prompt := p.styles.InputPrompt.Render("> ")
	
	var inputLine string
	if p.Focused {
		// Show cursor
		beforeCursor := p.Value[:p.CursorPos]
		afterCursor := ""
		cursorChar := " "
		
		if p.CursorPos < len(p.Value) {
			cursorChar = string(p.Value[p.CursorPos])
			afterCursor = p.Value[p.CursorPos+1:]
		}
		
		inputLine = p.styles.InputText.Render(beforeCursor) +
			p.styles.InputCursor.Render(cursorChar) +
			p.styles.InputText.Render(afterCursor)
		
		// Add ghost text after the value
		if p.GhostText != "" && p.CursorPos == len(p.Value) {
			inputLine += p.styles.InputGhost.Render(p.GhostText)
		}
	} else {
		inputLine = p.styles.InputText.Render(p.Value)
	}

	content := prompt + inputLine

	// Add padding to fill width
	contentWidth := lipgloss.Width(content)
	if contentWidth < p.Width-6 {
		content += strings.Repeat(" ", p.Width-6-contentWidth)
	}

	// Render panel with title in border
	panel := p.styles.InputPanel.
		Width(p.Width - 2).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true).
		Render(p.styles.InputPanelTitle.Render("⌨ Command") + "\n" + content)

	return panel
}
