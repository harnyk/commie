package main

import (
	"context"
	_ "embed"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/harnyk/commie/pkg/banner"
	"github.com/harnyk/commie/pkg/colorlog"
	"github.com/harnyk/commie/pkg/profile"
	"github.com/harnyk/commie/pkg/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	OpenAIKey    string `mapstructure:"OPENAI_KEY"`
	OpenAIModel  string `mapstructure:"OPENAI_MODEL"`
	OpenAIAPIURL string `mapstructure:"OPENAI_API_URL"`
	LogLevel     string `mapstructure:"LOG_LEVEL"`
}

var (
	version     = "development"
	cfg         Config
	cfgFile     string
	fileFlag    string
	commandFlag string
)

//go:embed commie.prompt.md
var promptText string

func strToLevel(level string) slog.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func getConfigDir() string {
	var configDir string

	switch runtime.GOOS {
	case "windows":
		configDir = os.Getenv("APPDATA")
	case "darwin":
		configDir = filepath.Join(os.Getenv("HOME"), "Library", "Application Support")
	default:
		configDir = os.Getenv("XDG_CONFIG_HOME")
		if configDir == "" {
			configDir = filepath.Join(os.Getenv("HOME"), ".config")
		}
	}
	return filepath.Join(configDir, "commie")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		configDir := getConfigDir()
		viper.AddConfigPath(configDir)
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
	}

	viper.AutomaticEnv()

	viper.SetDefault("LOG_LEVEL", "WARN")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Println("Unable to decode into struct:", err)
		os.Exit(1)
	}
}

// Function to create a new agent with the necessary configurations

func main() {
	helpCmd := &cobra.Command{
		Use:   "help",
		Short: "Displays help information",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	chatCmd := &cobra.Command{
		Use:   "chat",
		Short: "Start the chat session",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(banner.GetBanner())

			logLevel := strToLevel(cfg.LogLevel)
			log := slog.New(colorlog.NewColorConsoleHandler(os.Stderr, logLevel))

			profileResolver := profile.New(log)
			profileDir, err := profileResolver.Get()
			if err != nil {
				log.Error("failed to get profile dir", "error", err)
				os.Exit(1)
			}
			log.Debug("profile dir", "path", profileDir)

			agent := createAgent(
				profileDir,
				log,
			)

			if fileFlag != "" || commandFlag != "" {
				var commandFile string

				if commandFlag != "" {
					commandFile = filepath.Join(profileDir, "commands", commandFlag+".md")
				} else {
					commandFile = fileFlag
				}

				log.Debug("command file", "path", commandFile)

				content, err := os.ReadFile(commandFile)
				if err != nil {
					fmt.Println("Error reading command file:", err)
					return
				}

				question := string(content)

				answer, err := agent.Ask(context.Background(), question)
				if err != nil {
					fmt.Println("Error processing question:", err)
					return
				}
				answerRendered := ui.RenderMarkdown(answer)
				fmt.Println(answerRendered)
			}

			// reader := bufio.NewReader(os.Stdin)
			for {
				// fmt.Print(">>> ")
				question, err := ui.TextInput()
				if err != nil {
					if err == huh.ErrUserAborted {
						os.Exit(0)
					}
					fmt.Println("Error reading input:", err)
					continue
				}
				question = strings.TrimSpace(question)
				if question == "" {
					continue
				}
				answer, err := agent.Ask(context.Background(), question)
				if err != nil {
					fmt.Println("Error processing question:", err)
					continue
				}
				answerRendered := ui.RenderMarkdown(answer)
				fmt.Println("")
				fmt.Println("-------------------------------------")
				fmt.Println("> ", question)
				fmt.Println("-------------------------------------")
				fmt.Println(answerRendered)
			}
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Commie",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Commie version:", version)
		},
	}

	rootCmd := &cobra.Command{
		Use:   "commie",
		Short: "An AI-powered CLI tool",
		Run:   chatCmd.Run,
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is the OS-specific config path)")

	chatCmd.Flags().StringVarP(&fileFlag, "file", "f", "", "file with user task")
	rootCmd.Flags().StringVarP(&fileFlag, "file", "f", "", "file with user task")
	chatCmd.Flags().StringVarP(&commandFlag, "command", "c", "", "command")
	rootCmd.Flags().StringVarP(&commandFlag, "command", "c", "", "command")

	rootCmd.AddCommand(helpCmd, chatCmd, versionCmd)

	cobra.OnInitialize(initConfig)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
