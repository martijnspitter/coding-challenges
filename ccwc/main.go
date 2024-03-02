package main

import (
	handleinput "cc/ccwc/handle-input"
	readfile "cc/ccwc/read-file"
	readinput "cc/ccwc/read-input"
)

func main() {
	countPath, linePath := readinput.ReadInput()

	content := readfile.ReadFile(countPath, linePath)

	handleinput.HandleInput(content, countPath, linePath)

}
