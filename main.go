package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Check if a directory path is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory_path>")
		return
	}

	// Get the directory path from the command line arguments
	directory := os.Args[1]

	// Read the directory contents
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Map to keep track of extension counts
	extCount := make(map[string]int)

	// Loop through each file in the directory
	for _, file := range files {
		if !file.IsDir() { // Check if it is not a directory
			extension := filepath.Ext(file.Name())
			extCount[extension]++ // Increment the count for the extension
		}
	}

	// Print the count for each extension
	for ext, count := range extCount {
		fmt.Printf("Extension: %s, Count: %d\n", ext, count)
	}
}
