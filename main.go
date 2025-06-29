package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Define command line flag for directory
	dir := flag.String("dir", ".", "Directory to serve static files from")
	port := flag.String("port", "8080", "Port to listen on")
	flag.Parse()

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
