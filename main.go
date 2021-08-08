package main

import (
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

var (
	systemPath string
	err        error
)

func init() {
	// Check to see if any user claims have been transmitted.
	if len(os.Args) < 1 {
		log.Fatal("Error: No flags provided. Please use -help for more information.")
	} else {
		systemPath = os.Args[1]
	}
}

func main() {
	// Remove a file
	if fileExists(systemPath) {
		// Replace the content of the file multiple times and than remove it.
		for loop := 0; loop < 3; loop++ {
			secureDelete(systemPath)
		}
		err = os.Remove(systemPath)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Remove the folder
	if folderExists(systemPath) {
		// Remove the content of the file multiple times and than delete it.
		for loop := 0; loop < 3; loop++ {
			err = filepath.Walk(systemPath, func(path string, info fs.FileInfo, err error) error {
				secureDelete(path)
				return nil
			})
			if err != nil {
				log.Println(err)
			}
		}
		err = os.RemoveAll(systemPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Securely wipe documents
func secureDelete(filepath string) {
	// open the file
	file, err := os.OpenFile(systemPath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
	}
	// Write random data to the file, same as the original file size.
	_, err = file.WriteString(randomString(fileSize(systemPath)))
	// Report any error if while writing to the file.
	if err != nil {
		log.Fatalln(err)
	}
	// close the file
	err = file.Close()
	if err != nil {
		log.Fatalln(err)
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

// Generate a random string
func randomString(bytesSize int64) string {
	rand.Seed(time.Now().UTC().UnixNano())
	randomBytes := make([]byte, bytesSize)
	rand.Read(randomBytes)
	randomString := fmt.Sprintf("%X", randomBytes)
	return randomString
}

// Get the size of a file
func fileSize(filepath string) int64 {
	file, err := os.Stat(filepath)
	if err != nil {
		log.Println(err)
	}
	return file.Size()
}
