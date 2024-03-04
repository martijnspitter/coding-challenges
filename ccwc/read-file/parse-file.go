package readfile

import (
	"bufio"
	"fmt"
	"os"
)

func GetContent(path string) []byte {
	if path != "" {
		return readFileFromPath(path)
	}

	return readContentFromStdIn()
}

func readFileFromPath(path string) []byte {
	content, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("Something went wrong reading the file")
		return make([]byte, 0)
	}

	return content
}

func readContentFromStdIn() []byte {
	scanner := bufio.NewScanner(os.Stdin)

	var content []byte

	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {

		content = append(content, scanner.Bytes()...)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input from console")
	}
	return content
}
