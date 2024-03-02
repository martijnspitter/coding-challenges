package readfile

import (
	"flag"
	"fmt"
	"os"
)

func ReadFile(path *string) []byte {
	content, err := os.ReadFile(*path)

	if err != nil {
		fmt.Println("Something went wrong reading the file")
		flag.Usage()
	}

	return content
}
