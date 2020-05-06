package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/mikkyang/id3-go"
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

		if extensioncheck(file.Name()) == true {
			mp3count++

			if depth != 0 {
				for y <= depth-1 {
					fmt.Print("│  ")
					y++
				}
				y = 0
			}
			fmt.Println("├─", file.Name())
		}

		flag := file.IsDir()
		if flag == true {
			if depth != 0 {
				for y <= depth-1 {
					fmt.Print("│  ")
					y++
				}
			}
			fmt.Println("├─", file.Name())
			path := dir + "/" + file.Name()
			foldercount++
			scandir(path, depth+1)
		}
	}
}

//function for scanning tags
func deepscantwo(dir string, depth int) {

	//Get files and directories in path
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	//Check files and directories and print
	y := 0
	for _, file := range files {

		if extensioncheck(file.Name()) == true {
			mp3count++

			if depth != 0 {
				for y <= depth-1 {
					fmt.Print("│  ")
					y++
				}
				y = 0
			}
			fmt.Println("├─", file.Name())

			//open file
			path := dir + "/" + file.Name()
			mp3file, err := id3.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			i := 0
			y = 0
			for i < 3 {
				if depth != 0 {
					for y <= depth-1 {
						fmt.Print("│  ")
						y++
					}
					y = 0
				} else {
					fmt.Print("   ")
				}
				fmt.Print("   ")
				switch i {
				case 0:
					fmt.Println("Artist: ", mp3file.Artist())
				case 1:
					fmt.Println("Album:  ", mp3file.Album())
				case 2:
					fmt.Println("Title:  ", mp3file.Title())
				}
				i++
				mp3file.Close()
			}
		}

		flag := file.IsDir()
		if flag == true {
			if depth != 0 {
				for y <= depth-1 {
					fmt.Print("│  ")
					y++
				}
			}
			fmt.Println("├─", file.Name())
			path := dir + "/" + file.Name()
			foldercount++
			deepscantwo(path, depth+1)
		}
	}
}
