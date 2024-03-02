package handleinput

func HandleInput(content []byte, countPath *string, linePath *string, wordPath *string) {
	if *countPath != "" {
		reportByteCount(content, countPath)
	}

	if *linePath != "" {
		reportLineCount(content, linePath)
	}

	if *wordPath != "" {
		reportWordCount(content, wordPath)
	}

}
