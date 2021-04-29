package main

import (
	"flag"
	"log"
	"os"
)

var (
	systemPath string
	err        error
)

func init() {
	if len(os.Args) > 1 {
		tempSystemPath := flag.String("path", "example", "Path in system")
		flag.Parse()
		systemPath = *tempSystemPath
	} else {
		log.Fatal("Error: The system path has not been given.")
	}
	// System path
	if systemPath == "" {
		log.Fatal("Error: The system path has not been given.")
	}
}

func main() {
	if fileExists(systemPath) {
		deleteFile()
	}
	if folderExists(systemPath) {
		deleteFolder()
	}
}

func deleteFile() {
	err = os.Remove(systemPath)
	handleErrors(err)
}

func deleteFolder() {
	err = os.RemoveAll(systemPath)
	handleErrors(err)
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
