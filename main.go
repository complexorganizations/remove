package main

import (
	"flag"
	"log"
	"os"
)

var (
	systemPath string
	secureWipe bool
	err        error
)

func init() {
	secureWipeFlag := flag.Bool("secure", false, "You can secure wipe a file.")
	flag.Parse()
	secureWipe = *secureWipeFlag
	systemPath = flag.Args()[0]
}

func main() {
	if fileExists(systemPath) {
		if secureWipe {
			file, err := os.Open(systemPath)
			if err != nil {
				log.Fatal(err)
			}
			fileSize := len(file.Name())
			for loop := 0; loop < fileSize; loop++ {
				file.WriteString("0")
			}
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
