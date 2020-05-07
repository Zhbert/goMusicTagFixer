package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/mikkyang/id3-go"
)

//func to set mp3 tags
func mp3settag(dir string) {
	count := 0
	//Get the directory name
	r, _ := regexp.Compile(`[0-9a-zA-Z_ &]*$`)
	albumName := r.FindString(dir)

	//Get files and directories in path
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if extensioncheck(file.Name()) == true {
			path := dir + "/" + file.Name()
			mp3file, err := id3.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			mp3file.SetAlbum(albumName)
			mp3file.Close()
			count++
			//View info
			fmt.Println(file.Name(), "setiing", albumName, "album name")
		}

		flag := file.IsDir()
		if flag == true {
			path := dir + "/" + file.Name()
			mp3settag(path)
		}
	}
}
