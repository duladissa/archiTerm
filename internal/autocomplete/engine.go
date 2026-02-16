package autocomplete

import (
	"sort"
	"strings"
)

// Match represents a matching command with its score
type Match struct {
	Command     string
	Description string
	Score       int
}

// Engine provides autocomplete functionality
type Engine struct {
	trie     *Trie
	commands []Match
}

// NewEngine creates a new autocomplete engine
func NewEngine() *Engine {
	return &Engine{
		trie:     NewTrie(),
		commands: make([]Match, 0),
	}
}

// AddCommand adds a command to the engine
func (e *Engine) AddCommand(command, description string) {
	e.trie.Insert(command, description)
	e.commands = append(e.commands, Match{
		Command:     command,
		Description: description,
	})
}

// AddCommands adds multiple commands
func (e *Engine) AddCommands(commands []struct{ Command, Description string }) {
	for _, cmd := range commands {
		e.AddCommand(cmd.Command, cmd.Description)
	}
}

// GetSuggestions returns matching commands for the input
func (e *Engine) GetSuggestions(input string, limit int) []Match {
	if input == "" {
		// Return first N commands
		if limit > len(e.commands) {
			limit = len(e.commands)
		}
		return e.commands[:limit]
	}

	// First try prefix matching via trie
	prefixMatches := e.trie.Search(input)

	// Then do fuzzy matching for additional results
	fuzzyMatches := e.fuzzySearch(input)

	// Combine and dedupe results
	seen := make(map[string]bool)
	var results []Match

	// Add prefix matches first (they have higher priority)
	for _, m := range prefixMatches {
		if !seen[m.Command] {
			seen[m.Command] = true
			results = append(results, m)
		}
	}

	// Add fuzzy matches
	for _, m := range fuzzyMatches {
		if !seen[m.Command] {
			seen[m.Command] = true
			results = append(results, m)
		}
	}

	// Sort by score (descending)
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	if limit > 0 && len(results) > limit {
		results = results[:limit]
	}

	return results
}

// GetGhostText returns the completion text to show as ghost text
func (e *Engine) GetGhostText(input string) string {
	if input == "" {
		return ""
	}

	// Get the best prefix match
	completion := e.trie.GetCompletion(input)
	return completion
}

// fuzzySearch performs fuzzy matching on commands
func (e *Engine) fuzzySearch(query string) []Match {
	query = strings.ToLower(query)
	var matches []Match

	for _, cmd := range e.commands {
		command := strings.ToLower(cmd.Command)
		desc := strings.ToLower(cmd.Description)

		score := e.calculateFuzzyScore(query, command, desc)
		if score > 0 {
			matches = append(matches, Match{
				Command:     cmd.Command,
				Description: cmd.Description,
				Score:       score,
			})
		}
	}

	return matches
}

// calculateFuzzyScore calculates a fuzzy match score
func (e *Engine) calculateFuzzyScore(query, command, description string) int {
	score := 0

	// Check if query is contained in command
	if strings.Contains(command, query) {
		score += 50
		// Bonus for prefix match
		if strings.HasPrefix(command, query) {
			score += 30
		}
	}

	// Check word-by-word matching
	queryWords := strings.Fields(query)
	commandWords := strings.Fields(command)

	matchedWords := 0
	for _, qw := range queryWords {
		for _, cw := range commandWords {
			if strings.HasPrefix(cw, qw) {
				matchedWords++
				score += 10
				break
			}
		}
	}

	// Check description for matches
	if strings.Contains(description, query) {
		score += 20
	}

	// Check if all characters appear in order (subsequence)
	if e.isSubsequence(query, command) && score == 0 {
		score += 15
	}

	return score
}

// isSubsequence checks if query chars appear in order in target
func (e *Engine) isSubsequence(query, target string) bool {
	qi := 0
	for ti := 0; ti < len(target) && qi < len(query); ti++ {
		if target[ti] == query[qi] {
			qi++
		}
	}
	return qi == len(query)
}

// GetBestMatch returns the best matching command for the input
func (e *Engine) GetBestMatch(input string) *Match {
	suggestions := e.GetSuggestions(input, 1)
	if len(suggestions) > 0 {
		return &suggestions[0]
	}
	return nil
}
