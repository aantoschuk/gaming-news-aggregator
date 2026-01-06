/*
This is a file for a 'root' command. Which would be called immideately
after running the app.

TODO: add more description later
*/
package root

import (
	"fmt"
	"os"

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
		fmt.Println("Root command")
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
	// initialize flags here
}
