/*
This is a file for a 'root' command. Which would be called immideately
after running the app.

TODO: add more description later
*/
package root

import (
	"fmt"
	"os"

	"github.com/aantoschuk/feed/internal/app_logger"
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
		logger := app_logger.NewAppLogger(false)
		flags, err := retrieveFlags(cmd)
		if err != nil {
			appErr := apperr.NewInternalError("cannot retrieve -u flag", "RETRIEVE_U_FLAG_EROR", 1, err)
			logger.Error(appErr)
		}
		logger.SetVerbose(flags.v)
		logger.Info("executing root command")

		if err != nil {
			logger.Error(err)
		}

		fmt.Println(flags.u)
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
	t := false
	rootCmd.PersistentFlags().BoolVarP(&t, "verbose", "v", false,
		"Enables verbose mode in the app. Which displays all the messages with the full error information.")
	// i prefer to have a shorthand along with the full flag name
	rootCmd.Flags().StringP("url", "u", "", "Provide url to aggregate")
}
