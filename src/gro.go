package main

import (
	"flag"
	"os"
	"log"
	"fileReader"
	"linePrinter"
)

// FIXME a bit more info perhaps?
var help_flag = flag.Bool("help", false, "help")

var help_str = "gro [FILENAME] or $COMMAND | gro"

func main() {
	flag.Parse()
	if *help_flag {
		log.Print(help_str)
		os.Exit(0)
	}

	filename := ""
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}

	lines := fileReader.CreateListFromFile(filename)
	linePrinter.Simple(lines)
}
