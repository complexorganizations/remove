package main

import (
	"log"
	"os"
	"path/filepath"
)

var (
	systemPath string
	err        error
)

// Check for arguments
func init() {
	if len(os.Args) > 1 {
		systemPath = os.Args[1]
	} else {
		log.Fatal("Error: no arguments passed.")
	}
}

// Decide what type of file it is.
func main() {
	decideType()
}

func decideType() {
	// Simply delete the file if it is a file.
	if fileExists(systemPath) {
		err = os.Remove(systemPath)
		handlePrintlnErrors(err)
	}
	// If it's a folder, go through all of the files and delete them, then go through the folder and delete it.
	if folderExists(systemPath) {
		err = filepath.Walk(systemPath, func(path string, info os.FileInfo, err error) error {
			// Handle error for pathwalk return
			handlePrintlnErrors(err)
			// Remove the file
			if fileExists(path) {
				err = os.Remove(path)
				handlePrintlnErrors(err)
			}
			return nil
		})
		// Handle error for the walk function
		handlePrintlnErrors(err)
		// Remove the directory from your system
		if folderExists(systemPath) {
			err = os.RemoveAll(systemPath)
			handlePrintlnErrors(err)
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

// Error Handle
func handlePrintlnErrors(err error) {
	if err != nil {
		log.Println(err)
	}
}
