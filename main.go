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
	// System path
	if len(systemPath) < 1 || systemPath == "/user/example/folder/file" {
		log.Fatal("Error: The system path has not been given.")
	}
}

func main() {
	if fileExists(systemPath) {
		err = os.Remove(systemPath)
		if err != nil {
			log.Println(err)
		}
	}
	if folderExists(systemPath) {
		err = os.RemoveAll(systemPath)
		if err != nil {
			log.Println(err)
		}
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
