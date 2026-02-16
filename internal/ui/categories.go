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

// View renders the categories panel
func (p *CategoriesPanel) View() string {
	titleText := p.styles.SuggestionsPanelTitle.Render("🛠 Supported Technologies")

	// Build category badges
	var badges []string
	for _, cat := range p.Categories {
		badge := p.renderBadge(cat)
		badges = append(badges, badge)
	}

	// Join badges with spacing
	content := strings.Join(badges, "  ")

	// Render panel with title inside border
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

// renderBadge renders a single category badge with icon
func (p *CategoriesPanel) renderBadge(category string) string {
	icon := getCategoryIcon(category)
	return p.styles.SuggestionCategory.Render(icon + " " + category)
}

// getCategoryIcon returns an emoji icon for a category
func getCategoryIcon(category string) string {
	icons := map[string]string{
		"docker":     "🐳",
		"kubernetes": "☸️",
		"gcloud":     "☁️",
		"azure":      "⚡",
		"curl":       "🌐",
		"git":        "📦",
		"terraform":  "🏗️",
		"ansible":    "🔧",
		"helm":       "⎈",
		"aws":        "🔶",
		"ssh":        "🔐",
		"tcpdump":    "🔬",
		"netstat":    "📡",
		"linux":      "🐧",
		"nginx":      "🌿",
		"conda":      "🐍",
		"tmux":       "🖥️",
		"grep":       "🔍",
		"find":       "📂",
		"custom":     "⚙️",
	}

	if icon, ok := icons[strings.ToLower(category)]; ok {
		return icon
	}
	return "📌"
}
