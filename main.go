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
	// Check to see if any user claims have been transmitted.
	if len(os.Args) > 1 {
		tempSystemPath := flag.String("path", "/user/example/folder/file", "The location of the file(s) to be deleted.")
		flag.Parse()
		systemPath = *tempSystemPath
	} else {
		log.Fatal("Error: The system path has not been given.")
	}
	// System path
	if len(systemPath) < 1 || systemPath == "/user/example/folder/file" {
		log.Fatal("Error: The system path has not been given.")
	}
}

func main() {
	if fileExists(systemPath) {
		deleteFile()
	} else if folderExists(systemPath) {
		deleteFolder()
	} else {
		log.Fatal("Error: The machine direction is invalid.")
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
		log.Println(err)
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
