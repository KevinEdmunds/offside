package fpl

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchBootstrap(client *http.Client, fplBase string) (*Bootstrap, error) {

	resp, err := client.Get(fplBase + "/bootstrap-static/")
	if err != nil {
		return nil, fmt.Errorf("fetch bootstrap: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var bootstrap Bootstrap
	if err := json.NewDecoder(resp.Body).Decode(&bootstrap); err != nil {
		return nil, fmt.Errorf("decode bootstrap: %w", err)
	}

	return &bootstrap, nil
}