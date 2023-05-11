package palindrome

import (
	"errors"
	"fmt"
)

var ErrInvalidPalindrome = errors.New("invalid palindrome")

// ValidPalindrome takes in a string (that only consists ASCII characters, which is why we can take
// the index of the string with no unintended edge case bugs) and determines using two pointers if
// the string is a palindrome or not. If its not an ErrInvalidPalindrome is returned.
// Time complexity is O(n) given we loop through all characters of the string.
// Space complexity is O(1) since we use constant space to store two indices.
func ValidPalindrome(s string) (bool, error) {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false, fmt.Errorf("string %s: %w", s, ErrInvalidPalindrome)
		}
		left++
		right--
	}
	return true, nil
}
