package main

import (
	"log"
	"math/rand"
	"os"
	"io/fs"
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
		log.Fatal("Error: The system path has not been given.")
	} else {
		systemPath = os.Args[1]
	}
}

func main() {
	// Remove a file
	if fileExists(systemPath) {
		secureDelete(systemPath)
	}
	// Remove the folder
	if folderExists(systemPath) {
		err = filepath.Walk(systemPath, func(path string, info fs.FileInfo, err error) error {
			secureDelete(path)
			return nil
		})
		if err != nil {
			log.Println(err)
		}
	}
}

// Securely wipe documents
func secureDelete(filepath string) {
	// open the file
	file, err := os.Open(systemPath)
	if err != nil {
		err = os.Remove(filepath)
		if err != nil {
			log.Print("Coudent open the file so tried to delete it but failed.")
		}
		log.Print(err)
	}
	// Write random data to the file, same as the original file size.
	randomData := randomString(fileSize(systemPath))
	file.WriteString(string(randomData))
	// close the file
	err = file.Close()
	if err != nil {
		log.Fatal(err)
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
func randomString(bytesSize int64) []byte {
	rand.Seed(time.Now().UTC().UnixNano())
	randomByte := make([]byte, bytesSize)
	rand.Read(randomByte)
	return randomByte
}

// Get the size of a file
func fileSize(filepath string) int64 {
	file, err := os.Stat(filepath)
	if err != nil {
		log.Print(err)
	}
	return file.Size()
}
