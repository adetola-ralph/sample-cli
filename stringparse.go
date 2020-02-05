package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	metricChoices := [3]string{"chars", "words", "lines"}

	textPtr := flag.String("text", "", "Text to parse (Required)")
	metricPtr := flag.String("metric", "char", "Metric {chars|words:lines};. (Required)")
	uniquePtr := flag.Bool("unique", false, "Measure unique value of a mteric")
	flag.Parse()

	if *textPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if exists, _ := includes(metricChoices[:], *metricPtr); !exists {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println(textPtr)

	fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *textPtr, *metricPtr, *uniquePtr)
}

func includes(array []string, value string) (bool, int) {
	for i, v := range array {
		if value == v {
			return true, i
		}
	}

	return false, -1
}
