package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {

	root := "/mnt/d/torrent"

	// var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			//is file
			name := info.Name()
			pattern := `[\w]{4,5}[-]{1}[\d]{1,4}`
			match, err := regexp.MatchString(pattern, name)
			if err == nil && match {
				fmt.Println(name)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(files)
}
