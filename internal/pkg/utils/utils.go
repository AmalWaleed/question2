// pkg/utils/utils.go
package utils

// IsUnique checks if all characters in a string are unique
func IsUnique(s string) bool {
	charSet := make(map[rune]bool)

	for _, char := range s {
		if charSet[char] {
			return false
		}
		charSet[char] = true
	}

	return true
}
