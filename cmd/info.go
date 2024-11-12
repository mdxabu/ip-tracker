
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type IPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Information about the IP address",
	Long:  `Information about the IP address`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide an IP address")
			return
		}
		ip := args[0]
		info, err := getIPInfo(ip)
		if err != nil {
			fmt.Printf("Error retrieving information for IP %s: %v\n", ip, err)
			return
		}
		fmt.Printf("Information for IP %s: %v\n", ip, info)
	},
}

func getIPInfo(ip string) (string, error) {
	url := fmt.Sprintf("https://ipinfo.io/%s/json", ip)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get IP info: %s", resp.Status)
	}

	var info IPInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return "", err
	}

	return fmt.Sprintf("IP: %s\nCity: %s\nRegion: %s\nCountry: %s\nLocation: %s\nPostal: %s\nTimezone: %s",
		info.IP, info.City, info.Region, info.Country, info.Loc, info.Postal, info.Timezone), nil
	
}

func init() {
	rootCmd.AddCommand(infoCmd)

	
}
