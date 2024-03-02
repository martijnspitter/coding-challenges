package readfile

import (
	"flag"
	"fmt"
	"os"
)

func ReadFile(countPath *string, linePath *string, wordPath *string) []byte {
	path := ""

	if *countPath != "" {
		path = *countPath
	}

	if *linePath != "" {
		path = *linePath
	}

	if *wordPath != "" {
		path = *wordPath
	}

	content, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("Something went wrong reading the file")
		flag.Usage()
	}

	return content
}
