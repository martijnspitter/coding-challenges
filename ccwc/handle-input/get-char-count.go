package handleinput

import (
	"fmt"
	"unicode/utf8"
)

func reportCharCount(content []byte, charPath string) {
	count := utf8.RuneCount(content)

	fmt.Printf("%v, %v", count, charPath)
}
