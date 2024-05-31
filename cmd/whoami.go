package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Print information about current user",
	Run: func(cmd *cobra.Command, args []string) {
		viper.ReadInConfig()
		token := viper.GetString("shortcut_token")

		if token == "" {
			log.Fatal("Not logged in")
		}

		var currentUser Member

		viper.ReadInConfig()
		viper.UnmarshalKey("current_user", &currentUser)

		userJson, _ := json.Marshal(currentUser)
		fmt.Println(string(userJson))
	},
}

func init() {
	rootCmd.AddCommand(whoamiCmd)
}
