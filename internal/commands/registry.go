package commands

import (
	"log"
	"sort"
	"strings"
)

// Command represents a single command template
type Command struct {
	Template    string   `yaml:"template" json:"template"`
	Description string   `yaml:"description" json:"description"`
	Category    string   `yaml:"category" json:"category"`
	Tags        []string `yaml:"tags,omitempty" json:"tags,omitempty"`
}

// Registry holds all registered commands
type Registry struct {
	commands []Command
}

// NewRegistry creates a new command registry with default commands
func NewRegistry() *Registry {
	r := &Registry{
		commands: make([]Command, 0),
	}
	r.loadDefaults()
	return r
}

// loadDefaults loads all built-in commands from embedded JSON files
func (r *Registry) loadDefaults() {
	embeddedCmds, err := LoadEmbeddedCommands()
	if err != nil {
		log.Printf("Warning: failed to load embedded commands: %v", err)
		return
	}
	r.commands = append(r.commands, embeddedCmds...)
}

// AddCommands adds custom commands to the registry
func (r *Registry) AddCommands(cmds []Command) {
	r.commands = append(r.commands, cmds...)
}

// GetAll returns all commands
func (r *Registry) GetAll() []Command {
	return r.commands
}

// GetTemplates returns all command templates
func (r *Registry) GetTemplates() []string {
	templates := make([]string, len(r.commands))
	for i, cmd := range r.commands {
		templates[i] = cmd.Template
	}
	return templates
}

// Search finds commands matching the query
func (r *Registry) Search(query string) []Command {
	if query == "" {
		return r.commands
	}

	query = strings.ToLower(query)
	var matches []Command

	for _, cmd := range r.commands {
		template := strings.ToLower(cmd.Template)
		desc := strings.ToLower(cmd.Description)
		category := strings.ToLower(cmd.Category)

		// Check if query matches template, description, or category
		if strings.Contains(template, query) ||
			strings.Contains(desc, query) ||
			strings.Contains(category, query) {
			matches = append(matches, cmd)
		}

		// Check tags
		for _, tag := range cmd.Tags {
			if strings.Contains(strings.ToLower(tag), query) {
				matches = append(matches, cmd)
				break
			}
		}
	}

	// Sort by relevance (prefix matches first)
	sort.Slice(matches, func(i, j int) bool {
		iTemplate := strings.ToLower(matches[i].Template)
		jTemplate := strings.ToLower(matches[j].Template)
		iPrefix := strings.HasPrefix(iTemplate, query)
		jPrefix := strings.HasPrefix(jTemplate, query)
		if iPrefix && !jPrefix {
			return true
		}
		if !iPrefix && jPrefix {
			return false
		}
		return iTemplate < jTemplate
	})

	return matches
}

// GetByCategory returns commands in a specific category
func (r *Registry) GetByCategory(category string) []Command {
	var matches []Command
	category = strings.ToLower(category)
	for _, cmd := range r.commands {
		if strings.ToLower(cmd.Category) == category {
			matches = append(matches, cmd)
		}
	}
	return matches
}

// GetCategories returns a list of unique categories
func (r *Registry) GetCategories() []string {
	seen := make(map[string]bool)
	var categories []string
	
	for _, cmd := range r.commands {
		cat := strings.ToLower(cmd.Category)
		if cat != "" && !seen[cat] {
			seen[cat] = true
			categories = append(categories, cat)
		}
	}
	
	// Sort alphabetically
	sort.Strings(categories)
	return categories
}
