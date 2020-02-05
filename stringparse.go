package main

import (
	"flag"
	"fmt"
)

func main() {
	textPtr := flag.String("text", "", "Text to parse")
	metricPtr := flag.String("metric", "char", "Metric {chars|words:lines};.")
	uniquePtr := flag.Bool("unique", false, "Measure unique value of a mteric")
	flag.Parse()

	fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *textPtr, *metricPtr, *uniquePtr)
}
