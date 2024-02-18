package cmd

import (
	"bufio"
	"errors"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("You need to pass the name of at least one person to remove from the list")
		}
		filePath, err := rootCmd.PersistentFlags().GetString("filePath")
		if err != nil {
			return err
		}

		fileContents, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(strings.NewReader(string(fileContents)))

		var result []string

	scannerLoop:
		for scanner.Scan() {
			for _, arg := range args {
				if arg == scanner.Text() {
					continue scannerLoop
				}
			}
			result = append(result, scanner.Text())

		}

		resultString := strings.Join(result, "\n")
		err = os.WriteFile(filePath, []byte(resultString), 0644)
		if err != nil {
			return err
		}
		return nil

	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
