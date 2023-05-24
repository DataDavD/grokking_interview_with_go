package two_pointers

// validPalindromeAdv takes a string as input and checks whether it is a valid palindrome with by
// removing, at most, one character.
// Time complexity is O(n) given we loop through all characters of the string.
// Space complexity is O(1) since we use constant space to store two indices.
func validPalindromeAdv(s string) (bool, error) {
	var isValid bool
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			// If the characters are not equal, then we need to check if removing either the left or right
			// character will result in a palindrome.
			isValid = isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
			if !isValid {
				return false, ErrInvalidPalindrome
			}
		}

		left++
		right--
	}

	return true, nil
}

// isPalindrome checks that the substring is a valid palindrome.
func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
