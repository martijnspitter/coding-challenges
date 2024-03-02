package handleinput

func HandleInput(content []byte, countPath *string, linePath *string) {
	if *countPath != "" {
		reportByteCount(content, countPath)
	}

	if *linePath != "" {
		reportLineCount(content, linePath)
	}

}
