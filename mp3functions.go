package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/mikkyang/id3-go"
)

//func to set mp3 tags
func mp3settag(dir string, depth int) {
	//Get the directory name
	r, _ := regexp.Compile(`[0-9a-zA-Z0-9А-Яа-я_ &]*$`)
	albumName := r.FindString(dir)
	println(albumName)

	//Get files and directories in path
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		flag := file.IsDir()
		if flag == true {
			path := dir + "/" + file.Name()
			mp3settag(path, depth+1)
		}

		if extensioncheck(file.Name()) == true {
			path := dir + "/" + file.Name()
			mp3file, err := id3.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			mp3file.SetAlbum(albumName)
			mp3file.SetArtist(albumName)
			defer mp3file.Close()
			//View info
			fmt.Println("\x1b[35m", file.Name(), "\x1b[0m", "setting", "\x1b[32m", albumName, "\x1b[0m", "album and artist name")
		}
	}
}
