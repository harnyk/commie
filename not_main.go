package main

import (
	"fmt"
	"os"

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

	rootCmd.AddCommand(commitCmd, helpCmd)

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
