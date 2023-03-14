package cmd

import (
	"fmt"
	"foss_toolconverter/internal"
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
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		inputPath := args[0]
		outputPath := args[1]
		npmPath, _ := cmd.Flags().GetString("npm")
		var manager internal.Manager
		if npmPath != "" {
			//If npmPath wasnt set it will be set to true by cobra so we can check for that
			if npmPath == "true" {
				return fmt.Errorf("npm flag requires a path argument")
			}
			manager.SyftToDep(inputPath, outputPath, true, npmPath)
		} else {
			manager.SyftToDep(inputPath, outputPath, false, npmPath)
		}
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
