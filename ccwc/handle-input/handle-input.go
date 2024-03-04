package handleinput

func HandleInput(content []byte, command string, path string) {
	if command == "-c" {
		reportByteCount(content, path)
	}

	if command == "-l" {
		reportLineCount(content, path)
	}

	if command == "-w" {
		reportWordCount(content, path)
	}

	if command == "-m" {
		reportCharCount(content, path)
	}

	if command == "none" {
		reportNoOptionsCount(content, path)
	}

}
