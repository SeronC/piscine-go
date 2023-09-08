package piscine

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func SplitWhiteSpaces(s string) []string {
	var words []string
	start := 0
	inWord := false

	for i := 0; i < len(s); i++ {
		if isWhitespace(s[i]) {
			if inWord {
				words = append(words, s[start:i])
				inWord = false
			}
		} else if !inWord {
			start = i
			inWord = true
		}
	}

	if inWord {
		words = append(words, s[start:])
	}

	return words
}
