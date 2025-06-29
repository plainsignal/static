package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Define command line flags
	dir := flag.String("dir", ".", "Directory to serve static files from")
	port := flag.String("port", "8080", "Port to listen on")
	help := flag.Bool("help", false, "Display usage information")
	flag.Parse()

	// Display help and exit if -help flag is provided
	if *help {
		fmt.Println("Static File Server")
		fmt.Println("A simple Go program to serve static files from a specified directory over HTTP.")
		fmt.Println("\nUsage:")
		fmt.Println("  go run main.go [flags]\n")
		fmt.Println("  # If installed via go install github.com/plainsignal/static\n")
		fmt.Println("  static [flags]\n")
		fmt.Println("Flags:")
		flag.PrintDefaults()
		fmt.Println("\nExample:")
		fmt.Println("  go run main.go -dir=/path/to/your/directory -port=8080\n")
		fmt.Println("  # If installed via go install github.com/plainsignal/static\n")
		fmt.Println("  static -dir=/path/to/your/directory -port=8080")
		os.Exit(0)
	}

	// Check if directory exists
	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		log.Fatalf("Directory %s does not exist", *dir)
	}

	// Create file server handler
	fs := http.FileServer(http.Dir(*dir))

	// Set up handler with logging
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		fs.ServeHTTP(w, r)
	})

	// Start server
	addr := fmt.Sprintf(":%s", *port)
	fmt.Printf("Serving files from %s on http://localhost%s\n", *dir, addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
