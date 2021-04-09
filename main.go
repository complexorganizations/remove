package main

import (
	"os"
	"log"
)

var filePath string

func init() {
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	} else {
		log.Fatal("Error: No such file or directory")
	}
}

func main() {
	if fileExists(filePath) {
		os.Remove(filePath)
	} else if folderExists(filePath) {
		os.RemoveAll(filePath)
	} else {
		log.Fatal("Error: No such file or directory")
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
