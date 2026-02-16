package ui

import (
	"fmt"
	"strings"

	"github.com/architerm/architerm/internal/autocomplete"
)

// SuggestionsPanel represents the suggestions area
type SuggestionsPanel struct {
	Items         []autocomplete.Match
	SelectedIndex int
	MaxVisible    int
	ScrollOffset  int
	Width         int
	Height        int
	styles        *Styles
}

// NewSuggestionsPanel creates a new suggestions panel
func NewSuggestionsPanel(styles *Styles) *SuggestionsPanel {
	return &SuggestionsPanel{
		Items:         make([]autocomplete.Match, 0),
		SelectedIndex: 0,
		MaxVisible:    5,
		ScrollOffset:  0,
		Width:         80,
		Height:        7,
		styles:        styles,
	}
}

// SetItems sets the suggestion items
func (p *SuggestionsPanel) SetItems(items []autocomplete.Match) {
	p.Items = items
	p.SelectedIndex = 0
	p.ScrollOffset = 0
}

// MoveUp moves selection up
func (p *SuggestionsPanel) MoveUp() {
	if p.SelectedIndex > 0 {
		p.SelectedIndex--
		// Adjust scroll if needed
		if p.SelectedIndex < p.ScrollOffset {
			p.ScrollOffset = p.SelectedIndex
		}
	}
}

// MoveDown moves selection down
func (p *SuggestionsPanel) MoveDown() {
	if p.SelectedIndex < len(p.Items)-1 {
		p.SelectedIndex++
		// Adjust scroll if needed
		if p.SelectedIndex >= p.ScrollOffset+p.MaxVisible {
			p.ScrollOffset = p.SelectedIndex - p.MaxVisible + 1
		}
	}
}

// GetSelected returns the currently selected item
func (p *SuggestionsPanel) GetSelected() *autocomplete.Match {
	if len(p.Items) == 0 || p.SelectedIndex >= len(p.Items) {
		return nil
	}
	return &p.Items[p.SelectedIndex]
}

// SetWidth sets the panel width
func (p *SuggestionsPanel) SetWidth(width int) {
	p.Width = width
}

// SetHeight sets the panel height
func (p *SuggestionsPanel) SetHeight(height int) {
	p.Height = height
	p.MaxVisible = height - 2 // Account for borders
	if p.MaxVisible < 1 {
		p.MaxVisible = 1
	}
}

// View renders the suggestions panel
func (p *SuggestionsPanel) View() string {
	// Title with count
	countStr := ""
	if len(p.Items) > 0 {
		countStr = fmt.Sprintf(" (%d)", len(p.Items))
	}
	titleText := p.styles.SuggestionsPanelTitle.Render("📋 Suggestions" + countStr)

	// Build suggestion lines
	var lines []string
	
	if len(p.Items) == 0 {
		lines = append(lines, p.styles.SuggestionDesc.Render("  Type to search commands..."))
	} else {
		// Calculate visible range
		endIndex := p.ScrollOffset + p.MaxVisible
		if endIndex > len(p.Items) {
			endIndex = len(p.Items)
		}

		for i := p.ScrollOffset; i < endIndex; i++ {
			item := p.Items[i]
			
			// Format command and description
			cmd := truncateString(item.Command, p.Width-30)
			desc := truncateString(item.Description, 20)
			
			if i == p.SelectedIndex {
				line := p.styles.SuggestionSelected.Render(fmt.Sprintf("▶ %s", cmd))
				lines = append(lines, line)
			} else {
				line := fmt.Sprintf("  %s  %s", cmd, p.styles.SuggestionDesc.Render(desc))
				lines = append(lines, p.styles.SuggestionItem.Render(line))
			}
		}

		// Show scroll indicator if needed
		if len(p.Items) > p.MaxVisible {
			scrollInfo := fmt.Sprintf("  [%d-%d of %d]", p.ScrollOffset+1, endIndex, len(p.Items))
			lines = append(lines, p.styles.SuggestionDesc.Render(scrollInfo))
		}
	}

	// Pad lines to fill height - account for title
	maxLines := p.MaxVisible - 1
	if maxLines < 1 {
		maxLines = 1
	}
	for len(lines) < maxLines {
		lines = append(lines, "")
	}

	content := strings.Join(lines, "\n")

	// Render panel with title inside
	panel := p.styles.SuggestionsPanel.
		Width(p.Width - 2).
		Height(p.Height).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true).
		Render(titleText + "\n" + content)

	return panel
}

// truncateString truncates a string to maxLen, adding "..." if needed
func truncateString(s string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}
