package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Define the nocopy flag with two names
	nocopy := flag.Bool("nocopy", false, "Do not copy GUIDs to clipboard")
	nocopyShort := flag.Bool("n", false, "Do not copy GUIDs to clipboard (short)")

	// Custom usage function to show only the executable name
	flag.Usage = func() {
		exeName := filepath.Base(os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [-nocopy | -n] [-h]\n", exeName)
		flag.PrintDefaults()
	}

	flag.Parse()

	// Define the URL of the external service to get the public IP
	url := "https://api.ipify.org?format=text"

	// Make an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching IP address: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
		os.Exit(1)
	}

	ip := string(body)

	// Copy the IP address to the clipboard using the clip command if neither -nocopy nor -n is specified
	if !*nocopy && !*nocopyShort {
		cmd := exec.Command("cmd", "/c", "echo|set /p="+ip+"|clip")
		cmd.Run()
	}

	// Print the IP address
	fmt.Printf("%s\n", ip)
}
