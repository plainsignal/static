# Static File Server

A simple command-line Go program that serves static files from a specified directory over HTTP.

## Features

- Serves static files from any specified directory
- Configurable port via command-line flag
- Logs incoming requests with client IP, method, and URL
- Error handling for invalid directories
- Default settings for quick setup

## Prerequisites

Go 1.16 or higher

## Installation

Save the program as main.go
Ensure you have Go installed
No additional dependencies required

## Usage

Run the server with default settings (current directory, port 8080):
go run main.go

Specify a custom directory and port:
go run main.go -dir=/path/to/your/directory -port=8080

Access the files through a web browser at:
http://localhost:8080

### Command-line Flags

-dir: Directory to serve files from (default: current directory .)
-port: Port to listen on (default: 8080)

### Example

To serve files from /var/www/html on port 3000:
go run main.go -dir=/var/www/html -port=3000

## Error Handling

The program checks if the specified directory exists
Logs errors if the server fails to start
Provides clear error messages for invalid inputs

## Logging

Each HTTP request is logged with:

- Client IP address
- HTTP method
- Requested URL

### Notes

Ensure the specified directory has appropriate read permissions
The server runs indefinitely until manually stopped (Ctrl+C)
Uses Go's standard library net/http package

## License

MIT License

Copyright (c) 2025

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
