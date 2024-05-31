package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var loginCmd = &cobra.Command{
	Use:   "login [token]",
	Short: "Log in using a Shortcut API token",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		client := http.Client{}

		requestURL := fmt.Sprintf("%s/member", SHORTCUT_API_URL)

		req, _ := http.NewRequest(http.MethodGet, requestURL, nil)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Shortcut-Token", token)

		res, _ := client.Do(req)
		body, _ := io.ReadAll(res.Body)

		var apiError ApiError
		json.Unmarshal([]byte(body), &apiError)

		if apiError.Tag == "" {
			var member Member
			json.Unmarshal([]byte(body), &member)

			viper.ReadInConfig()
			viper.Set("shortcut_token", token)
			viper.Set("current_user", member)
			viper.WriteConfig()
		} else {
			log.Fatal(apiError.Message)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
