package handleinput

import "fmt"

func reportByteCount(content []byte, path *string) {
	byteCount := len(content)

	fmt.Printf("%v %v", byteCount, *path)

}
