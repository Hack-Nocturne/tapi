package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
)

func main() {
	// Print all environment variables in a prettified way
	fmt.Println("=== Environment Variables ===")
	printEnvVariables()
	fmt.Println("\n=== Starting HTTP Server ===")

	// Set up HTTP routes
	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/env", envHandler)

	// Start the server
	port := getPort()
	fmt.Printf("Server starting on port %s\n", port)
	fmt.Printf("Hello World: http://localhost:%s/\n", port)
	fmt.Printf("Environment Variables: http://localhost:%s/env\n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// helloWorldHandler handles the hello world route
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]any{
		"message": "Hey I am v2 🌿",
		"status":  "success",
		"path":    r.URL.Path,
		"method":  r.Method,
	}

	json.NewEncoder(w).Encode(response)
}

// envHandler returns all environment variables as JSON
func envHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	envMap := make(map[string]string)
	for _, env := range os.Environ() {
		// Split environment variable into key and value
		for i := 0; i < len(env); i++ {
			if env[i] == '=' {
				key := env[:i]
				value := env[i+1:]
				envMap[key] = value
				break
			}
		}
	}

	response := map[string]any{
		"environment_variables": envMap,
		"count":                 len(envMap),
	}

	// Pretty print JSON
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(response)
}

// printEnvVariables prints all environment variables in a prettified way to console
func printEnvVariables() {
	envVars := os.Environ()

	// Create a map to store environment variables
	envMap := make(map[string]string)
	for _, env := range envVars {
		// Split environment variable into key and value
		for i := 0; i < len(env); i++ {
			if env[i] == '=' {
				key := env[:i]
				value := env[i+1:]
				envMap[key] = value
				break
			}
		}
	}

	// Sort keys for consistent output
	keys := make([]string, 0, len(envMap))
	for key := range envMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Print in a prettified format
	fmt.Printf("Found %d environment variables:\n\n", len(envMap))
	for _, key := range keys {
		value := envMap[key]
		// Truncate very long values for readability
		if len(value) > 100 {
			value = value[:97] + "..."
		}
		fmt.Printf("%-30s = %s\n", key, value)
	}
}

// getPort returns the port to use, defaulting to 8080 if PORT env var is not set
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
