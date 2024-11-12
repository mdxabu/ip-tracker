package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/mdxabu/ipscout/pkg"
)

func ipInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Extract IP from URL path: /geo/{ip-address}
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 3 {
		http.Error(w, "Invalid request format. Use /geo/{ip-address}", http.StatusBadRequest)
		return
	}

	ip := segments[2]
	info, err := pkg.GetIPInfo(ip)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving information for IP %s: %v", ip, err), http.StatusInternalServerError)
		return
	}

	// Return the geolocation data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func StartWebServer() {
	http.HandleFunc("/geo/", ipInfoHandler)

	// Start the server on port 8080
	log.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
