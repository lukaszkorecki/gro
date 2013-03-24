package linePrinter

import (
	"log"
	"fmt"
	"fileReader"
	"sort"
)

func Simple(lines []fileReader.Line) {
	for _, l := range lines {
		log.Print(l.Key)
	}
}

func BarGraph(lines []fileReader.Line, screenWidth int) {
	graph := graphData(lines, float64(screenWidth))
	for i, item := range graph {
		s := ""
		for i := 0 ; i <= int(item) ; i++ { s += "=" }
		fmt.Printf("%s\n", s)
		fmt.Printf("%s %s\n", lines[i].Key, lines[i].Value)
	}
}

func graphData(lines []fileReader.Line, screenWidth float64) []float64{

	var (
		res, sorted, graphData []float64
		max float64
	)
	/*
	screenwith = 80
	data = 10, 20, 40, 10
	max = 40

	percents f
	*/
	for _, item := range lines {
		res = append(res, item.NumValue)
		sorted = append(sorted, item.NumValue)
	}

	sort.Float64s(sorted)
	max = sorted[len(sorted) - 1]

	for _, item := range res {
		percent := ratioOf(max, item)
		barWidth := (screenWidth * percent)
		graphData = append(graphData, barWidth)

	}
	return graphData

}

func ratioOf(max, item float64)  float64 {
	return (item  / max)
}
