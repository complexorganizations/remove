package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	systemPath string
	secureWipe bool
	err        error
)

func init() {
	// Check to see if any user claims have been transmitted.
	if len(os.Args) > 1 {
		secureWipeFlag := flag.Bool("secure", false, "You can secure wipe a file.")
		flag.Parse()
		secureWipe = *secureWipeFlag
		systemPath = flag.Args()[0]
	} else {
		log.Fatal("No arguments given.")
	}
}

func main() {
	if fileExists(systemPath) {
		// if we are using secure wipe
		if secureWipe {
			// open the file
			file, err := os.Open(systemPath)
			if err != nil {
				log.Fatal(err)
			}
			// Write random data to the file, same as the original file size.
			file.WriteString((randomString(fileSize(systemPath))))
			// close the file
			err = file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}
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

// Generate a random string
func randomString(bytesSize int64) string {
	randomBytes := make([]byte, bytesSize)
	rand.Read(randomBytes)
	randomString := fmt.Sprintf("%X", randomBytes)
	return randomString
}

// Get the size of a file
func fileSize(filepath string) int64 {
	file, err := os.Stat(filepath)
	if err != nil {
		log.Print(err)
	}
	return file.Size()
}
