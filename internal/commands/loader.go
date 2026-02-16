package commands

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed embedded/*.json
var embeddedCommands embed.FS

// Config represents the configuration file structure
type Config struct {
	Commands []Command `yaml:"commands" json:"commands"`
}

// EmbeddedConfig represents the structure of embedded JSON files
type EmbeddedConfig struct {
	Category string `json:"category"`
	Commands []struct {
		Template    string   `json:"template"`
		Description string   `json:"description"`
		Tags        []string `json:"tags"`
	} `json:"commands"`
}

// LoadEmbeddedCommands loads all commands from embedded JSON files
func LoadEmbeddedCommands() ([]Command, error) {
	var allCommands []Command

	// Read all files from the embedded directory
	entries, err := embeddedCommands.ReadDir("embedded")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded commands directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		data, err := embeddedCommands.ReadFile("embedded/" + entry.Name())
		if err != nil {
			return nil, fmt.Errorf("failed to read embedded file %s: %w", entry.Name(), err)
		}

		var config EmbeddedConfig
		if err := json.Unmarshal(data, &config); err != nil {
			return nil, fmt.Errorf("failed to parse embedded file %s: %w", entry.Name(), err)
		}

		// Convert to Command structs
		for _, cmd := range config.Commands {
			allCommands = append(allCommands, Command{
				Template:    cmd.Template,
				Description: cmd.Description,
				Category:    config.Category,
				Tags:        cmd.Tags,
			})
		}
	}

	return allCommands, nil
}

// LoadConfig loads commands from a YAML or JSON configuration file
func LoadConfig(path string) ([]Command, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(data, &config); err != nil {
			return nil, fmt.Errorf("failed to parse YAML config: %w", err)
		}
	case ".json":
		if err := json.Unmarshal(data, &config); err != nil {
			return nil, fmt.Errorf("failed to parse JSON config: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported config format: %s (use .yaml, .yml, or .json)", ext)
	}

	return config.Commands, nil
}

// GetDefaultConfigPath returns the default config file path
func GetDefaultConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".config", "architerm", "commands.yaml")
}

// LoadUserConfig attempts to load user configuration from default location
func LoadUserConfig() ([]Command, error) {
	configPath := GetDefaultConfigPath()
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Try JSON as fallback
		jsonPath := strings.TrimSuffix(configPath, ".yaml") + ".json"
		if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
			return nil, nil // No user config, that's fine
		}
		configPath = jsonPath
	}
	return LoadConfig(configPath)
}
