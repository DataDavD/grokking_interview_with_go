package palindrome

import (
	"regexp"
	"strings"
)

// reverseStr reverses a string character by character and returns a byte slice.
// Note, this is pretty poor because we change the string in place, and also return it. But it gets
// the job quickly :).
func reverseStr(str []byte, start int, end int) []byte {
	for start < end {
		temp := str[start] // temp store for swapping
		str[start] = str[end]
		str[end] = temp

		start++ // Move forwards towards the center
		end--   // Move backwards towards the center
	}

	return str
}

func trimString(str string) string {
	// Trim extra spaces from end and beginning of input string
	str = strings.TrimSpace(str)

	// Replace multiple spaces with a single space
	regex := regexp.MustCompile("\\s+")
	str = regex.ReplaceAllString(str, " ")

	return str
}

// reverseWords reverses the order of a sentence's words without affecting the order of the letters
// in the given words. Input sentence must only contain English uppercase and lowercase letters,
// digits, and spaces.
// Note, the input string may contain leading or trailing spaces (or multiple spaces), but the
// reverseWords returns a string with only single spaces separating the words.
// Time complexity is O(n).
// Space complexity is O(n) since we use additional space (new byte slice to store the reversed
// string, and this byte slice takes O(n) space) to store the list of characters for reversal since
// strings are immutable in Go.
func reverseWords(sentence string) string {
	// Remove leading, trailing, and multiple spaces from the input sentence.
	trimmedSentence := trimString(sentence)

	// Convert input string to byte slice since strings are immutable in Go
	sentenceBytes := []byte(trimmedSentence)

	// Reverse the entire string.
	reversedSentence := reverseStr(sentenceBytes, 0, len(sentenceBytes)-1)

	for start, end := 0, 0; start < len(reversedSentence); {
		// Move end forward until we reach a space/end of the word/end of the sentence.
		for end < len(reversedSentence) && sentenceBytes[end] != ' ' {
			end++
		}

		reverseStr(sentenceBytes, start, end-1) // Reverse word
		start = end + 1
		end += 1
	}

	return string(sentenceBytes)
}
