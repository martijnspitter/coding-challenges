package main

import (
	handleinput "cc/ccwc/handle-input"
	readfile "cc/ccwc/read-file"
	readinput "cc/ccwc/read-input"
)

func main() {
	command, path := readinput.ReadInput()

	content := readfile.GetContent(path)

	handleinput.HandleInput(content, command, path)

}
