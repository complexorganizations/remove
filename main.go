package main

import (
	"log"
	"os"
)

var (
	filePath string
	err      error
)

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
		err = os.Remove(filePath)
		if err != nil {
			log.Println(err)
		}
	} else if folderExists(filePath) {
		err = os.RemoveAll(filePath)
		if err != nil {
			log.Println(err)
		}
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
