package history

// History manages session command history
type History struct {
	commands []string
	position int
	maxSize  int
}

// NewHistory creates a new history manager
func NewHistory(maxSize int) *History {
	return &History{
		commands: make([]string, 0),
		position: -1,
		maxSize:  maxSize,
	}
}

// Add adds a command to history
func (h *History) Add(command string) {
	if command == "" {
		return
	}

	// Don't add duplicates consecutively
	if len(h.commands) > 0 && h.commands[len(h.commands)-1] == command {
		h.position = len(h.commands)
		return
	}

	h.commands = append(h.commands, command)

	// Trim if exceeds max size
	if len(h.commands) > h.maxSize {
		h.commands = h.commands[1:]
	}

	// Reset position to end
	h.position = len(h.commands)
}

// Previous returns the previous command in history
func (h *History) Previous() string {
	if len(h.commands) == 0 {
		return ""
	}

	if h.position > 0 {
		h.position--
	}

	if h.position >= 0 && h.position < len(h.commands) {
		return h.commands[h.position]
	}

	return ""
}

// Next returns the next command in history
func (h *History) Next() string {
	if len(h.commands) == 0 {
		return ""
	}

	if h.position < len(h.commands)-1 {
		h.position++
		return h.commands[h.position]
	}

	// At the end, return empty to allow new input
	h.position = len(h.commands)
	return ""
}

// Reset resets the position to the end
func (h *History) Reset() {
	h.position = len(h.commands)
}

// GetAll returns all commands in history
func (h *History) GetAll() []string {
	return h.commands
}

// Len returns the number of commands in history
func (h *History) Len() int {
	return len(h.commands)
}

// Clear clears all history
func (h *History) Clear() {
	h.commands = make([]string, 0)
	h.position = -1
}

// Search finds commands containing the query
func (h *History) Search(query string) []string {
	if query == "" {
		return h.commands
	}

	var matches []string
	for _, cmd := range h.commands {
		if containsIgnoreCase(cmd, query) {
			matches = append(matches, cmd)
		}
	}
	return matches
}

func containsIgnoreCase(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if equalFoldAt(s, substr, i) {
			return true
		}
	}
	return false
}

func equalFoldAt(s, substr string, start int) bool {
	for i := 0; i < len(substr); i++ {
		c1 := s[start+i]
		c2 := substr[i]
		if c1 != c2 && toLower(c1) != toLower(c2) {
			return false
		}
	}
	return true
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}
