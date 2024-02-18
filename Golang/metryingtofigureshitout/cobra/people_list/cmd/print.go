/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("print called")
		filePath, err := rootCmd.PersistentFlags().GetString("filePath")
		if err != nil {
			return err
		}

		fileContents, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return err
		}
		fmt.Println(string(fileContents))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
