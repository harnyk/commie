# CLI Application

This is a Command Line Interface (CLI) application written in Go, utilizing the Cobra and Viper libraries for command handling and configuration management.

## Features
- **Command Line Interface**: Built using Cobra to handle CLI interactions.
- **Configuration Management**: Utilizes Viper for flexible configuration handling through files and environment variables.

## Configuration
The application makes use of a configuration file (`config.toml`) or environment variables to manage its settings. The following configuration options are available:
- `OPENAI_KEY`: API key for accessing the OpenAI service.
- `OPENAI_MODEL`: The model to be used with OpenAI services.

### Configuring via Environment Variables
You can set the following environment variables to configure the application:
- `OPENAI_KEY`: Your API key for OpenAI.
- `OPENAI_MODEL`: The OpenAI model you wish to use.

## Getting Started

1. **Clone the repository**:
   ```sh
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Build the application**:
   ```sh
   go build -o cli-app
   ```

3. **Run the application**:
   ```sh
   ./cli-app
   ```

## Dependencies
- Go

This application requires Go to be installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).

## License
This project is licensed under the MIT License.