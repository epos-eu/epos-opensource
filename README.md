# EPOS Open Source CLI

A command-line tool for deploying the EPOS Platform locally using Docker Compose or Kubernetes.

![Image](https://github.com/user-attachments/assets/adb46bfd-b5b1-47d8-9c56-4aa7cfb24479)

---

## 🚀 Quick Start

**Install the CLI (Linux/macOS/WSL):**

```bash
curl -fsSL https://raw.githubusercontent.com/epos-eu/epos-opensource/main/install.sh | bash
```

**Check installation:**

```bash
epos-opensource --version
```

**Deploy a local environment with Docker Compose:**

```bash
epos-opensource docker deploy myenv
```

**Populate it with your own data:**

```bash
epos-opensource docker populate myenv /path/to/ttl/files
```

---

## What is EPOS Open Source CLI?

A command-line tool to easily deploy and manage EPOS Platform environments on your computer or in the cloud, using either Docker Compose or Kubernetes.

---

## What is an "Environment"?

An "environment" is a named, isolated instance of the EPOS Platform, with its own configuration and data. You can have multiple environments for testing, development, etc.

---

## What is a "TTL file"?

A TTL file is a metadata file in [Turtle format](https://www.w3.org/TR/turtle/), used to describe datasets for EPOS. You can find or create these files to load your own data and visualize it in the GUI.

---

## Usage Requirements

- **Docker** and **Docker Compose** (for Docker-based setups)
- **kubectl** and access to a Kubernetes cluster (for Kubernetes-based setups)
- **Internet connection** (for downloading images and updates)

---

## Installation

### Easiest: Installation Script

```bash
curl -fsSL https://raw.githubusercontent.com/epos-eu/epos-opensource/main/install.sh | bash
```

_This script works on Linux, macOS, and WSL. It will update the CLI if already installed._

> [!IMPORTANT]
> ALWAYS REVIEW SCRIPTS YOURSELF BEFORE EXECUTING THEM ON YOUR SYSTEM!

### Using Go

```shell
go install github.com/epos-eu/epos-opensource@latest
```

_Make sure `$GOPATH/bin` or `$HOME/go/bin` is in your `$PATH`._

### Pre-built Binaries

1. Download the appropriate archive from the [releases](https://github.com/epos-eu/epos-opensource/releases).
2. Make the binary executable and move it to your `$PATH`:

   ```shell
   chmod +x epos-opensource-{your version}
   mv epos-opensource-{your version} /usr/local/bin/epos-opensource
   ```

3. Verify the installation:

   ```shell
   epos-opensource --version
   ```

### Build from source

Build using `make` with Go 1.24.4 or later:

```shell
make build
```

---

## Discovering All Commands

The CLI has many commands and options. To see everything you can do, use the built-in help:

```bash
epos-opensource --help
epos-opensource docker --help
epos-opensource kubernetes --help
epos-opensource docker deploy --help
```

This will always show the most up-to-date list of commands and flags.

---

## Usage Examples

### List available commands

```shell
epos-opensource --help
```

### Docker example: Deploy and populate an environment

```shell
epos-opensource docker deploy myenv
epos-opensource docker populate myenv /path/to/ttl-files
```

After deploying, the CLI will print the URLs for:

- EPOS Data Portal
- EPOS API Gateway
- EPOS Backoffice

Look for these in your terminal output.

---

## Troubleshooting & Tips

- **Docker/Kubernetes not found:** Make sure Docker and/or kubectl are installed and running.
- **Environment/Directory already exists:** Use a new name, or delete the old environment first.
- **Problems with `.ttl` files:** Make sure the directory exists and contains valid `.ttl` files and that their path are valid (no spaces, weird symbols, ...).

If you get stuck, run with `--help` for more info, or feel free to [open an issue](https://github.com/epos-eu/epos-opensource/issues).

---

## Development

Follow these steps to set up your local development environment and enable the shared Git hooks.

### Prerequisites

- **Go 1.24.4+**
- **Make** (GNU Make)
- **golangci‑lint**

### Clone & Enter the Repo

```bash
git clone https://github.com/epos-eu/epos-opensource.git
cd epos-opensource
```

### Install the Shared Git Hooks

We include a set of pre‑commit hooks under `.githooks/` that will automatically run your Makefile checks before each commit:

```bash
git config core.hooksPath .githooks
```

### Makefile Targets

Our `Makefile` provides common commands for development like:

- **`make build`**
  Compile the CLI binary.

- **`make test`**
  Run all tests.

- **`make lint`**
  Execute linters (using `golangci-lint`).

### Workflow

1. **Edit code** in your favorite editor.
2. **Run** your chosen `make` targets to verify everything passes.
3. **Commit:** the pre‑commit hook will automatically invoke `make lint test`.
4. **Push** your changes when you're ready.

---

## Contributions Welcome

We welcome all contributions, including bug reports, feature ideas, documentation, or code changes.

If you have questions or are unsure how to get started, feel free to [open an issue](https://github.com/epos-eu/epos-opensource/issues). We are happy to assist!

## Contributing

1. Fork the repository and clone it locally.
2. Create a new branch from `main` or `develop`.
3. Make your changes and follow the existing code style.
4. Add or update tests and documentation as needed.
5. Commit your changes using clear, present-tense messages.
6. Push your branch and open a pull request.

Before submitting:

- Pull the latest changes from upstream.
- Squash commits if needed.
- Ensure all tests pass.

After submitting:

- Address any feedback on your pull request.
- Update your branch until it is approved and merged.

Thank you for your contribution!

```

```
