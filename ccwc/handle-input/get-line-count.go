package handleinput

import (
	"bytes"
	"fmt"
)

func reportLineCount(content []byte, linePath string) {
	lineSep := []byte{'\n'}
	lineCount := bytes.Count(content, lineSep)

	fmt.Printf("%v %v", lineCount, linePath)
}
