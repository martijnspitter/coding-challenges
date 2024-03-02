package main

import (
	handleinput "cc/ccwc/handle-input"
	readfile "cc/ccwc/read-file"
	readinput "cc/ccwc/read-input"
)

func main() {
	countPath, linePath, wordPath := readinput.ReadInput()

	content := readfile.ReadFile(countPath, linePath, wordPath)

	handleinput.HandleInput(content, countPath, linePath, wordPath)

}
