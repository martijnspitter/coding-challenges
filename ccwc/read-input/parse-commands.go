package readinput

import (
	"flag"
	"fmt"
)

func ReadInput() (*string, *string, *string) {
	countPath := flag.String("c", "", "File to parse")
	linePath := flag.String("l", "", "File to parse")
	wordPath := flag.String("w", "", "File to parse")

	flag.Parse()

	if *countPath == "" && *linePath == "" && *wordPath == "" {
		fmt.Println("Please provide a file path.")
		flag.Usage()
	}

	return countPath, linePath, wordPath
}
