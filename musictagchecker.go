package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"
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

			dir, err := os.Getwd() //current directory
			if err != nil {
				log.Fatal(err)
			}
			//View main directory name
			r, _ := regexp.Compile(`\w*$`)
			fmt.Println(r.FindString(dir))

			scandir(dir, 0) //start scanning with 0 depth
		case "exit": //Exit program
			os.Exit(0)
		default: //Incorrect command
			fmt.Println("You have not entered any command!")
		}
	}
}
