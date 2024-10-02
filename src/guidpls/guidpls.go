package main

import (
    "flag"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strconv"
    "strings"

    "github.com/google/uuid"
)

func main() {
    // Define the nocopy flag with two names
    nocopy := flag.Bool("nocopy", false, "Do not copy GUIDs to clipboard")
    nocopyShort := flag.Bool("n", false, "Do not copy GUIDs to clipboard (short)")

    // Custom usage function to show only the executable name
    flag.Usage = func() {
        exeName := filepath.Base(os.Args[0])
        fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [count] [-nocopy | -n]\n", exeName)
        fmt.Fprintf(flag.CommandLine.Output(), "  count: Number of GUIDs to generate (default is 1)\n")
        flag.PrintDefaults()
    }

    flag.Parse()

    // Check if the count argument is provided
    var count int
    var err error
    if flag.NArg() > 0 {
        count, err = strconv.Atoi(flag.Arg(0))
        if err != nil || count < 1 {
            fmt.Println("Invalid count. Please provide a positive integer.")
            return
        }
    } else {
        count = 1 // Default count is 1
    }

    // Generate the specified number of GUIDs
    guids := make([]string, count)
    for i := 0; i < count; i++ {
        guids[i] = uuid.New().String()
    }

    // Concatenate the GUIDs into a comma-separated list
    guidsStr := strings.Join(guids, ",")

    // Copy the GUIDs to the clipboard using the clip command if neither -nocopy nor -n is specified
    if !*nocopy && !*nocopyShort {
        cmd := exec.Command("cmd", "/c", "echo|set /p="+guidsStr+"| clip")
        cmd.Run()
    }

    // Print each GUID on a separate line
    for _, guid := range guids {
        fmt.Println(guid)
    }
}