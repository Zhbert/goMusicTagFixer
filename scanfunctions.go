package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

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

		flag := file.IsDir()
		if flag == true {
			path := dir + "/" + file.Name()
			scandir(path, depth+1)
		}

	}
}
