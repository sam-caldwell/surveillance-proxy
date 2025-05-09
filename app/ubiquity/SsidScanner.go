package ubiquity

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Configuration for UniFi OS token-based access
const (
	BaseURL     = "https://%s/proxy/network"
	AccessToken = "eyJhbGciOi..." // Paste your valid token here
	OutputFile  = "ssids.txt"
)

// RFScanEntry represents a single SSID scan result
type RFScanEntry struct {
	SSID    string `json:"ssid"`
	BSSID   string `json:"bssid"`
	Channel int    `json:"channel"`
	RSSI    int    `json:"rssi"`
	Band    string `json:"band"`
}

// CollectSsidData - queries the given device API for all discovered SSID information
func CollectSsidData(address, token *string) {

	var (
		data []RFScanEntry
		err  error
	)

	if data, err = fetchRFScan(address); err != nil {
		fmt.Println("RF scan fetch failed:", err)
		return
	}

	if err = writeToFile(data); err != nil {
		fmt.Println("File write failed:", err)
	}
}

// fetchRFScan performs an authenticated request using the bearer token.
func fetchRFScan(address *string) ([]RFScanEntry, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(BaseURL, *address)+"/api/s/default/stat/rf-scan", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data []RFScanEntry `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

// writeToFile writes the scan entries to a plain text file.
func writeToFile(entries []RFScanEntry) error {
	f, err := os.Create(OutputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, entry := range entries {
		line := fmt.Sprintf("SSID: %s, BSSID: %s, RSSI: %d, Channel: %d, Band: %s\n",
			entry.SSID, entry.BSSID, entry.RSSI, entry.Channel, entry.Band)
		_, err := io.WriteString(f, line)
		if err != nil {
			return err
		}
	}
	return nil
}
