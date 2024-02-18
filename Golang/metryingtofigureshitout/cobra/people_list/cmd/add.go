package cmd

import (
	"bufio"
	"errors"
	"github.com/spf13/cobra"
	"os"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("You need to pass the name of at least one person to add to the list")
		}
		filePath, err := rootCmd.PersistentFlags().GetString("filePath")
		if err != nil {
			return err
		}

		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		writer := bufio.NewWriter(f)

		for _, person := range args {
			writer.WriteString(person + "\n")
			if err != nil {
				return err
			}
		}
		err = writer.Flush()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
