package ui

import (
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/lipgloss"
)

// OutputEntry represents a single command output entry
type OutputEntry struct {
	Command  string
	Output   string
	FullText string // Complete formatted output for copying
}

// OutputPanel represents the command output area
type OutputPanel struct {
	Content       string
	Lines         []string
	Entries       []OutputEntry // Individual command outputs
	SelectedEntry int           // Currently selected entry for copying
	ScrollOffset  int
	Width         int
	Height        int
	styles        *Styles
	CopyMessage   string // Temporary message shown after copy
	
	// Mouse selection state
	IsSelecting     bool
	SelectionStart  int // Line number where selection started
	SelectionEnd    int // Line number where selection ended
	SelectedText    string
	LastClickTime   int64 // For double-click detection (unix nano)
}

// NewOutputPanel creates a new output panel
func NewOutputPanel(styles *Styles) *OutputPanel {
	return &OutputPanel{
		Content:       "",
		Lines:         make([]string, 0),
		Entries:       make([]OutputEntry, 0),
		SelectedEntry: -1,
		ScrollOffset:  0,
		Width:         80,
		Height:        15,
		styles:        styles,
		CopyMessage:   "",
	}
}

// SetContent sets the output content
func (p *OutputPanel) SetContent(content string) {
	p.Content = content
	p.Lines = strings.Split(content, "\n")
	// Auto-scroll to bottom
	p.ScrollToBottom()
}

// AppendContent appends content to the output
func (p *OutputPanel) AppendContent(content string) {
	if p.Content == "" {
		p.Content = content
	} else {
		p.Content += content
	}
	p.Lines = strings.Split(p.Content, "\n")
	p.ScrollToBottom()
}

// AddEntry adds a new command output entry (replaces previous output display)
func (p *OutputPanel) AddEntry(command, output, fullText string) {
	entry := OutputEntry{
		Command:  command,
		Output:   output,
		FullText: fullText,
	}
	// Keep entries for copy history
	p.Entries = append(p.Entries, entry)
	p.SelectedEntry = len(p.Entries) - 1 // Select the latest entry
	
	// Replace content with only the latest output
	p.Content = fullText
	p.Lines = strings.Split(p.Content, "\n")
	p.ScrollOffset = 0 // Reset scroll to top for new output
	p.CopyMessage = "" // Clear any previous copy message
}

// GetEntryCount returns the number of entries
func (p *OutputPanel) GetEntryCount() int {
	return len(p.Entries)
}

// SelectPreviousEntry selects the previous entry
func (p *OutputPanel) SelectPreviousEntry() {
	if len(p.Entries) > 0 && p.SelectedEntry > 0 {
		p.SelectedEntry--
		p.CopyMessage = ""
	}
}

// SelectNextEntry selects the next entry
func (p *OutputPanel) SelectNextEntry() {
	if len(p.Entries) > 0 && p.SelectedEntry < len(p.Entries)-1 {
		p.SelectedEntry++
		p.CopyMessage = ""
	}
}

// CopySelectedEntry copies the selected entry to clipboard
func (p *OutputPanel) CopySelectedEntry() error {
	if len(p.Entries) == 0 || p.SelectedEntry < 0 || p.SelectedEntry >= len(p.Entries) {
		return nil
	}
	entry := p.Entries[p.SelectedEntry]
	err := clipboard.WriteAll(entry.FullText)
	if err != nil {
		p.CopyMessage = "❌ Copy failed!"
		return err
	}
	p.CopyMessage = "✅ Copied to clipboard!"
	return nil
}

// CopyLastOutput copies the last command output to clipboard
func (p *OutputPanel) CopyLastOutput() error {
	if len(p.Entries) == 0 {
		return nil
	}
	entry := p.Entries[len(p.Entries)-1]
	err := clipboard.WriteAll(entry.Output)
	if err != nil {
		p.CopyMessage = "❌ Copy failed!"
		return err
	}
	p.CopyMessage = "✅ Output copied!"
	return nil
}

// CopyLastCommand copies the last command to clipboard
func (p *OutputPanel) CopyLastCommand() error {
	if len(p.Entries) == 0 {
		return nil
	}
	entry := p.Entries[len(p.Entries)-1]
	err := clipboard.WriteAll(entry.Command)
	if err != nil {
		p.CopyMessage = "❌ Copy failed!"
		return err
	}
	p.CopyMessage = "✅ Command copied!"
	return nil
}

// ClearCopyMessage clears the copy message
func (p *OutputPanel) ClearCopyMessage() {
	p.CopyMessage = ""
}

// StartSelection starts mouse text selection at a given line
// Returns true if this was a double-click (select all)
func (p *OutputPanel) StartSelection(lineY int) bool {
	now := time.Now().UnixNano()
	doubleClickThreshold := int64(500 * time.Millisecond) // 500ms threshold
	
	// Check for double-click
	if now-p.LastClickTime < doubleClickThreshold {
		// Double-click detected - select all
		p.SelectAll()
		p.LastClickTime = 0 // Reset to prevent triple-click
		return true
	}
	
	p.LastClickTime = now
	
	actualLine := p.ScrollOffset + lineY - 2 // Adjust for border and title
	if actualLine >= 0 && actualLine < len(p.Lines) {
		p.IsSelecting = true
		p.SelectionStart = actualLine
		p.SelectionEnd = actualLine
		p.SelectedText = ""
	}
	return false
}

// SelectAll selects all content in the output panel
func (p *OutputPanel) SelectAll() {
	if len(p.Lines) == 0 {
		return
	}
	p.SelectionStart = 0
	p.SelectionEnd = len(p.Lines) - 1
	p.SelectedText = strings.Join(p.Lines, "\n")
	p.IsSelecting = false
	
	// Auto-copy on select all
	p.CopySelection()
}

// UpdateSelection updates the selection end point
func (p *OutputPanel) UpdateSelection(lineY int) {
	if !p.IsSelecting {
		return
	}
	actualLine := p.ScrollOffset + lineY - 2 // Adjust for border and title
	if actualLine >= 0 && actualLine < len(p.Lines) {
		p.SelectionEnd = actualLine
	}
}

// EndSelection ends the selection and builds the selected text
func (p *OutputPanel) EndSelection() string {
	if !p.IsSelecting {
		return ""
	}
	p.IsSelecting = false
	
	// Ensure start <= end
	start, end := p.SelectionStart, p.SelectionEnd
	if start > end {
		start, end = end, start
	}
	
	// Build selected text
	var selectedLines []string
	for i := start; i <= end && i < len(p.Lines); i++ {
		selectedLines = append(selectedLines, p.Lines[i])
	}
	p.SelectedText = strings.Join(selectedLines, "\n")
	return p.SelectedText
}

// CopySelection copies the selected text to clipboard
func (p *OutputPanel) CopySelection() error {
	if p.SelectedText == "" {
		return nil
	}
	err := clipboard.WriteAll(p.SelectedText)
	if err != nil {
		p.CopyMessage = "❌ Copy failed!"
		return err
	}
	p.CopyMessage = "✅ Selection copied!"
	return nil
}

// ClearSelection clears the current selection
func (p *OutputPanel) ClearSelection() {
	p.IsSelecting = false
	p.SelectionStart = 0
	p.SelectionEnd = 0
	p.SelectedText = ""
}

// HasSelection returns true if there's an active selection
func (p *OutputPanel) HasSelection() bool {
	return p.SelectedText != ""
}

// IsLineSelected returns true if the given line index is part of the selection
func (p *OutputPanel) IsLineSelected(lineIndex int) bool {
	if p.SelectedText == "" && !p.IsSelecting {
		return false
	}
	start, end := p.SelectionStart, p.SelectionEnd
	if start > end {
		start, end = end, start
	}
	return lineIndex >= start && lineIndex <= end
}

// Clear clears the output
func (p *OutputPanel) Clear() {
	p.Content = ""
	p.Lines = make([]string, 0)
	p.Entries = make([]OutputEntry, 0)
	p.SelectedEntry = -1
	p.ScrollOffset = 0
	p.CopyMessage = ""
}

// ScrollUp scrolls the output up
func (p *OutputPanel) ScrollUp() {
	if p.ScrollOffset > 0 {
		p.ScrollOffset--
	}
}

// ScrollDown scrolls the output down
func (p *OutputPanel) ScrollDown() {
	maxOffset := len(p.Lines) - p.visibleLines()
	if maxOffset < 0 {
		maxOffset = 0
	}
	if p.ScrollOffset < maxOffset {
		p.ScrollOffset++
	}
}

// ScrollToBottom scrolls to the bottom
func (p *OutputPanel) ScrollToBottom() {
	maxOffset := len(p.Lines) - p.visibleLines()
	if maxOffset < 0 {
		maxOffset = 0
	}
	p.ScrollOffset = maxOffset
}

// ScrollToTop scrolls to the top
func (p *OutputPanel) ScrollToTop() {
	p.ScrollOffset = 0
}

// visibleLines returns the number of visible lines
func (p *OutputPanel) visibleLines() int {
	return p.Height - 2 // Account for borders
}

// SetWidth sets the panel width
func (p *OutputPanel) SetWidth(width int) {
	p.Width = width
}

// SetHeight sets the panel height
func (p *OutputPanel) SetHeight(height int) {
	p.Height = height
}

// View renders the output panel
func (p *OutputPanel) View() string {
	// Title with entry count and copy hints
	titleParts := []string{"📺 Output"}
	if len(p.Entries) > 0 {
		titleParts = append(titleParts, lipgloss.NewStyle().Foreground(ColorMuted).Render(
			" │ Ctrl+Y: copy output │ Ctrl+B: copy cmd"))
	}
	if p.CopyMessage != "" {
		titleParts = append(titleParts, " "+p.CopyMessage)
	}
	titleText := p.styles.OutputPanelTitle.Render(strings.Join(titleParts, ""))

	visibleCount := p.visibleLines() - 1 // Account for title
	if visibleCount < 1 {
		visibleCount = 1
	}
	var lines []string

	if len(p.Lines) == 0 {
		lines = append(lines, p.styles.SuggestionDesc.Render("  Press Enter to execute a command..."))
		lines = append(lines, "")
		lines = append(lines, p.styles.SuggestionDesc.Render("  Copy shortcuts:"))
		lines = append(lines, p.styles.SuggestionDesc.Render("  • Ctrl+Y - Copy last output"))
		lines = append(lines, p.styles.SuggestionDesc.Render("  • Ctrl+B - Copy last command"))
	} else {
		endIndex := p.ScrollOffset + visibleCount
		if endIndex > len(p.Lines) {
			endIndex = len(p.Lines)
		}

		for i := p.ScrollOffset; i < endIndex; i++ {
			line := p.Lines[i]
			// Truncate long lines
			if len(line) > p.Width-6 {
				line = line[:p.Width-9] + "..."
			}
			
			// Apply Unix/Linux style coloring based on line type
			styledLine := p.styleLine(line)
			
			// Highlight selected lines
			if p.IsLineSelected(i) {
				styledLine = lipgloss.NewStyle().
					Background(lipgloss.Color("#3B82F6")).
					Foreground(lipgloss.Color("#FFFFFF")).
					Render(line)
			}
			
			lines = append(lines, styledLine)
		}

		// Show scroll and copy info if content exceeds view
		if len(p.Lines) > visibleCount {
			scrollInfo := lipgloss.NewStyle().Foreground(ColorMuted).Render(
				" [↑/↓ scroll] [Ctrl+Y copy output] [Ctrl+B copy cmd]",
			)
			lines = append(lines, scrollInfo)
		}
	}

	// Pad lines to fill height
	for len(lines) < visibleCount {
		lines = append(lines, "")
	}

	content := strings.Join(lines, "\n")

	// Render panel with title inside border
	panel := p.styles.OutputPanel.
		Width(p.Width - 2).
		Height(p.Height).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true).
		Render(titleText + "\n" + content)

	return panel
}

// styleLine applies Unix/Linux style coloring to a line based on its content
func (p *OutputPanel) styleLine(line string) string {
	// Command prompt line (starts with $)
	if strings.HasPrefix(line, "$ ") {
		prompt := p.styles.OutputPrompt.Render("$ ")
		command := p.styles.OutputCommand.Render(line[2:])
		return prompt + command
	}

	// Separator line (─────)
	if strings.HasPrefix(line, "═") || strings.HasPrefix(line, "─") || strings.HasPrefix(line, "━") {
		return p.styles.OutputSeparator.Render(line)
	}

	// Exit code success
	if strings.Contains(line, "[Exit code: 0]") {
		return p.styles.OutputExitOK.Render("✓ ") + p.styles.OutputDuration.Render(line)
	}

	// Exit code failure
	if strings.Contains(line, "[Exit code:") && !strings.Contains(line, "[Exit code: 0]") {
		return p.styles.OutputExitFail.Render("✗ ") + p.styles.OutputError.Render(line)
	}

	// Duration line
	if strings.HasPrefix(line, "[Duration:") {
		return p.styles.OutputDuration.Render(line)
	}

	// Error lines
	if strings.HasPrefix(strings.ToLower(line), "error") ||
		strings.HasPrefix(strings.ToLower(line), "fatal") ||
		strings.HasPrefix(strings.ToLower(line), "failed") {
		return p.styles.OutputError.Render(line)
	}

	// Warning lines
	if strings.HasPrefix(strings.ToLower(line), "warning") ||
		strings.HasPrefix(strings.ToLower(line), "warn") {
		return lipgloss.NewStyle().Foreground(ColorWarning).Render(line)
	}

	// Default output text
	return p.styles.OutputText.Render(line)
}
