package fileReader

import (
	"bufio"
	"os"
	"log"
)

func openFile(filename string) *os.File {
	if len(filename) == 0 {
		return os.Stdin
	}

	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
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
		log.Print(string(line))
	}
}

func OpenAndRead(filename string) {
	file := openFile(filename)
	reader := bufio.NewReader(file)
	readFile(reader)
}
