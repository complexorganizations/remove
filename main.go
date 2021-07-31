package main

import (
	"log"
	"os"
)

var (
	systemPath string
	err        error
)

func init() {
	// Check to see if any user claims have been transmitted.
	if len(os.Args) < 1 {
		log.Fatal("Error: The system path has not been given.")
	} else {
		systemPath = os.Args[1]
	}
}

func main() {
	if fileExists(systemPath) {
		err = os.Remove(systemPath)
		if err != nil {
			log.Println(err)
		}
	} else if folderExists(systemPath) {
		err = os.RemoveAll(systemPath)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Fatal("Error: The document could not be found on your local system.")
	}
}

// Check if a file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// Check if a folder exists
func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if err != nil {
		return false
	}
	return info.IsDir()
}
