package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go-reloaded/punctuations"
)

func main() {
	// Open the input file
	openFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	defer openFile.Close()
	// Read the file
	newScanner := bufio.NewScanner(openFile)
	if err != nil {
		fmt.Println(err)
	}
	// Create the output file
	outputFile, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}
	defer outputFile.Close()
	// fmt.Println("Text has been modified")

	var newScannerText string
	for newScanner.Scan() {
		newScannerText += newScanner.Text() + "\n"
	}

	words := strings.Fields(newScannerText)

	for i, word := range words {
		// Replace "a" with "an" if the next word starts with a vowel or "h"
		for i := 0; i < len(words); i++ {
			lowerWord := strings.ToLower(words[i])
			if lowerWord == "a" || lowerWord == "A" {
				nextWordfirstLetter := strings.ToLower(words[i+1][:1])
				if strings.ContainsAny(nextWordfirstLetter, "aeiouh") {
					if words[i] == "a" {
						words[i] = "an"
					} else if words[i] == "A" {
						words[i] = "An"
					}
				}
			}
		}

		if word == "(hex)" {
			decimal, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				fmt.Println("Error:", err)
			}
			words[i-1] = strconv.FormatInt(decimal, 10)
			words = append(words[:i], words[i+1:]...)
		} else if word == "(bin)" {
			decimal, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			words[i-1] = strconv.FormatInt(decimal, 10)
			words = append(words[:i], words[i+1:]...)
		} else if word == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if word == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if word == "(cap)" {
			words[i-1] = strings.Title(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if word == "(cap," {
			num, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
			if err != nil {
				fmt.Println(err)
			}
			for j := 0; j <= num; j++ {
				words[i-j] = strings.Title(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		} else if word == "(low," {
			num, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
			if err != nil {
				fmt.Println(err)
			}
			for j := 0; j <= num; j++ {
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		} else if word == "(up," {
			num, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
			if err != nil {
				fmt.Println("Error: ", err)
			}
			for j := 0; j <= num; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)

		}

	}

	newSentence := strings.Join((words), " ")
	newWord := punctuations.FormatPunctuation((newSentence))
	finalWords := punctuations.FormatSingleQuotes((newWord))

	doc := os.WriteFile(os.Args[2], []byte(finalWords+"\n"), 0o644)
	if doc != nil {
		panic(err)
	}
}

// func FormatPunctuation(text string) string {
// 	var result strings.Builder
// 	text = strings.TrimSpace(text) // Trim leading and trailing spaces
// 	length := len(text)

// 	for i := 0; i < length; i++ {
// 		current := text[i]
// 		next := byte(' ')
// 		if i+1 < length {
// 			next = text[i+1]
// 		}

// 		if strings.ContainsRune(".,:;!?-", rune(current)) {
// 			if current == '.' && i+2 < length && text[i+1] == '.' && text[i+2] == '.' {
// 				// Handle ellipsis
// 				result.WriteString("...")
// 				i += 2 // Skip next two dots
// 			} else {
// 				result.WriteByte(current) // Append the punctuation
// 			}

// 			// Ensure space after punctuation if the next character isn't punctuation or space
// 			if !strings.ContainsRune(" ,.:;!?-", rune(next)) {
// 				result.WriteByte(' ')
// 			}
// 		} else {
// 			// Add current character if it's not an unwanted space
// 			if !(current == ' ' && strings.ContainsRune(".,:;!?-", rune(next))) {
// 				result.WriteByte(current)
// 			}
// 		}
// 	}

// 	return result.String()
// }

// func FormatSingleQuotes(text string) string {
// 	var result strings.Builder
// 	length := len(text)
// 	i := 0

// 	for i < length {
// 		current := text[i]

// 		if current == '\'' && i+1 < length {
// 			start := i + 1
// 			end := strings.IndexByte(text[start:], '\'') + start
// 			if end > start && end < length {
// 				// Trim spaces around the quoted text
// 				quotedText := strings.TrimSpace(text[start:end])
// 				result.WriteByte('\'') // Write opening quote
// 				result.WriteString(quotedText)
// 				result.WriteByte('\'') // Write closing quote
// 				i = end + 1            // Skip past the closing quote
// 				continue
// 			}
// 		}

// 		// If not processing quotes, append the character as is
// 		result.WriteByte(current)
// 		i++
// 	}

// 	return result.String()
// }
