package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {

	root := "/mnt/d/torrent"

	// var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			//is file
			name := info.Name()
			pattern := `(hhd800.com@)+`
			//(?i)((snis)|(ssni)|(dvdms)|(sdmua)|(fsdss)|(fsdss)|(fsdss)|(fsdss)|(fsdss)|(fsdss)|(fsdss)|(fsdss)|(fsdss)|(fsdss)|(fsdss)|(fsdss)){1}[-]{0,1}[\d]{1,4}
			//[\w]{3,5}[-]{0,1}[\d]{1,4}
			match, err := regexp.MatchString(pattern, name)
			if err == nil && match {
				
				renamed := strings.Replace(path, "hhd800.com@", "", 1)
				fmt.Println(path + " => " + renamed)
				err := os.Rename(path, renamed)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(files)
}
