package main

import (
    "encoding/base64"
    "encoding/json"
    "flag"
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "syscall"
    "unsafe"
)

// decodeBase64 decodes a base64 URL-encoded string with padding if necessary
func decodeBase64(data string) ([]byte, error) {
    // Add padding if necessary
    padding := len(data) % 4
    if padding > 0 {
        data += strings.Repeat("=", 4-padding)
    }
    return base64.URLEncoding.DecodeString(data)
}

// getClipboardText retrieves text from the clipboard on Windows
func getClipboardText() (string, error) {
    user32 := syscall.MustLoadDLL("user32.dll")
    kernel32 := syscall.MustLoadDLL("kernel32.dll")

    openClipboard := user32.MustFindProc("OpenClipboard")
    closeClipboard := user32.MustFindProc("CloseClipboard")
    getClipboardData := user32.MustFindProc("GetClipboardData")
    globalLock := kernel32.MustFindProc("GlobalLock")
    globalUnlock := kernel32.MustFindProc("GlobalUnlock")

    const CF_UNICODETEXT = 13

    // Open the clipboard
    r, _, err := openClipboard.Call(0)
    if r == 0 {
        return "", fmt.Errorf("failed to open clipboard: %v", err)
    }
    defer closeClipboard.Call()

    // Get clipboard data
    h, _, err := getClipboardData.Call(CF_UNICODETEXT)
    if h == 0 {
        return "", fmt.Errorf("failed to get clipboard data: %v", err)
    }

    // Lock the global memory object
    ptr, _, err := globalLock.Call(h)
    if ptr == 0 {
        return "", fmt.Errorf("failed to lock global memory: %v", err)
    }
    defer globalUnlock.Call(h)

    // Convert the clipboard data to a Go string
    text := syscall.UTF16ToString((*[1 << 20]uint16)(unsafe.Pointer(ptr))[:])

    return text, nil
}

// getJWTToken retrieves the JWT token from the positional argument or clipboard
func getJWTToken(args []string) (string, error) {
    if len(args) > 0 {
        return args[0], nil
    }
    token, err := getClipboardText()
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(token), nil
}

// decodeJWT decodes the JWT token and prints the payload as a formatted JSON
func decodeJWT(token string) error {
    // Split the token into header, payload, and signature
    parts := strings.Split(token, ".")
    if len(parts) != 3 {
        return fmt.Errorf("invalid JWT token")
    }

    // Decode the payload
    payload := parts[1]
    decodedPayload, err := decodeBase64(payload)
    if err != nil {
        return fmt.Errorf("failed to decode base64 payload: %v", err)
    }

    // Convert the decoded payload to a JSON object
    var decodedJSON map[string]interface{}
    if err := json.Unmarshal(decodedPayload, &decodedJSON); err != nil {
        return fmt.Errorf("failed to decode JWT token: %v", err)
    }

    // Print the formatted JSON
    encodedJSON, err := json.MarshalIndent(decodedJSON, "", "    ")
    if err != nil {
        return fmt.Errorf("failed to encode JSON: %v", err)
    }
    fmt.Println(string(encodedJSON))
    return nil
}

func main() {
		// Custom usage function to show only the executable name
		flag.Usage = func() {
        exeName := filepath.Base(os.Args[0])
        fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [token] [-h | -help]\n", exeName)
        fmt.Fprintf(flag.CommandLine.Output(), "  token: JWT token to decode\n")
    }

    // Parse command-line arguments
    flag.Parse()
    args := flag.Args()

    // Check if there are more than one positional argument
    if len(args) > 1 {
        fmt.Println("Error: Too many arguments. Only one positional argument is allowed.")
        os.Exit(1)
    }

    // Get the JWT token
    jwtToken, err := getJWTToken(args)
    if err != nil {
        fmt.Printf("Error retrieving JWT token: %v\n", err)
        os.Exit(1)
    }

    // Decode the JWT token
    if err := decodeJWT(jwtToken); err != nil {
        fmt.Printf("Error decoding JWT token: %v\n", err)
        os.Exit(1)
    }
}