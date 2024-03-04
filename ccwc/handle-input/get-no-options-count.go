package handleinput

import (
	"bytes"
	"fmt"
	"strings"
)

func reportNoOptionsCount(content []byte, path string) {
	lineSep := []byte{'\n'}
	lineCount := bytes.Count(content, lineSep)

	words := strings.Fields(string(content))
	wordCount := len(words)

	byteCount := len(content)

	fmt.Printf("%v %v %v %v", lineCount, wordCount, byteCount, path)
}
