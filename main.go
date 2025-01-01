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
	"github.com/harnyk/commie/pkg/tools/dump"
	"github.com/harnyk/commie/pkg/tools/git"
	"github.com/harnyk/commie/pkg/tools/list"
	"github.com/harnyk/commie/pkg/tools/ls"
	"github.com/harnyk/commie/pkg/tools/memory"
	"github.com/harnyk/commie/pkg/tools/patch"
	"github.com/harnyk/commie/pkg/tools/rm"
	"github.com/harnyk/gena"
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
	default: // Линукс и другие
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
func createAgent() *gena.Agent {
	memoryRepo := memory.NewMemoryRepoYAMLFile("./.commie/memory.yaml")

	promptTextWithMemory := strings.Builder{}
	promptTextWithMemory.WriteString(promptText)

	toc, _ := memoryRepo.GetTOC()
	if len(toc) > 0 {
		// promptTextWithMemory = promptTextWithMemory + "\nCurrent memory item ids: "
		promptTextWithMemory.WriteString("\nCurrent memory items:\n")
		for _, item := range toc {
			tagsString := strings.Join(item.Tags, ",")
			promptTextWithMemory.WriteString(fmt.Sprintf("- id:'%s', tags:%s\n", item.ID, tagsString))
		}
	}

	fmt.Println(promptTextWithMemory.String())

	return gena.NewAgent().
		WithOpenAIKey(cfg.OpenAIKey).
		WithOpenAIModel(cfg.OpenAIModel).
		WithSystemPrompt(promptTextWithMemory.String()).
		WithTool(ls.New()).
		WithTool(list.New()).
		WithTool(rm.New()).
		WithTool(dump.New()).
		WithTool(patch.New()).
		WithTool(git.NewStatus()).
		WithTool(git.NewDiff()).
		WithTool(git.NewCommit()).
		WithTool(git.NewPush()).
		WithTool(git.NewAdd()).
		WithTool(git.NewLog()).
		WithTool(memory.NewSet(memoryRepo)).
		WithTool(memory.NewGet(memoryRepo)).
		Build()
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "cli-app",
		Short: "A CLI app with configuration",
		Long:  "An example CLI application demonstrating Cobra and Viper for configuration.",
		Run:   chatCmd.Run, // Default command
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
		cmd.Help() // Show help for the root command
	},
}

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start the chat session",
	Run: func(cmd *cobra.Command, args []string) {
		inforg := createAgent()

		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Enter your question: ")
			question, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			answer, err := inforg.Ask(context.Background(), question)
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
