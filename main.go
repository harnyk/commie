package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/harnyk/commie/cmd/playground/tools/cat"
	"github.com/harnyk/commie/cmd/playground/tools/dump"
	"github.com/harnyk/commie/cmd/playground/tools/git"
	"github.com/harnyk/commie/cmd/playground/tools/ls"
	"github.com/harnyk/commie/cmd/playground/tools/rm"
	"github.com/harnyk/commie/pkg/agent"
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

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
	}

	viper.AutomaticEnv() // Load from environment variables

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

func main() {
	rootCmd := &cobra.Command{
		Use:   "cli-app",
		Short: "A CLI app with configuration",
		Long:  "An example CLI application demonstrating Cobra and Viper for configuration.",
		Run:   commitCmd.Run, // Default command
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.toml)")

	rootCmd.AddCommand(commitCmd, helpCmd, chatCmd)

	cobra.OnInitialize(initConfig)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Default commit command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Commit executed with model: %s and key: %s\n", cfg.OpenAIModel, cfg.OpenAIKey)
	},
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
		inforg := agent.NewAgent().
			WithOpenAIKey(cfg.OpenAIKey).
			WithOpenAIModel(cfg.OpenAIModel).
			WithSystemPrompt(`
				You are a helpful assistant which helps a user to work with the file system, terminal and git.
				Your responses will be rendered directly to the modern Linux terminal,
				so you should use ASCII art, emojis for icons, ASCII terminal codes for colors.
				Markdown is not allowed, if you use it, the whole response will be broken.
				Reply with just a plain text with no markdown.

				If the user asks to do something, you should do your best and provide deep analysis using the
				available tools.

				If you compose commit messages, you should
				 - analyze the changes
				 - read the git diffs
				 - if necessary, read through the sources
				 - reason about the changes
				 - compose a concise commit message as a summary of the changes in "conventional commits" format.

				If you are asked to write some file, first, read it until the end, and only then incorporate changes
			`).
			WithTool(ls.New()).
			WithTool(cat.New()).
			WithTool(git.NewStatus()).
			WithTool(git.NewDiff()).
			WithTool(git.NewCommit()).
			WithTool(git.NewAdd()).
			WithTool(git.NewPush()).
			WithTool(dump.New()).
			WithTool(rm.New()).
			Build()

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

			fmt.Println("")
			fmt.Println("-------------------------------------")
			fmt.Println("> Question: ", question)
			fmt.Println("-------------------------------------")
			fmt.Println("Answer:", answer)
		}
	},
}
