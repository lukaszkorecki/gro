package fileReader

import (
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

type Line struct {
	Key, Value string
	NumValue float64
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

func extractLineData(line string) Line {
	set := strings.Split(line, " ")
	key, val := set[0], set[1]
	num_val, _ := strconv.ParseFloat(val, 32)

	l := new(Line)
	l.Key = key
	l.Value = val
	l.NumValue = float64(num_val)

	return *l
}

func lineList(file *bufio.Reader) []string {
	var (
		err error = nil
		line []byte
		lines []string
	)

	for err == nil {
		line, _, err = file.ReadLine()
		if len(line) > 0 {
			lines = append(lines, string(line))
		}
	}

	return lines
}

func convertLines(lines []string) []Line {
	var converted []Line
	for _, line := range lines {
		l := extractLineData(string(line))
		converted = append(converted, l)
	}
	return converted
}

func CreateListFromFile(filename string) []Line {
	file := openFile(filename)
	reader := bufio.NewReader(file)
	list := lineList(reader)
	converted := convertLines(list)
	return converted
}
