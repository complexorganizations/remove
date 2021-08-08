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
		flag.Parse()
		systemPath = flag.Args()[0]
	} else {
		log.Fatalln("Error: There are no parameters, therefore please give the path of the document(s) you want to remove.")
	}
}

func main() {
	documentType(systemPath)
}

func documentType(filePathInSystem string) {
	if fileExists(filePathInSystem) {
		err = os.Remove(filePathInSystem)
		if err != nil {
			log.Println(err)
		}
	} else if folderExists(filePathInSystem) {
		err = os.RemoveAll(filePathInSystem)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Fatalln("Error: The path you have entered does not exist.")
	}
}

// Check to see whether a file already exists.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// Check to see if a folder already exists.
func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if err != nil {
		return false
	}
	return info.IsDir()
}
