# Local Development Setup

A quick-start guide for local development with Go CLI tools, FluxCD, and Kubernetes using Kind and Mise.

## Prerequisites
- [podman](https://podman.io/get-started)/[orbstack](https://orbstack.dev/) installed and running

### Optional
- [Mise](https://mise.jdx.dev/) for tool version management 

### Initial Setup

#### 1. Install Mise
For more detail information around mise and how to get started visit the projects [getting started page](https://mise.jdx.dev/getting-started.html)
```bash
# macOS
brew install mise

# Linux
curl https://mise.run | sh

# Add to your shell profile (assuming it's installed in .local/bin)
#fish
echo '~/.local/bin/mise activate fish | source' >> ~/.config/fish/config.fish
# bash
echo 'eval "$(mise activate bash)"' >> ~/.bashrc
# or for zsh
echo 'eval "$(mise activate zsh)"' >> ~/.zshrc
```

#### 2. Configure Tool Versions

Create a `.mise.toml` file in your project root:

```toml
[tools]
golang = "1.24"
kubectl = "1.34"
kind = "0.30"
flux = "2.6"
helm = "3.18"
```

Install the tools:

```bash
mise install
```

