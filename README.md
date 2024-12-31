# Commie

Commie is a command-line interface (CLI) application developed in Go, aimed at providing a versatile tool for interacting with the filesystem, terminal, and Git operations. It utilizes the powerful Cobra library for managing commands and Viper for handling configuration settings.

## Features

- **Configurable via File or Environment**: Supports configuration using a `config.toml` file or environment variables.
- **Commands**:
  - `help`: Returns assistance and documentation on command usage.
  - `chat`: Starts an interactive session allowing the CLI to act as an informative assistant.

## Configuration

Commie can be configured with:

- **Configuration File**: `config.toml`
- **Environment Variables**:
  - `OPENAI_KEY`
  - `OPENAI_MODEL`

Ensure these are set up correctly for the application to run smoothly, especially the `chat` functionalities which rely on OpenAI services.

## Installation

Simply run `eget harnyk/commie` to install it.
