package punctuations

import "strings"

func FormatPunctuation(text string) string {
	var result strings.Builder
	text = strings.TrimSpace(text) // Trim leading and trailing spaces
	length := len(text)

	for i := 0; i < length; i++ {
		current := text[i]
		next := byte(' ')
		if i+1 < length {
			next = text[i+1]
		}

		if strings.ContainsRune(".,:;!?-", rune(current)) {
			if current == '.' && i+2 < length && text[i+1] == '.' && text[i+2] == '.' {
				// Handle ellipsis
				result.WriteString("...")
				i += 2 // Skip next two dots
			} else {
				result.WriteByte(current) // Append the punctuation
			}

			// Ensure space after punctuation if the next character isn't punctuation or space
			if !strings.ContainsRune(" ,.:;!?-", rune(next)) {
				result.WriteByte(' ')
			}
		} else {
			// Add current character if it's not an unwanted space
			if !(current == ' ' && strings.ContainsRune(".,:;!?-", rune(next))) {
				result.WriteByte(current)
			}
		}
	}

	return result.String()
}
