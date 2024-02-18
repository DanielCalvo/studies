package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const people = `joe
max
bob
alice`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "Changed this on root -- is this where the help command goes?",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Unable to get working directory")
		os.Exit(1)
	}

	var filePath string
	rootCmd.PersistentFlags().StringVar(&filePath, "filePath", pwd+string(os.PathSeparator)+"people.txt", "File in which a list of people's names are stored")
}
