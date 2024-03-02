package readfile

import (
	"flag"
	"fmt"
	"os"
)

func ReadFile(countPath *string, linePath *string) []byte {
	path := ""

	if *countPath != "" {
		path = *countPath
	}

	if *linePath != "" {
		path = *linePath
	}

	content, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("Something went wrong reading the file")
		flag.Usage()
	}

	return content
}
