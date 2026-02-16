package app

import (
	"fmt"

	"github.com/duladissa/architerm/internal/autocomplete"
	"github.com/duladissa/architerm/internal/commands"
	"github.com/duladissa/architerm/internal/executor"
	"github.com/duladissa/architerm/internal/history"
	"github.com/duladissa/architerm/internal/theme"
	"github.com/duladissa/architerm/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

// Model represents the application state
type Model struct {
	// UI components
	styles      *ui.Styles
	layout      *ui.Layout
	inputPanel  *ui.InputPanel
	suggestions *ui.SuggestionsPanel
	categories  *ui.CategoriesPanel
	outputPanel *ui.OutputPanel

	// Core components
	registry   *commands.Registry
	engine     *autocomplete.Engine
	executor   *executor.Executor
	history    *history.History

	// State
	width      int
	height     int
	status     string
	isRunning  bool
	configPath string
}

// CommandResultMsg is sent when a command finishes executing
type CommandResultMsg struct {
	Result *executor.Result
}

// NewModel creates a new application model
func NewModel(configPath string) *Model {
	styles := ui.DefaultStyles()
	
	m := &Model{
		styles:      styles,
		layout:      ui.NewLayout(styles),
		inputPanel:  ui.NewInputPanel(styles),
		suggestions: ui.NewSuggestionsPanel(styles),
		categories:  ui.NewCategoriesPanel(styles),
		outputPanel: ui.NewOutputPanel(styles),
		registry:    commands.NewRegistry(),
		engine:      autocomplete.NewEngine(),
		executor:    executor.NewExecutor(),
		history:     history.NewHistory(100),
		width:       80,
		height:      24,
		status:      "",
		isRunning:   false,
		configPath:  configPath,
	}

	// Load custom config if provided
	if configPath != "" {
		customCmds, err := commands.LoadConfig(configPath)
		if err != nil {
			m.status = fmt.Sprintf("Config error: %v", err)
		} else {
			m.registry.AddCommands(customCmds)
		}
	}

	// Try loading user config from default location
	userCmds, _ := commands.LoadUserConfig()
	if len(userCmds) > 0 {
		m.registry.AddCommands(userCmds)
	}

	// Populate autocomplete engine
	for _, cmd := range m.registry.GetAll() {
		m.engine.AddCommand(cmd.Template, cmd.Description)
	}

	// Set supported categories from registry
	m.categories.SetCategories(m.registry.GetCategories())

	// Initialize suggestions
	m.updateSuggestions()

	return m
}

// Init implements tea.Model
func (m *Model) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyPress(msg)

	case tea.MouseMsg:
		return m.handleMouseEvent(msg)

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.updateLayout()
		return m, nil

	case CommandResultMsg:
		m.isRunning = false
		m.status = ""
		fullText := executor.FormatResult(msg.Result)
		// Add as entry for easy copying
		m.outputPanel.AddEntry(msg.Result.Command, msg.Result.Output, fullText)
		return m, nil
	}

	return m, nil
}

// handleMouseEvent handles mouse input (scroll wheel and selection)
func (m *Model) handleMouseEvent(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	// Calculate if mouse is in output panel area
	// Output panel is on the right side (approximately after 45% of width)
	outputPanelStartX := m.layout.GetLeftPanelWidth() + 2
	isInOutputPanel := msg.X >= outputPanelStartX
	
	switch msg.Type {
	case tea.MouseWheelUp:
		// Scroll output up
		m.outputPanel.ScrollUp()
		return m, nil
	case tea.MouseWheelDown:
		// Scroll output down
		m.outputPanel.ScrollDown()
		return m, nil
	case tea.MouseLeft:
		// Start selection in output panel (or select all on double-click)
		if isInOutputPanel {
			isDoubleClick := m.outputPanel.StartSelection(msg.Y)
			if isDoubleClick {
				m.status = "All output selected & copied!"
			}
		}
		return m, nil
	case tea.MouseMotion:
		// Update selection while dragging
		if isInOutputPanel && m.outputPanel.IsSelecting {
			m.outputPanel.UpdateSelection(msg.Y)
		}
		return m, nil
	case tea.MouseRelease:
		// End selection and copy to clipboard
		if isInOutputPanel && m.outputPanel.IsSelecting {
			selectedText := m.outputPanel.EndSelection()
			if selectedText != "" {
				m.outputPanel.CopySelection()
			}
		}
		return m, nil
	case tea.MouseRight, tea.MouseMiddle:
		// Consume these events without doing anything
		return m, nil
	}
	return m, nil
}

// handleKeyPress handles keyboard input
func (m *Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Check for Shift+Arrow keys for output scrolling
	if msg.Alt {
		switch msg.Type {
		case tea.KeyUp:
			m.outputPanel.ScrollUp()
			return m, nil
		case tea.KeyDown:
			m.outputPanel.ScrollDown()
			return m, nil
		}
	}

	switch msg.Type {
	case tea.KeyCtrlC:
		if m.isRunning {
			m.executor.Cancel()
			m.status = "Command cancelled"
			return m, nil
		}
		return m, tea.Quit

	case tea.KeyEsc:
		m.inputPanel.Clear()
		m.updateSuggestions()
		return m, nil

	case tea.KeyTab:
		// Accept ghost text or selected suggestion
		if m.inputPanel.GhostText != "" {
			m.inputPanel.AcceptGhostText()
		} else if selected := m.suggestions.GetSelected(); selected != nil {
			m.inputPanel.SetValue(selected.Command)
		}
		m.updateSuggestions()
		return m, nil

	case tea.KeyEnter:
		if m.inputPanel.Value != "" && !m.isRunning {
			return m.executeCommand()
		}
		return m, nil

	case tea.KeyUp:
		if len(m.suggestions.Items) > 0 {
			m.suggestions.MoveUp()
		} else {
			// Navigate history
			if prev := m.history.Previous(); prev != "" {
				m.inputPanel.SetValue(prev)
				m.updateSuggestions()
			}
		}
		return m, nil

	case tea.KeyDown:
		if len(m.suggestions.Items) > 0 {
			m.suggestions.MoveDown()
		} else {
			// Navigate history
			if next := m.history.Next(); next != "" {
				m.inputPanel.SetValue(next)
				m.updateSuggestions()
			}
		}
		return m, nil

	case tea.KeyPgUp:
		// Page up scrolls output
		for i := 0; i < 5; i++ {
			m.outputPanel.ScrollUp()
		}
		return m, nil

	case tea.KeyPgDown:
		// Page down scrolls output
		for i := 0; i < 5; i++ {
			m.outputPanel.ScrollDown()
		}
		return m, nil

	case tea.KeyLeft:
		m.inputPanel.MoveCursorLeft()
		return m, nil

	case tea.KeyRight:
		m.inputPanel.MoveCursorRight()
		return m, nil

	case tea.KeyHome:
		m.inputPanel.MoveCursorStart()
		return m, nil

	case tea.KeyEnd:
		m.inputPanel.MoveCursorEnd()
		return m, nil

	case tea.KeyBackspace:
		m.inputPanel.DeleteChar()
		m.updateSuggestions()
		return m, nil

	case tea.KeyDelete:
		m.inputPanel.DeleteCharForward()
		m.updateSuggestions()
		return m, nil

	case tea.KeyCtrlL:
		m.outputPanel.Clear()
		return m, nil

	case tea.KeyCtrlU:
		m.inputPanel.Clear()
		m.updateSuggestions()
		return m, nil

	case tea.KeyCtrlY:
		// Copy last output to clipboard
		m.outputPanel.CopyLastOutput()
		return m, nil

	case tea.KeyCtrlB:
		// Copy last command to clipboard
		m.outputPanel.CopyLastCommand()
		return m, nil

	case tea.KeyCtrlT:
		// Cycle through themes
		m.cycleTheme()
		return m, nil

	case tea.KeyRunes:
		// Filter out mouse escape sequence characters that might leak through
		// Mouse sequences typically have multiple characters with digits and special chars
		// e.g., "<64;123;45M" or similar patterns
		
		// Only filter if it looks like a mouse escape sequence (has digits + special chars)
		runeStr := string(msg.Runes)
		if len(msg.Runes) > 3 {
			// Check if it contains digits mixed with semicolons (likely mouse coords)
			hasDigit := false
			hasSemicolon := false
			for _, r := range msg.Runes {
				if r >= '0' && r <= '9' {
					hasDigit = true
				}
				if r == ';' {
					hasSemicolon = true
				}
			}
			if hasDigit && hasSemicolon {
				// Likely a mouse escape sequence, ignore
				return m, nil
			}
		}
		
		// Check for specific mouse sequence patterns like "<0;123;45M"
		if len(runeStr) > 5 && (runeStr[0] == '<' || runeStr[0] == '[') {
			return m, nil
		}
		
		for _, r := range msg.Runes {
			// Only filter out control characters
			if r < 32 || r == 127 {
				continue
			}
			m.inputPanel.InsertChar(r)
		}
		m.updateSuggestions()
		return m, nil

	case tea.KeySpace:
		m.inputPanel.InsertChar(' ')
		m.updateSuggestions()
		return m, nil
	}

	return m, nil
}

// executeCommand runs the current input command
func (m *Model) executeCommand() (tea.Model, tea.Cmd) {
	command := m.inputPanel.Value
	m.history.Add(command)
	m.history.Reset()
	m.inputPanel.Clear()
	m.updateSuggestions()
	m.isRunning = true
	m.status = "Running..."

	// Execute command asynchronously
	return m, func() tea.Msg {
		result := m.executor.Execute(command)
		return CommandResultMsg{Result: result}
	}
}

// updateSuggestions updates the suggestions based on current input
func (m *Model) updateSuggestions() {
	input := m.inputPanel.Value
	
	// Get suggestions from engine
	suggestions := m.engine.GetSuggestions(input, 20)
	m.suggestions.SetItems(suggestions)

	// Update ghost text
	ghostText := m.engine.GetGhostText(input)
	m.inputPanel.SetGhostText(ghostText)
}

// updateLayout updates panel sizes based on terminal size
func (m *Model) updateLayout() {
	m.layout.SetSize(m.width, m.height)
	
	// Left panel (input + suggestions + categories) width
	leftWidth := m.layout.GetLeftPanelWidth()
	m.inputPanel.SetWidth(leftWidth)
	m.suggestions.SetWidth(leftWidth)
	m.suggestions.SetHeight(m.layout.GetSuggestionsHeight())
	m.categories.SetWidth(leftWidth)
	m.categories.SetHeight(m.layout.GetCategoriesHeight())
	
	// Right panel (output) width
	rightWidth := m.layout.GetRightPanelWidth()
	m.outputPanel.SetWidth(rightWidth)
	m.outputPanel.SetHeight(m.layout.GetOutputHeight())
}

// refreshStyles recreates all styles with the current theme
func (m *Model) refreshStyles() {
	m.styles = ui.DefaultStyles()
	m.layout.SetStyles(m.styles)
	m.inputPanel.SetStyles(m.styles)
	m.suggestions.SetStyles(m.styles)
	m.categories.SetStyles(m.styles)
	m.outputPanel.SetStyles(m.styles)
}

// cycleTheme switches to the next available theme
func (m *Model) cycleTheme() {
	newTheme := theme.CycleTheme()
	m.refreshStyles()
	m.status = fmt.Sprintf("Theme: %s", newTheme)
}

// View implements tea.Model
func (m *Model) View() string {
	header := m.layout.RenderHeader()
	input := m.inputPanel.View()
	suggestions := m.suggestions.View()
	categories := m.categories.View()
	output := m.outputPanel.View()
	statusBar := m.layout.RenderStatusBar(m.status)

	return m.layout.Render(header, input, suggestions, categories, output, statusBar)
}

// Run starts the application
func Run(configPath string) error {
	model := NewModel(configPath)
	// Use WithMouseAllMotion for better compatibility, avoiding raw escape sequences
	p := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseAllMotion())
	_, err := p.Run()
	return err
}
