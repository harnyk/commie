# Commie

Commie is a command-line interface (CLI) application developed in Go, aimed at providing a versatile tool for interacting with the filesystem, terminal, and Git operations. It utilizes the powerful Cobra library for managing commands and Viper for handling configuration settings.

## Features

- **Configurable via File or Environment**: Supports configuration using a `config.toml` file or environment variables.
- **Commands**:
  - `commit`: Executes a commit operation utilizing an OpenAI model and key for intelligent suggestions.
  - `help`: Returns assistance and documentation on command usage.
  - `chat`: Starts an interactive session allowing the CLI to act as an informative assistant.

## Configuration

Commie can be configured with:

- **Configuration File**: `config.toml`
- **Environment Variables**:
  - `OPENAI_KEY`
  - `OPENAI_MODEL`

Ensure these are set up correctly for the application to run smoothly, especially the `chat` and `commit` functionalities which rely on OpenAI services.

## Installation

1. **Dependencies**:
   - Ensure all required Go packages and dependencies are installed, including Cobra, Viper, and relevant tool packages for functionality like `ls`, `cat`, `git`, etc.

2. **Build**:
   - Run `go build` to compile the application.

## Assistant Behavior

- Before executing Git operations such as 'commit' or 'push', the assistant will always ask for user confirmation to ensure control over changes.