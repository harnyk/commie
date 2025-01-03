package main

import (
	"bufio"
	"context"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/harnyk/commie/pkg/banner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	OpenAIKey   string `mapstructure:"OPENAI_KEY"`
	OpenAIModel string `mapstructure:"OPENAI_MODEL"`
}

var (
	cfg     Config
	cfgFile string
)

//go:embed commie.prompt.md
var promptText string

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

	viper.AutomaticEnv() // Загрузка из переменных окружения

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
	rootCmd := &cobra.Command{
		Use:   "commie",
		Short: "An AI-powered CLI tool",
		Run:   chatCmd.Run,
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is the OS-specific config path)")

	rootCmd.AddCommand(helpCmd, chatCmd)

	cobra.OnInitialize(initConfig)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Displays help information",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start the chat session",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(banner.GetBanner())

		agent := createAgent()

		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print(">>>: ")
			question, err := reader.ReadString('\n')
			if err != nil {
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

			answerRendered := string(markdown.Render(answer, 80, 0))

			fmt.Println("")
			fmt.Println("-------------------------------------")
			fmt.Println("> ", question)
			fmt.Println("-------------------------------------")
			fmt.Println(answerRendered)
		}
	},
}
