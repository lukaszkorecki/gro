package linePrinter

import (
	"log"
	"fileReader"
)

func Simple(lines []fileReader.Line) {
	for _, l := range lines {
		log.Print(l.Key)
	}
}
