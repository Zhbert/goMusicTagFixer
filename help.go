package main

import "fmt"

func helpView() {
	fmt.Print(`
	goMusicFixer is a simple command-line utility for to manage tags in mp3 files.
	Its main task is to edit tags in directories with one command. 
	For example, you have a catalog with a couple hundred different songs in mp3 
	format. When you fill this in a physical player or add it to the library of a 
	software player, you risk getting not just one folder with the name, say, 
	"Collection", but a couple of hundred artists and albums in the library.
	This utility allows you to solve this problem with a single command.

	Basic commands:

	scandir   -   recursively scans the directory from which the utility was 
	launched (where the terminal's "focus" was currently located) for mp3 music 
	files. The contents of directories are displayed as a tree.

	deepscan  -   its behavior is similar to the previous command, but there is 
	additional functionality - it shows three main tags under each of the tracks:
	artist, album and title.

	mp3settag -   the main command that the utility was designed for. It sets 
	the album tag value for all tracks equal to the name of the directory where 
	this tag is located.

	help      -   view this help page.

	The library is used for working with tags - mikkyang/id3-go (GitHub)
	`)
	fmt.Println(" ")
}
