package fileReader

import (
	"testing"
	"os"
)

var fixture = "./test/example.txt"


func TestExtractingDataFromALine(t *testing.T) {
	exp := Line{"ok", "10", 10.0}
	extracted := extractLineData("ok 10")
	if extracted != exp {
		t.Errorf("expected %v to == %v", extracted, exp)
	}

}

func TestFileReading(t *testing.T) {
	// slightly dirty, TODO how to mock out
	// a Reader?
	pwd, _ := os.Getwd()
	var (
		fname string = pwd+"/../../fixtures/example.txt"
		exp []Line
		acc int = 0
	)

	exp = append(exp, Line{"foo", "20", 20.0})
	exp = append(exp, Line{"yo", "40", 40.0})
	exp = append(exp, Line{"baz", "10", 10.0})
	exp = append(exp, Line{"qux", "20", 20.0})

	data := CreateListFromFile(fname)
	for i, item := range  exp {
		if data[i] == item {
			acc++
		}
	}
	if acc != 4 {
		t.Errorf("exptected %v to == %v", exp, data)
	}

}
