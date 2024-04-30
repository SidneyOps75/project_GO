package punctuations

import "strings"

func FormatSingleQuotes(text string) string {
	var result strings.Builder
	length := len(text)
	i := 0

	for i < length {
		current := text[i]

		if current == '\'' && i+1 < length {
			start := i + 1
			end := strings.IndexByte(text[start:], '\'') + start
			if end > start && end < length {
				// Trim spaces around the quoted text
				quotedText := strings.TrimSpace(text[start:end])
				result.WriteByte('\'') // Write opening quote
				result.WriteString(quotedText)
				result.WriteByte('\'') // Write closing quote
				i = end + 1            // Skip past the closing quote
				continue
			}
		}

		// If not processing quotes, append the character as is
		result.WriteByte(current)
		i++
	}

	return result.String()
}
