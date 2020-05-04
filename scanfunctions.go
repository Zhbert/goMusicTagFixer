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

//scan id3 tags of each file
func deepscan(dir string, depth int) {
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
			println(depth)
			fmt.Println("├─", file.Name())

			path := dir + "/" + file.Name()
			mp3file, err := id3.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			//Artist info
			if depth != 0 {
				for y <= depth-1 {
					fmt.Print("│  :  ")
					y++
				}
				y = 0
			} else {
				print("│  │")
			}
			fmt.Println("Artist: ", mp3file.Artist())

			//Album info
			if depth != 0 {
				for y <= depth-1 {
					fmt.Print("│  :  ")
					y++
				}
				y = 0
			} else {
				print("│  │")
			}
			fmt.Println("Album:  ", mp3file.Album())

			//Title info
			if depth != 0 {
				for y <= depth-1 {
					fmt.Print("│  :  ")
					y++
				}
				y = 0
			} else {
				print("│  │")
			}
			fmt.Println("Title:  ", mp3file.Title())

			mp3file.Close()

			//find the longest string
			long := len(mp3file.Artist())
			if len(mp3file.Album()) > long {
				long = len(mp3file.Album())
			}
			if len(mp3file.Title()) > long {
				long = len(mp3file.Title())
			}

			if depth != 0 {
				for y <= depth-1 {
					fmt.Print("│  :  ")
					y++
				}
				y = 0
			} else {
				print("│  │")
			}

			y = 0
			for y < long {
				print("─")
				y++
			}
			fmt.Println(" ")
		}

		flag := file.IsDir()
		if flag == true {
			if depth != 0 {
				for y <= depth-1 {
					fmt.Print("│  ")
					y++
				}
				y = 0
			}
			fmt.Println("├─", file.Name())
			path := dir + "/" + file.Name()
			foldercount++
			deepscan(path, depth+1)
		}
	}
}
