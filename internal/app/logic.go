// internal/app/logic.go
package app

import (
	"sort"
	"github.com/your-username/your-project/pkg/utils"

)

// ReorganizeString rearranges the characters of s so that any two adjacent characters are not the same
func (a *App) ReorganizeString(s string) string {
	charCount := make(map[rune]int)

	// Count the frequency of each character
	for _, char := range s {
		charCount[char]++
	}

	// Create a slice of structs to store characters and their frequencies
	var charFreq []struct {
		char  rune
		count int
	}

	for char, count := range charCount {
		charFreq = append(charFreq, struct{ char rune; count int }{char, count})
	}

	// Sort characters by frequency in descending order
	sort.Slice(charFreq, func(i, j int) bool {
		return charFreq[i].count > charFreq[j].count
	})

	// Create a result slice to store the rearranged characters
	result := make([]rune, len(s))
	index := 0

	// Rearrange characters in the result slice
	for _, cf := range charFreq {
		for cf.count > 0 {
			result[index] = cf.char
			index += 2

			// Wrap around if the index goes beyond the length of the result slice
			if index >= len(result) {
				index = 1
			}

			cf.count--
		}
	}

	// Check if any two adjacent characters are the same
	for i := 1; i < len(result); i++ {
		if result[i] == result[i-1] {
			return ""
		}
	}
	
		// Check if any two adjacent characters are the same using the IsUnique utility function
		if !utils.IsUnique(string(result)) {
			return ""
		}
	

	return string(result)
}
