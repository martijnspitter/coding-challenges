package readinput

import (
	"flag"
	"fmt"
)

func ReadInput() (*string, *string) {
	countPath := flag.String("c", "", "File to parse")
	linePath := flag.String("l", "", "File to parse")

	flag.Parse()

	if *countPath == "" && *linePath == "" {
		fmt.Println("Please provide a file path.")
		flag.Usage()
	}

	return countPath, linePath
}
