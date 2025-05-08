package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Define flags
	celsiusFlag := flag.Bool("c", false, "Display temperature in Celsius")
	flag.BoolVar(celsiusFlag, "celsius", false, "Display temperature in Celsius")
	fahrenheitFlag := flag.Bool("f", false, "Display temperature in Fahrenheit (default)")
	flag.BoolVar(fahrenheitFlag, "fahrenheit", false, "Display temperature in Fahrenheit (default)")
	kelvinFlag := flag.Bool("k", false, "Display temperature in Kelvin")
	flag.BoolVar(kelvinFlag, "kelvin", false, "Display temperature in Kelvin")

	// Custom usage function to show only the executable name
	flag.Usage = func() {
		exeName := filepath.Base(os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [-c | --celsius] [-f | --fahrenheit] [-k | --kelvin]\n", exeName)
		fmt.Fprintf(flag.CommandLine.Output(), "  -c, --celsius      Display temperature in Celsius\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  -f, --fahrenheit   Display temperature in Fahrenheit (default)\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  -k, --kelvin       Display temperature in Kelvin\n")
	}

	flag.Parse()

	// Check if no flags were explicitly set, default to Fahrenheit
	if !*celsiusFlag && !*fahrenheitFlag && !*kelvinFlag {
		*fahrenheitFlag = true
	}

	// Ensure only one flag is set
	if (*celsiusFlag && *fahrenheitFlag) || (*celsiusFlag && *kelvinFlag) || (*fahrenheitFlag && *kelvinFlag) {
		log.Fatalf("Error: Cannot specify multiple temperature units at the same time")
	}

	// Define the PowerShell command
	psCommand := `
        Get-WmiObject MSAcpi_ThermalZoneTemperature -Namespace "root/wmi" | 
        Select-Object CurrentTemperature | 
        ForEach-Object { ($_.CurrentTemperature - 2732) / 10 }
    `

	// Execute the PowerShell command
	cmd := exec.Command("powershell", "-Command", psCommand)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing PowerShell script: %v", err)
	}

	// Trim the output to remove extra whitespace
	trimmedOutput := strings.TrimSpace(string(output))

	// Convert the temperature to the desired unit
	tempCelsius := 0.0
	fmt.Sscanf(trimmedOutput, "%f", &tempCelsius)

	var tempOutput string
	if *celsiusFlag {
		tempOutput = fmt.Sprintf("%.2f°C", tempCelsius)
	} else if *kelvinFlag {
		tempKelvin := tempCelsius + 273.15
		tempOutput = fmt.Sprintf("%.2fK", tempKelvin)
	} else {
		tempFahrenheit := (tempCelsius * 9 / 5) + 32
		tempOutput = fmt.Sprintf("%.2f°F", tempFahrenheit)
	}

	// Print the output
	fmt.Printf("%s\n", tempOutput)
}
