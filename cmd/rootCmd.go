package cmd

import (
	"fmt"
	"foss_toolconverter/internal"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hello",
	Short: "This is the first command",
	Long: `A longer description 
	for the first command`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the first cobra example")
	},
}

var convertCmd = &cobra.Command{
	Use:     "convert",
	Short:   "Converts Tool Output",
	Long:    "Converts the Output of a Tool to a common structure",
	Aliases: []string{"con"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case args[0] == "syft":
			var tool internal.Syft
			tool.Convert()
		default:
			fmt.Println("No Tool specified")
		}

	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
