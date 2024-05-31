package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove saved Shortcut token and user data",
	Run: func(cmd *cobra.Command, args []string) {
		var member Member

		viper.ReadInConfig()
		viper.Set("shortcut_token", "")
		viper.Set("current_user", member)
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
