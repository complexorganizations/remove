package main

import (
	"log"
	"os"
)

var err error

func init() {
	//
}

func main() {
	//
}

// Log errors
func handleErrors(err error) {
	if err != nil {
		log.Print(err)
	}
}

// Check if a file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Check if a folder exists
func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
