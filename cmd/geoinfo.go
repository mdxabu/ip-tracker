package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/mdxabu/ipscout/pkg"
)

var geoCmd = &cobra.Command{
	Use:   "geo",
	Short: "Geolocation information about the IP address",
	Long:  `Retrieve detailed geolocation information about the given IP address, including its city, country, region, timezone, and a map link.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide an IP address")
			return
		}
		ip := args[0]
		info, err := pkg.GetIPInfo(ip)
		if err != nil {
			fmt.Printf("Error retrieving information for IP %s: %v\n", ip, err)
			return
		}
		// Print the IP geolocation info
		fmt.Println(info)
	},
}

func init() {
	rootCmd.AddCommand(geoCmd)
}
