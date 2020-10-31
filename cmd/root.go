package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"

	cli "github.com/spf13/cobra"
	config "github.com/spf13/viper"
)

var (

	// Config and global logger
	configFile string
	pidFile    string

	// Root handle for other commands
	rootCmd = &cli.Command{
		Version: "1.0",
		Use:     "hdd-files-recovery",
		PersistentPreRunE: func(cmd *cli.Command, args []string) error {
			// Create Pid File
			pidFile = config.GetString("pidfile")
			if pidFile != "" {
				file, err := os.OpenFile(pidFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
				if err != nil {
					log.Error().Msgf("Could not create pid file: %s Error:%v", pidFile, err)
					return err
				}
				defer file.Close()
				_, err = fmt.Fprintf(file, "%d\n", os.Getpid())
				if err != nil {
					log.Error().Msgf("Could not create pid file: %s Error:%v", pidFile, err)
					return err
				}
			}
			return nil
		},
		PersistentPostRun: func(cmd *cli.Command, args []string) {
			// Remove Pid file
			if pidFile != "" {
				os.Remove(pidFile)
			}
		},
	}
)

// Execute is the entrypoint for cmdline execution
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("%s\n", err.Error())
	}
}
