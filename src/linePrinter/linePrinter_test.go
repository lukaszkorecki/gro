package linePrinter

import  (
	"testing"
	"fileReader"
)

func TestGeneratingGraphData(t *testing.T) {

	expectedData := []float64 {20.0, 40, 10, 20}
	width := float64(40)
	lines := []fileReader.Line {
		fileReader.Line{"foo", "2", 2.0},
		fileReader.Line{"yo" , "4", 4.0},
		fileReader.Line{"baz", "1", 1.0},
		fileReader.Line{"qux", "2", 2.0},
	}
	c := 0

	data := graphData(lines, width)
	for i, item := range data {
		if expectedData[i] != item { c++ }
	}

	if  c != 0 {
		t.Errorf("expected %v to == %v" , expectedData, data)
	}
}
