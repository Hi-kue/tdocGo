package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Helper function to print out the path of the file or directory
func helperPath(info os.FileInfo, path string) {
	if info.IsDir() {
		fmt.Println("Directory Found: ", path)
	} else {
		fmt.Println("File Found: ", path)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	time := time.NewTicker(4 * time.Second)
	cwd, err := os.Getwd()
	checkErr(err)

	for {
		fileCount, dirCount := 0, 0

		filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				helperPath(info, path)
				dirCount++
			} else {
				helperPath(info, path)
				fileCount++
			}
			return nil
		})

		fmt.Printf("There are %d files and %d directories in %s\n", fileCount, dirCount, cwd)
		<-time.C
		fmt.Println("Loading...")
		<-time.C
		fmt.Println("There is now a malicious crawler running in the background, thanks for testing the program :D.")
		<-time.C
		os.Exit(0)
	}
}
