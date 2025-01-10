package main

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	// Start the application in a goroutine
	go func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Application failed to start: %v", r)
			}
		}()
		main()
	}()

	// Allow some time for the server to start
	time.Sleep(2 * time.Second)

	// Test the /metrics endpoint
	resp, err := http.Get("http://localhost:8080/metrics")
	if err != nil {
		t.Fatalf("Failed to connect to the /metrics endpoint: %v", err)
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	// Check response body (optional: validate specific metrics content)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %v", err)
	}

	if len(body) == 0 {
		t.Errorf("Metrics endpoint returned an empty response")
	}

	t.Log("Metrics endpoint tested successfully")
}
