package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

//check name for music extension
func extensioncheck(name string) bool {
	var ok bool

	ok = false

	r, _ := regexp.Compile(`.mp3*$`)
	ok = r.MatchString(name)

	return ok
}

//function for scanning dir
func scandir(dir string, depth int) {

	//Get files and directories in path
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	//Check files and directories and print
	y := 0
	for _, file := range files {
		if depth != 0 {
			for y <= depth-1 {
				fmt.Print("│  ")
				y++
			}
			y = 0
		}

		fmt.Println("├─", file.Name())

		if extensioncheck(file.Name()) == true {
			mp3count++
		}

		flag := file.IsDir()
		if flag == true {
			path := dir + "/" + file.Name()
			foldercount++
			scandir(path, depth+1)
		}
	}
}
