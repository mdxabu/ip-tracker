package cmd

import (
	

	"github.com/spf13/cobra"	
)

var (
	
	rootCmd = &cobra.Command{
		Use:   "ipscout",
		Short: "ipscout cli tool",
		Long:  `ipscout cli tool`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

