package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"
)

func scanSecondDir(dir string, level int) {
	//View main directory name
	r, _ := regexp.Compile(`\w*$`)
	fmt.Println(r.FindString(dir))

	//Get files and directories in path
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	//Check files and directories and print
	i := 1
	for _, file := range files {
		if i < len(files) {
			fmt.Println("├─", file.Name(), file.IsDir())
			i++
		} else {
			fmt.Println("└─", file.Name(), file.IsDir())
		}
	}
}

//function for scanning dir
func scandir() {
	dir, err := os.Getwd() //current directory
	if err != nil {
		log.Fatal(err)
	}

	//View main directory name
	r, _ := regexp.Compile(`\w*$`)
	fmt.Println(r.FindString(dir))

	//Get files and directories in path
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	//Check files and directories and print
	i := 1
	for _, file := range files {
		if i < len(files) {
			fmt.Println("├─", file.Name(), file.IsDir())
			flag := file.IsDir()
			if flag == true {
				path := dir + "/" + file.Name()
				scanSecondDir(path, 1)
			}
			i++
		} else {
			fmt.Println("└─", file.Name(), file.IsDir())
		}
	}
}

func main() {
	var command string

	//Greetings
	dt := time.Now()
	fmt.Println("Starting program at", dt.Format("15:04:05 01-02-2006"))

	//Main cycle
	for {
		fmt.Println("Please enter the command:")
		fmt.Scanf("%s", &command)

		//check the command
		switch command {
		case "scandir": //Scanning current directory
			dt := time.Now()
			fmt.Println("Starting the folder scan at", dt.Format("15:04:05 01-02-2006"))
			scandir()
		case "exit": //Exit program
			os.Exit(0)
		default: //Incorrect command
			fmt.Println("You have not entered any command!")
		}
	}
}
