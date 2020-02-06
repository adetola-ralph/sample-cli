package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/adetola-ralph/sample-cli/util"
)

func main() {
	// subcommands
	countCommand := flag.NewFlagSet("count", flag.ExitOnError)
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	// flag.Parse()

	// count subcommand
	countTextPtr := countCommand.String("text", "", "Text to parse (Required)")
	countMetricPtr := countCommand.String("metric", "chars", "Metric {chars|words|lines|subsring};. (Required)")
	countSubstringPtr := countCommand.String("substring", "", "The substring to be counted. Required for --metric=substring")
	countUniquePtr := countCommand.Bool("unique", false, "Measure unique value of a mteric")

	// list subcommand
	listTextPtr := listCommand.String("text", "", "Text to parse (Required)")
	listMetricPtr := listCommand.String("metric", "chars", "Metric {chars|words|lines};. (Required)")
	listUniquePtr := listCommand.Bool("unique", false, "Measure unique value of a mteric")

	// verify subcommand has been given
	// os.Arg[0] - main command
	// os.Arg[1] - sub command
	if len(os.Args) < 2 {
		fmt.Println("list or count subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		listCommand.Parse(os.Args[2:])
	case "count":
		countCommand.Parse(os.Args[2:])
	default:
		// next line seems useless because there are no defaults on the flag set
		flag.PrintDefaults()
		os.Exit(1)
	}

	if listCommand.Parsed() {
		metricChoices := [3]string{"chars", "words", "lines"}
		// required flags
		if *listTextPtr == "" {
			listCommand.PrintDefaults()
			os.Exit(1)
		}

		if exists, _ := includes(metricChoices[:], *listMetricPtr); !exists {
			listCommand.PrintDefaults()
			os.Exit(1)
		}

		listUtil := util.ListUtil{Text: *listTextPtr, Unique: *listUniquePtr}

		switch *listMetricPtr {
		case "words":
			for _, v := range listUtil.Words() {
				fmt.Println(v)
			}
		case "chars":
			for _, v := range listUtil.Chars() {
				fmt.Println(v)
			}
		case "lines":
			for _, v := range listUtil.Lines() {
				fmt.Println(v)
			}
		default:
			fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *listTextPtr, *listMetricPtr, *listUniquePtr)
		}

	}

	if countCommand.Parsed() {
		metricChoices := map[string]bool{"chars": true, "lines": true, "words": true, "substring": true}
		// required flags
		if *countTextPtr == "" {
			countCommand.PrintDefaults()
			os.Exit(1)
		}

		if _, exists := metricChoices[*countMetricPtr]; !exists {
			countCommand.PrintDefaults()
			os.Exit(1)
		}

		if *countMetricPtr == "substring" && *countSubstringPtr == "" {
			countCommand.PrintDefaults()
			os.Exit(1)
		}

		if *countMetricPtr != "substring" && *countSubstringPtr != "" {
			fmt.Println("--substring may only be used with --metric=substring.")
			countCommand.PrintDefaults()
			os.Exit(1)
		}

		countUtil := util.CountUtil{Text: *countTextPtr, Unique: *countUniquePtr, Substring: *countSubstringPtr}

		switch *countMetricPtr {
		case "words":
			fmt.Println(countUtil.Words())
		case "chars":
			fmt.Println(countUtil.Chars())
		case "lines":
			fmt.Println(countUtil.Lines())
		case "substring":
			fmt.Println(countUtil.SubstringCount())
		default:
			fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *countTextPtr, *countMetricPtr, *countUniquePtr)
		}
	}
}

func includes(array []string, value string) (bool, int) {
	for i, v := range array {
		if value == v {
			return true, i
		}
	}

	return false, -1
}
