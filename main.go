package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// API represents the structure of an API definition.
type API struct {
	Path       string          `json:"path"`
	Method     string          `json:"method"`
	HttpStatus int             `json:"http_status"`
	Body       json.RawMessage `json:"body"`
	Latency    int             `json:"latency"`
}

// APIs is a slice of API definitions.
type APIs []API

// pingHandler responds with a simple "pong" message.
func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func test2(){

}

// apiHandler processes incoming requests based on the API definition.
func apiHandler(w http.ResponseWriter, r *http.Request, api API) {
	time.Sleep(time.Duration(api.Latency) * time.Millisecond)

	w.WriteHeader(api.HttpStatus)
	w.Write(api.Body)
}

// readAPIs reads the API definitions from a JSON file.
func readAPIs(filePath string) (APIs, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file) // Updated to use io.ReadAll
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var apis APIs
	if err := json.Unmarshal(data, &apis); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return apis, nil
}

// setupRoutes initializes the HTTP routes based on the provided APIs.
func setupRoutes(apis APIs) {
	for _, api := range apis {
		api := api // capture range variable
		http.HandleFunc(api.Path, func(w http.ResponseWriter, r *http.Request) {
			if r.Method == api.Method {
				apiHandler(w, r, api)
			} else {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		})
		// Print the route information
		log.Printf("Route registered: %s %s", api.Method, api.Path)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the path to the API definitions JSON file.")
	}

	filePath := os.Args[1]
	apis, err := readAPIs(filePath)
	if err != nil {
		log.Fatalf("Error reading APIs: %v", err)
	}

	setupRoutes(apis)
	http.HandleFunc("/ping", pingHandler)

	port := 8080
	if portEnv, exists := os.LookupEnv("PORT"); exists {
		if p, err := strconv.Atoi(portEnv); err == nil {
			port = p
		}
	}

	fmt.Printf("Server is running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
