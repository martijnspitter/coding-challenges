package handleinput

import (
	"fmt"
	"strings"
)

func reportWordCount(content []byte, wordPath *string) {
	words := strings.Fields(string(content))
	count := len(words)

	fmt.Printf("%v, %v", count, *wordPath)
}
