package cmd

import (
	

	"github.com/spf13/cobra"	
)

var (
	
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "ip-tracker cli tool",
		Long:  `ip-tracker cli tool`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

