package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var query string

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for stories",
	Run: func(cmd *cobra.Command, args []string) {
		viper.ReadInConfig()
		token := viper.GetString("shortcut_token")

		if token == "" {
			log.Fatal("Not logged in")
		}

		client := http.Client{}

		requestURL := fmt.Sprintf("%s/search/stories", SHORTCUT_API_URL)

		if query == "owner:<you>" {
			query = "owner:nscheurich"
		}

		queryParams := url.Values{}
		queryParams.Set("query", query)

		req, _ := http.NewRequest(http.MethodGet, requestURL, nil)

		req.URL.RawQuery = queryParams.Encode()
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Shortcut-Token", token)

		res, _ := client.Do(req)
		body, _ := io.ReadAll(res.Body)
		fmt.Println(string(body))
	},
}

func init() {
	searchCmd.Flags().StringVarP(&query, "query", "q", "owner:<you>", "search query")
	rootCmd.AddCommand(searchCmd)
}
