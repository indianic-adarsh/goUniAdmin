package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Get the home directory from the environment
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		fmt.Println("Error: HOME environment variable is not set")
		return
	}

	// Construct the full path to the swag binary
	swagBinary := filepath.Join(homeDir, "go", "bin", "swag")

	// Define the base directory to search from
	baseDir := "./"

	// List to hold the directories containing Go files
	var dirs []string

	// Walk through the base directory to find Go files
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// If the file is a Go file, add the directory to dirs
		if strings.HasSuffix(info.Name(), ".go") {
			dir := filepath.Dir(path)
			dirs = append(dirs, dir)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error while walking the directory:", err)
		return
	}

	// Remove duplicates from dirs (in case multiple Go files are in the same directory)
	uniqueDirs := unique(dirs)

	// Join directories with commas
	dirStr := strings.Join(uniqueDirs, ",")

	// Run swag init with the dynamic directories
	cmd := exec.Command(swagBinary, "init", "--dir", dirStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to generate swagger docs:", err)
	}
}

// Utility to remove duplicates from a slice of strings
func unique(strs []string) []string {
	uniqueMap := make(map[string]struct{})
	var uniqueStrs []string
	for _, str := range strs {
		if _, ok := uniqueMap[str]; !ok {
			uniqueMap[str] = struct{}{}
			uniqueStrs = append(uniqueStrs, str)
		}
	}
	return uniqueStrs
}
