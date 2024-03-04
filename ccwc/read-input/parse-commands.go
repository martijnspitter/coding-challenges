package readinput

import (
	"log"
	"os"
	"strings"
)

func ReadInput() (string, string) {
	args := os.Args

	var command, path string

	switch len(args) {
	case 1:
		log.Fatal("positional arg1 is required")
	case 2:
		command = getCommand(args[1])
		path = ""
	case 3:
		command = getCommand(args[1])
		path = getPath(args[2])
	default:
		command = "none"
		path = ""
	}

	return command, path
}

func getCommand(arg string) string {
	if strings.Contains(arg, "-") {
		return arg
	}

	return "none"
}

func getPath(arg string) string {
	if len(arg) > 0 {
		return arg
	}

	return ""
}
