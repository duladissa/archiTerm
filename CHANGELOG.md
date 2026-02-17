# Changelog

All notable changes to archiTerm will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-02-16

### Added

#### Core Features
- **Smart Autocomplete Engine**: Gmail-style ghost text suggestions with fuzzy matching
- **Trie-based Search**: Fast prefix matching for command suggestions
- **Command Execution**: Real-time command execution with live output streaming
- **Session History**: Navigate through previously executed commands with ↑/↓ keys

#### User Interface
- **Three-panel Layout**: Input, suggestions, and output panels in a modern TUI
- **Category Browser**: Visual display of all supported technologies
- **Mouse Support**: Scroll output with mouse wheel
- **Responsive Design**: Adapts to terminal size changes

#### Theme System
- **4 Built-in Themes**: dark, dracula, nord, gruvbox
- **Runtime Theme Switching**: Press `Ctrl+T` to cycle through themes instantly
- **Custom Theme Support**: Load themes from JSON files
- **Full Background Colors**: Themes work correctly on any terminal background

#### Pre-loaded Commands
- **Docker**: Container and image management commands
- **Kubernetes**: kubectl commands for cluster management
- **Git**: Version control commands
- **gcloud**: Google Cloud Platform CLI commands
- **Azure**: Microsoft Azure CLI commands
- **curl**: HTTP request commands
- **SSH**: Secure shell commands
- **Linux**: Common system administration commands
- **nginx**: Web server configuration commands
- **tmux**: Terminal multiplexer commands
- **conda**: Python environment management
- **grep/find**: File search commands
- **netstat/tcpdump**: Network diagnostic commands

#### Clipboard Support
- **Copy Output**: `Ctrl+Y` to copy command output
- **Copy Command**: `Ctrl+B` to copy last executed command

#### Configuration
- **YAML/JSON Config**: Add custom commands via configuration files
- **Embedded Commands**: Pre-loaded commands compiled into the binary

#### Build & Distribution
- **Cross-platform**: Windows, Linux, and macOS support
- **Multi-architecture**: amd64 and arm64 binaries
- **Single Binary**: No runtime dependencies required
- **GitHub Actions CI/CD**: Automated testing and releases

### Technical Details
- Built with Go 1.21+
- Uses [Bubble Tea](https://github.com/charmbracelet/bubbletea) for the TUI framework
- Uses [Lip Gloss](https://github.com/charmbracelet/lipgloss) for styling
- Uses [Cobra](https://github.com/spf13/cobra) for CLI commands

---

## Future Plans

- [ ] Theme persistence (save preferred theme)
- [ ] Custom theme directory (`~/.config/architerm/themes/`)
- [ ] Plugin system for extending commands
- [ ] Shell integration (bash/zsh completions)
- [ ] Command favorites and bookmarks
- [ ] Syntax highlighting for output

[1.0.0]: https://github.com/duladissa/architerm/releases/tag/v1.0.0
