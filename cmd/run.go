package cmd

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spekulant/hdd-files-recovery/runner"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs through all files in the specified location",
	Long: `Using file extensions specified in the config.json file
			it will run through all files in the specified location 
			looking only for files ending with "lookFor" extensions 
			and omitting files whose names are covered by the "filterOut"
			pattern`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Print("Please supply a root path to go through")
			os.Exit(1)
		}
		if err := runner.Run(args[0]); err != nil {
			log.Printf("%s\n", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
