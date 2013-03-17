package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// FIXME a bit more info perhaps?
var help_flag = flag.Bool("help", false, "help")

var help_str = "gro [FILENAME] or $COMMAND | gro"

func log(s string) {
	fmt.Println(s)
}

func logerr(err error) {
	fmt.Fprintf(os.Stderr, "ERROR %s", err)
}

func logerrAndExit(err error) {
	logerr(err)
	os.Exit(1)
}

func openFile(filename string) *os.File {
	if len(filename) == 0 {
		return os.Stdin
	}

	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		logerrAndExit(err)
	}

	// trololo

	return file
}

func readFile(file *bufio.Reader) {
	var (
		err  error = nil
		line []byte
	)

	for err == nil {
		line, _, err = file.ReadLine()
		log(string(line))
	}
}

func main() {
	flag.Parse()
	if *help_flag {
		log(help_str)
		os.Exit(0)
	}

	filename := ""
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	file := openFile(filename)
	reader := bufio.NewReader(file)
	readFile(reader)

}
