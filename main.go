package main

import (
	"flag"
	"fmt"
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
	if len(os.Args) > 1 {
		flag.Parse()
		systemPath = flag.Args()[0]
	} else {
		log.Fatalln("Error: There are no parameters, therefore please give the path of the document(s) you want to remove.")
	}
}

func main() {
	// Remove a file
	if fileExists(systemPath) {
		secureDelete(systemPath)
	} else if folderExists(systemPath) {
		filepath.Walk(systemPath, func(path string, info os.FileInfo, err error) error {
			if fileExists(path) {
				secureDelete(path)
			}
			return err
		})
	} else {
		log.Fatalln("Error: The path you have entered does not exist.")
	}
}

// Securely wipe documents
func secureDelete(filepath string) {
	// Loop it for multiple times so its harder to recover.
	for loop := 0; loop < 3; loop++ {
		// open the file
		file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			log.Println(err)
		}
		// Write random data to the file, same as the original file size.
		_, err = file.WriteString(randomString(fileSize(filepath)))
		// Report any error if while writing to the file.
		if err != nil {
			log.Println(err)
		}
		// close the file
		err = file.Close()
		if err != nil {
			log.Println(err)
		}
	}
	// Once we have completed the loop we will remove the file.
	if fileExists(filepath) {
		err = os.Remove(filepath)
		if err != nil {
			log.Println(err)
		}
	} else if folderExists(filepath) {
		err = os.RemoveAll(filepath)
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
