// Package challenge6 contains the solution for Challenge 6.
package challenge6

import (
    "strings"
    "unicode"
)

// CountWordFrequency takes a string containing multiple words and returns
// a map where each key is a word and the value is the number of times that
// word appears in the string. The comparison is case-insensitive.
//
// Words are defined as sequences of letters and digits.
// All words are converted to lowercase before counting.
// All punctuation, spaces, and other non-alphanumeric characters are ignored.
//
// For example:
// Input: "The quick brown fox jumps over the lazy dog."
// Output: map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1}
func CountWordFrequency(text string) map[string]int {
	// Your implementation here
	finalMap := make(map[string]int)

	var word strings.Builder
	
	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			word.WriteRune(unicode.ToLower(r))
		} else if r == '\'' {
			// Ignore apostrophes inside words.
			continue
		} else {
			if word.Len() > 0 {
				finalMap[word.String()]++
				word.Reset()
			}
		}


	
	
}
if word.Len() > 0 {
		finalMap[word.String()]++
	}

return finalMap
} 