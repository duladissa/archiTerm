package ui

import (
	"strings"
)

// CategoriesPanel represents the supported technologies display
type CategoriesPanel struct {
	Categories []string
	Width      int
	Height     int
	styles     *Styles
}

// NewCategoriesPanel creates a new categories panel
func NewCategoriesPanel(styles *Styles) *CategoriesPanel {
	return &CategoriesPanel{
		Categories: make([]string, 0),
		Width:      80,
		Height:     3,
		styles:     styles,
	}
}

// SetCategories sets the list of supported categories
func (p *CategoriesPanel) SetCategories(categories []string) {
	p.Categories = categories
}

// SetWidth sets the panel width
func (p *CategoriesPanel) SetWidth(width int) {
	p.Width = width
}

// SetHeight sets the panel height
func (p *CategoriesPanel) SetHeight(height int) {
	p.Height = height
}

// SetStyles updates the styles for the categories panel
func (p *CategoriesPanel) SetStyles(styles *Styles) {
	p.styles = styles
}

// View renders the categories panel
func (p *CategoriesPanel) View() string {
	titleText := p.styles.SuggestionsPanelTitle.Render("🛠 Supported Technologies")

	// Fixed column width for proper alignment
	colWidth := 14
	numCols := 3
	
	// Adjust columns based on width
	if p.Width > 50 {
		numCols = 3
	}
	if p.Width < 40 {
		numCols = 2
	}

	// Build table rows with proper formatting
	var rows []string
	var currentRow []string

	for i, cat := range p.Categories {
		icon := getCategoryIcon(cat)
		// Format: "icon name" with fixed width (plain text, no individual styling)
		entry := icon + cat
		// Pad to fixed width
		padding := colWidth - len(cat) - 2
		if padding < 1 {
			padding = 1
		}
		paddedEntry := entry + strings.Repeat(" ", padding)
		currentRow = append(currentRow, paddedEntry)

		// Start new row when we reach numCols
		if (i+1)%numCols == 0 || i == len(p.Categories)-1 {
			// Join with separator
			rowText := " " + strings.Join(currentRow, "│ ")
			rows = append(rows, rowText)
			currentRow = []string{}
		}
	}

	// Add separator line between title and content
	separatorWidth := p.Width - 6
	if separatorWidth < 10 {
		separatorWidth = 10
	}
	separator := " " + strings.Repeat("─", separatorWidth)

	// Join rows with newlines
	content := separator + "\n" + strings.Join(rows, "\n")

	// Pad content to fill height
	contentLines := len(rows) + 1 // +1 for separator
	availableLines := p.Height - 3 // Account for title and borders
	if contentLines < availableLines {
		content += strings.Repeat("\n", availableLines-contentLines)
	}

	// Render panel with title inside border - panel style already has background from theme
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

// renderBadge renders a single category badge with icon (plain text, no background)
func (p *CategoriesPanel) renderBadge(category string) string {
	icon := getCategoryIcon(category)
	return icon + " " + category
}


// getCategoryIcon returns an emoji icon for a category
func getCategoryIcon(category string) string {
	icons := map[string]string{
		"docker":     "🐳 ",
		"kubernetes": "☸️  ",
		"gcloud":     "☁️  ",
		"azure":      "⚡ ",
		"curl":       "🌐 ",
		"git":        "📦 ",
		"terraform":  "🏗️  ",
		"ansible":    "🔧 ",
		"helm":       "⎈  ",
		"aws":        "🔶 ",
		"ssh":        "🔐 ",
		"tcpdump":    "🔬 ",
		"netstat":    "📡 ",
		"linux":      "🐧 ",
		"nginx":      "🌿 ",
		"conda":      "🐍 ",
		"tmux":       "🖥️  ",
		"grep":       "🔍 ",
		"find":       "📂 ",
		"custom":     "⚙️  ",
	}

	if icon, ok := icons[strings.ToLower(category)]; ok {
		return icon
	}
	return "📌 "
}
