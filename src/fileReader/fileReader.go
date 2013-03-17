package fileReader

import (
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

type Line struct {
	key, value string
	num_value float64
}


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
		s := string(line)
		if len(s) > 0 {
			log.Print(s)
		}
	}
}

func OpenAndRead(filename string) {
	file := openFile(filename)
	reader := bufio.NewReader(file)
	readFile(reader)
}

func ExtractLines(line string) Line {
	set := strings.Split(line, " ")
	key, val := set[0], set[1]
	num_val, _ := strconv.ParseFloat(val, 32)

	l := new(Line)
	l.key = key
	l.value = val
	l.num_value = num_val

	return *l
}

func lineList(file *bufio.Reader) []Line {
	var (
		err error = nil
		line []byte
		lines []Line
	)

	for err == nil {
		line, _, err = file.ReadLine()
		if len(line) > 0 {
			l := ExtractLines(string(line))
			lines = append(lines, l)

		}
	}

	return lines

}

func CreateList(filename string) []Line {
	file := openFile(filename)
	reader := bufio.NewReader(file)
	list := lineList(reader)
	return list
}
