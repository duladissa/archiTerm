# 🏛️ archiTerm

A lightweight, cross-platform TUI (Terminal User Interface) application for architects and developers to quickly access and execute common commands with smart autocomplete.

[![CI](https://github.com/duladissa/architerm/actions/workflows/ci.yml/badge.svg)](https://github.com/duladissa/architerm/actions/workflows/ci.yml)
[![Release](https://github.com/duladissa/architerm/actions/workflows/release.yml/badge.svg)](https://github.com/duladissa/architerm/actions/workflows/release.yml)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/duladissa/architerm?include_prereleases)](https://github.com/duladissa/architerm/releases)
[![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey)](https://github.com/duladissa/architerm/releases)

> 🤖 **Made with the help of AI** — Built collaboratively with [Claude](https://claude.ai) by Anthropic

## ✨ Features

- **🔮 Smart Autocomplete**: Gmail-style autocomplete with ghost text suggestions
- **📦 Pre-loaded Commands**: Top 10-15 commands for Docker, Kubernetes, gcloud, Azure, curl, and git
- **🎨 Beautiful TUI**: Side-by-side panel interface with modern styling
- **🌈 Multiple Themes**: 4 built-in color themes (dark, dracula, nord, gruvbox) with runtime switching
- **🎯 Unix/Linux Colors**: Familiar terminal color scheme (green prompt, yellow commands, etc.)
- **📋 Easy Copy**: One-key copy of command output or command text to clipboard
- **⚡ Lightweight**: Single binary, no dependencies required at runtime
- **🔧 Customizable**: Add your own commands via YAML or JSON configuration
- **🖥️ Cross-Platform**: Works on Windows, Linux, and macOS
- **📜 Session History**: Navigate through previously executed commands
- **🖱️ Mouse Support**: Scroll output with mouse wheel
- **🛠 Technology Overview**: Visual display of all supported technologies

## 📸 Screenshot

```
┌───────────────────────────────────────────────────────────────────────────────────┐
│  🏛️  archiTerm                     Tab: complete │ ↑↓: navigate │ Ctrl+C: exit  │
├───────────────────────────────────────┬───────────────────────────────────────────┤
│ ╭───────────────────────────────────╮ │ ╭───────────────────────────────────────╮ │
│ │ ⌨ Command                         │ │ │ 📺 Output │ Ctrl+Y: copy │ Ctrl+B: cmd│ │
│ │ > kubectl get pods -n default     │ │ │                                       │ │
│ ╰───────────────────────────────────╯ │ │ ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━  │ │
│ ╭───────────────────────────────────╮ │ │ $ kubectl get pods -n default         │ │
│ │ 📋 Suggestions (15)               │ │ │ ────────────────────────────────────  │ │
│ │ ▶ kubectl get pods -n NAMESPACE   │ │ │ NAME                 READY   STATUS   │ │
│ │   kubectl get pods --all-namespa..│ │ │ nginx-deployment     1/1     Running  │ │
│ │   kubectl get services            │ │ │ redis-cache          1/1     Running  │ │
│ │   kubectl get deployments         │ │ │ ✓ [Exit code: 0] [Duration: 245ms]   │ │
│ ╰───────────────────────────────────╯ │ │                                       │ │
│ ╭───────────────────────────────────╮ │ │ [↑/↓ scroll] [Ctrl+Y copy output]    │ │
│ │ 🛠 Supported Technologies         │ │ │                                       │ │
│ │ ⚡ azure  🌐 curl  🐳 docker       │ │ │                                       │ │
│ │ ☁️ gcloud  📦 git  ☸️ kubernetes   │ │ │                                       │ │
│ ╰───────────────────────────────────╯ │ ╰───────────────────────────────────────╯ │
├───────────────────────────────────────────────────────────────────────────────────┤
│ Tab complete │ Enter run │ Ctrl+Y copy out │ Ctrl+B copy cmd │ Ctrl+L clear      │
└───────────────────────────────────────────────────────────────────────────────────┘
```

### Output Color Scheme (Unix/Linux Style)

| Element | Color | Description |
|---------|-------|-------------|
| `$` prompt | 🟢 Green | Familiar bash-style prompt |
| Command text | 🟡 Yellow | The executed command |
| Separators | 🟣 Purple | Visual distinction between outputs |
| Output text | ⚪ Light Gray | Command output |
| Duration | 🔵 Blue | Timing information |
| Success ✓ | 🟢 Green | Exit code 0 |
| Failure ✗ | 🔴 Red | Non-zero exit code |
| Errors | 🔴 Red | Error messages |
| Warnings | 🟠 Amber | Warning messages |

## 🚀 Installation

### Using Go Install

```bash
go install github.com/duladissa/architerm@latest
```

### From Source

```bash
git clone https://github.com/duladissa/architerm.git
cd architerm
make build
./build/architerm
```

### Pre-built Binaries

Download the latest release for your platform from the [Releases](https://github.com/duladissa/architerm/releases) page.

| Platform | Architecture | Download |
|----------|--------------|----------|
| Linux    | amd64        | `architerm-linux-amd64` |
| Linux    | arm64        | `architerm-linux-arm64` |
| macOS    | amd64        | `architerm-darwin-amd64` |
| macOS    | arm64 (M1/M2)| `architerm-darwin-arm64` |
| Windows  | amd64        | `architerm-windows-amd64.exe` |

## 📖 Usage

### Basic Usage

```bash
# Start archiTerm
architerm

# Start with custom config
architerm --config /path/to/commands.yaml

# Show version
architerm version
```

### Keyboard Shortcuts

| Key | Action |
|-----|--------|
| `Tab` | Accept autocomplete suggestion |
| `Enter` | Execute the command |
| `↑` / `↓` | Navigate suggestions or history |
| `Page Up` / `Page Down` | Scroll output (5 lines) |
| `Alt + ↑` / `Alt + ↓` | Scroll output (1 line) |
| `Esc` | Clear input |
| `Ctrl+L` | Clear output |
| `Ctrl+U` | Clear input line |
| `Ctrl+T` | Cycle through color themes |
| `Ctrl+C` | Cancel running command / Exit |

### Copy Shortcuts

| Key | Action |
|-----|--------|
| `Ctrl+Y` | Copy last command **output** to clipboard |
| `Ctrl+B` | Copy last **command** to clipboard |

After copying, you'll see a confirmation message:
- ✅ **"Output copied!"** - output successfully copied
- ✅ **"Command copied!"** - command successfully copied

### Mouse Selection & Copy

You can select and copy text from the output panel using your mouse:

| Action | Result |
|--------|--------|
| **Click and drag** | Select specific lines |
| **Double-click** | Select and copy **entire output** |
| **Release mouse** | Auto-copy selection to clipboard |

**How it works:**
1. **Click and drag** in the output panel to select lines
2. **Selected lines** are highlighted in blue
3. **Release mouse** to automatically copy the selection to clipboard
4. **Double-click** anywhere in output to select & copy everything
5. A "✅ Selection copied!" message confirms the copy

This works just like selecting text in a normal terminal!

### Smart "Command Not Found" Handling

When you run a command that's not installed, archiTerm will:
1. ⚠️ Detect the missing command
2. 📦 Show installation instructions with a URL
3. 🔗 Display the install link you can copy

**Example output when `gcloud` is not installed:**
```
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
$ gcloud compute instances list
────────────────────────────────────────

⚠️  COMMAND NOT FOUND
────────────────────────────────────────
The command 'gcloud' is not installed or not in PATH.

📦 HOW TO INSTALL:
   Install Google Cloud SDK
   🔗 https://cloud.google.com/sdk/docs/install

🔍 TROUBLESHOOTING:
   • Verify installation: which <command>
   • Check PATH: echo $PATH
   • Reload shell: source ~/.bashrc (or ~/.zshrc)

✗ [Exit code: 127] [Duration: 5ms]
```

**Supported commands with install hints:**
| Command | Install URL |
|---------|-------------|
| `docker` | https://docs.docker.com/get-docker/ |
| `docker-compose` | https://docs.docker.com/compose/install/ |
| `kubectl` | https://kubernetes.io/docs/tasks/tools/ |
| `gcloud` | https://cloud.google.com/sdk/docs/install |
| `az` | https://docs.microsoft.com/cli/azure/install-azure-cli |
| `aws` | https://aws.amazon.com/cli/ |
| `terraform` | https://www.terraform.io/downloads |
| `helm` | https://helm.sh/docs/intro/install/ |
| `git` | https://git-scm.com/downloads |

### Mouse Support

| Action | Result |
|--------|--------|
| Scroll wheel up | Scroll output up |
| Scroll wheel down | Scroll output down |

## 🎨 Themes

archiTerm comes with 4 built-in color themes:

| Theme | Description |
|-------|-------------|
| `dark` | Default dark theme with purple/cyan accents |
| `dracula` | Popular Dracula color scheme |
| `nord` | Nord-inspired arctic color palette |
| `gruvbox` | Retro groove color scheme |

### Using Themes

```bash
# Start with a specific theme
architerm --theme dracula

# List available themes
architerm themes

# Preview a theme's colors
architerm theme-preview nord
```

### Runtime Theme Switching

Press **`Ctrl+T`** while running to cycle through themes instantly! The current theme name will briefly appear in the status bar.

## 🔧 Configuration

archiTerm looks for custom commands in the following locations:

1. `~/.config/architerm/commands.yaml` (or `.json`)
2. Custom path via `--config` flag

### YAML Configuration Example

```yaml
commands:
  - template: "ssh user@hostname"
    description: "SSH to a remote server"
    category: "custom"
    tags:
      - ssh
      - remote

  - template: "terraform plan"
    description: "Show Terraform execution plan"
    category: "terraform"
    tags:
      - terraform
      - iac
```

### JSON Configuration Example

```json
{
  "commands": [
    {
      "template": "terraform apply",
      "description": "Apply Terraform changes",
      "category": "terraform",
      "tags": ["terraform", "iac"]
    }
  ]
}
```

## 📦 Built-in Commands

archiTerm comes with **230+ pre-loaded commands** stored as embedded JSON files, making it easy to extend:

| Tool | Icon | Commands | Examples |
|------|------|----------|----------|
| **Docker** | 🐳 | 15 | `docker ps`, `docker logs -f`, `docker-compose up` |
| **Kubernetes** | ☸️ | 15 | `kubectl get pods`, `kubectl describe`, `kubectl apply` |
| **gcloud** | ☁️ | 15 | `gcloud compute instances list`, `gcloud auth login` |
| **Azure** | ⚡ | 15 | `az login`, `az vm list`, `az aks get-credentials` |
| **curl** | 🌐 | 15 | `curl -X POST`, `curl -H "Authorization: Bearer"` |
| **git** | 📦 | 15 | `git status`, `git pull`, `git checkout -b` |
| **SSH** | 🔐 | 15 | `ssh USER@HOST`, `ssh -i KEY`, `ssh-keygen`, `scp` |
| **tcpdump** | 🔬 | 15 | `tcpdump -i any`, `tcpdump port 80`, `tcpdump -w file.pcap` |
| **netstat** | 📡 | 15 | `netstat -tuln`, `ss -tulnp`, `lsof -i :PORT` |
| **Linux** | 🐧 | 23 | `systemctl status`, `iptables -L`, `ufw allow`, `journalctl` |
| **nginx** | 🌿 | 15 | `nginx -t`, `nginx -s reload`, `tail -f access.log` |
| **conda** | 🐍 | 15 | `conda create -n env`, `conda activate`, `conda install` |
| **tmux** | 🖥️ | 15 | `tmux new -s`, `tmux attach`, `tmux split-window` |
| **grep** | 🔍 | 15 | `grep -r PATTERN`, `grep -i`, `grep -n`, `grep -v` |
| **find** | 📂 | 15 | `find . -name`, `find -type f`, `find -mtime`, `find -exec` |

### Command Storage

Commands are stored in JSON files under `internal/commands/embedded/`:
```
internal/commands/embedded/
├── docker.json      # Docker & docker-compose commands
├── kubernetes.json  # kubectl commands
├── gcloud.json      # Google Cloud CLI commands
├── azure.json       # Azure CLI commands
├── curl.json        # curl HTTP commands
├── git.json         # Git version control commands
├── ssh.json         # SSH, SCP, ssh-keygen commands
├── tcpdump.json     # Network packet capture commands
├── netstat.json     # netstat, ss, lsof network commands
├── linux.json       # systemctl, iptables, ufw, journalctl
├── nginx.json       # nginx web server commands
├── conda.json       # Conda environment management
├── tmux.json        # tmux terminal multiplexer
├── grep.json        # grep pattern search commands
└── find.json        # find file search commands
```

These files are embedded into the binary at compile time, so no external files are needed at runtime.

## 🏗️ Building from Source

### Requirements

- Go 1.21 or later
- Make (optional, for convenience)

### Build Commands

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Build for specific platform
make build-linux
make build-darwin
make build-windows

# Run tests
make test

# Create release archives
make release
```

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Adding New Built-in Commands

To add a new technology category with built-in commands:

1. Create a new JSON file in `internal/commands/embedded/` (e.g., `terraform.json`)
2. Follow the JSON structure:
```json
{
  "category": "terraform",
  "commands": [
    {
      "template": "terraform init",
      "description": "Initialize Terraform working directory",
      "tags": ["init", "setup"]
    },
    {
      "template": "terraform plan",
      "description": "Show execution plan",
      "tags": ["plan", "preview"]
    }
  ]
}
```
3. Rebuild the application: `make build`

The new category will automatically appear in the "Supported Technologies" panel!

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📋 Changelog

See [CHANGELOG.md](CHANGELOG.md) for a detailed history of changes.

## 🙏 Acknowledgments

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - The TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [zsh-autocomplete](https://github.com/marlonrichert/zsh-autocomplete) - Inspiration for the autocomplete UX

---

<p align="center">
  Made with ❤️ for architects and developers who love the terminal.
  <br><br>
  <a href="https://github.com/duladissa/architerm/stargazers">⭐ Star us on GitHub</a> •
  <a href="https://github.com/duladissa/architerm/issues">🐛 Report Bug</a> •
  <a href="https://github.com/duladissa/architerm/issues">✨ Request Feature</a>
</p>
