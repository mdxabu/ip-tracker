package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type IPInfo struct {
	IP        string `json:"ip"`
	City      string `json:"city"`
	Region    string `json:"region"`
	Country   string `json:"country"`
	Loc       string `json:"loc"`
	Postal    string `json:"postal"`
	Timezone  string `json:"timezone"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	MapURL    string `json:"map_url"`
}

func GetIPInfo(ip string) (string, error) {
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

	// Extract latitude and longitude from the 'loc' field
	coords := strings.Split(info.Loc, ",")
	if len(coords) == 2 {
		info.Latitude = coords[0]
		info.Longitude = coords[1]

		// Generate a map URL using Google Maps
		info.MapURL = fmt.Sprintf("https://www.google.com/maps?q=%s,%s", info.Latitude, info.Longitude)
	}

	// Return the formatted IP info with map URL
	return fmt.Sprintf("IP: %s\nCity: %s\nRegion: %s\nCountry: %s\nLocation: %s\nPostal: %s\nTimezone: %s\nMap URL: %s",
		info.IP, info.City, info.Region, info.Country, info.Loc, info.Postal, info.Timezone, info.MapURL), nil
}
