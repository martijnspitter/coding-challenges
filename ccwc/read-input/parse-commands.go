package readinput

import (
	"flag"
	"fmt"
)

func ReadInput() *string {
	path := flag.String("c", "", "File to parse")

	flag.Parse()

	if *path == "" {
		fmt.Println("Please provide a file path.")
		flag.Usage()
	}

	return path
}
