package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "FOSS_toolconverter",
	Short: "Converts the Output of a tool to usable dependency format",
}

var convertCmd = &cobra.Command{
	Use:   "convert [path]",
	Short: "Convert file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		inputPath := args[0]
		npmPath, _ := cmd.Flags().GetString("npm")
		if npmPath != "" {
			if npmPath == "true" {
				return fmt.Errorf("npm flag requires a path argument")
			}
			fmt.Println("Converting with npm behavior using package.json file:", npmPath)
		} else {
			fmt.Println("Converting without npm behavior")
		}
		fmt.Println("Path to input file:", inputPath)
		return nil
	},
}

func init() {
	convertCmd.Flags().String("npm", "", "Path to package.json file")
	rootCmd.AddCommand(convertCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
