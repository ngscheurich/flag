package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const SHORTCUT_API_URL = "https://api.app.shortcut.com/api/v3"

var cfgFile string

type ApiError struct {
	Message string `json:"message"`
	Tag     string `json:"tag"`
}

type Member struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	MentionName string `json:"mention_name"`
}

var rootCmd = &cobra.Command{
	Use:   "flag",
	Short: "⚑ FLAG is Shortcut on the command-line",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.flag.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".flag")
	}

	viper.AutomaticEnv()
}
