package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Send a GET request to the URL
	resp, err := http.Get("https://ifconfig.me")
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body) // Using io.ReadAll instead of ioutil.ReadAll
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	// Print the response body (IP address or other info)
	fmt.Println(string(body))
}
