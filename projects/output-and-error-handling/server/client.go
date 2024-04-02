package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	resp, err := http.Get("http://localhost:8080") 
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error making HTTP request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	//checking conditions
	switch resp.StatusCode {
	case http.StatusOK:
		fmt.Println("Weather Information:")
		fmt.Println("Response Status:", resp.Status)
	case http.StatusTooManyRequests:
		fmt.Fprintf(os.Stderr, "Server too busy. Will retry after 5 seconds\n")
		time.Sleep(5 * time.Second) 
		
	default:
		fmt.Fprintf(os.Stderr, "Unexpected server response. Status code: %d\n", resp.StatusCode)
		os.Exit(1)
	}

	os.Exit(0)
}