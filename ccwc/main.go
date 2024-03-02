package main

import (
	readfile "cc/ccwc/read-file"
	readinput "cc/ccwc/read-input"
	"fmt"
)

func main() {
	path := readinput.ReadInput()

	content := readfile.ReadFile(path)

	contentLength := len(content)

	fmt.Printf("%v %v", contentLength, *path)

}
