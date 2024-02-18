package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generates a file named people.txt that has one name per line on it",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filePath, err := rootCmd.PersistentFlags().GetString("filePath")
		if err != nil {
			return err
		}

		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.WriteString(people)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
