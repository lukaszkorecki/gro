package fileReader

import "testing"

var fixture = "./test/example.txt"


func TestLineConversion(t *testing.T) {
	exp := Line{"ok", "10", 10.0}
	extracted := ExtractLines("ok 10")
	if extracted != exp {
		t.Errorf("expected %v to == %v", extracted, exp)
	}

}
