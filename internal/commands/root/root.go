/*
This is a file for a 'root' command. Which would be called immideately
after running the app.

TODO: add more description later
*/
package root

import (
	"fmt"
	"os"

	"github.com/aantoschuk/feed/internal/apperr"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-feed [command]",
	Short: "Collect and list news articles from configured sources.",
	Long: `Go-Feed is a CLI tool to collect and list recent news articles.

	This tool does not download full aricle content; it collects titles,
	links and metadata so you can quickly see what's new and decide what to read on
	the original site. It is designed to be fast, lightweight.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			fmt.Println("Hello")
			appErr := apperr.NewInternalError("cannot retrieve -u flag", "RETRIEVE_U_FLAG_EROR", 1, err)
			fmt.Println(appErr)
		}
		if url == "" {
			fmt.Println(apperr.ErrMissingRequiredFlag)
			os.Exit(1)
		}
		fmt.Println(url)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// i prefer to have a shorthand along with the full flag name
	rootCmd.Flags().StringP("url", "u", "", "Provide url to aggregate")
}
