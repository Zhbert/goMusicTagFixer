package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

var mp3count int
var foldercount int

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

			mp3count = 0
			foldercount = 0
			scandir(dir, 0) //start scanning with 0 depth

			//View all statistic
			y := 0
			for y < 40 {
				print("─")
				y++
			}
			fmt.Println("─")
			fmt.Println("Scanning complete at ", dt.Format("15:04:05 01-02-2006"))
			fmt.Println("Founded ", mp3count, " mp3 files in ", foldercount, " folders.")
			fmt.Println(" ")
		case "exit": //Exit program
			os.Exit(0)
		case "deepscan":
			dt := time.Now()
			fmt.Println("Deep scan started at ", dt.Format("15:04:05 01-02-2006"))

			dir, err := os.Getwd() //current directory
			if err != nil {
				log.Fatal(err)
			}

			//View main directory name
			r, _ := regexp.Compile(`\w*$`)
			fmt.Println(r.FindString(dir))

			mp3count = 0
			foldercount = 0
			deepscantwo(dir, 0)

			//View all statistic
			y := 0
			for y < 40 {
				print("─")
				y++
			}
			fmt.Println("─")
			fmt.Println("Scanning complete at ", dt.Format("15:04:05 01-02-2006"))
			fmt.Println("Founded ", mp3count, " mp3 files in ", foldercount, " folders.")
			fmt.Println(" ")
		default: //Incorrect command
			fmt.Println("You have not entered any command!")
		}
	}
}
