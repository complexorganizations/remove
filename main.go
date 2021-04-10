package main

import (
	"os"
	"log"
)

var filePath string

// Check for arguments
func init() {
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	} else {
		os.Exit(0)
	}
}

// Decide what type of file it is.
func main() {
	if fileExists(filePath) {
		os.Remove(filePath)
	} else if folderExists(filePath) {
		os.RemoveAll(filePath)
	} else {
		log.Fatalf("Error: %s No such file or directory.\n", filePath)
	}
}

// Directory Check
func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// File Check
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
