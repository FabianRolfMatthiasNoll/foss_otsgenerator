package cmd

import (
	"foss_otsgenerator/internal"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "foss_toolconverter",
	Short: "Converts the Output of a tool to usable dependency format",
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "fosser_otsgenerator init [Path/To/SBOM] [Path/To/Config] [Path/For/Output]",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]
		configPath := args[1]
		outputPath := args[2]
		var manager internal.Manager
		manager.GenerateOTS(inputPath, configPath, outputPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
