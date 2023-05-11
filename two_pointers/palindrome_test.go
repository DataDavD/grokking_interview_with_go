package palindrome

import (
	"errors"
	"testing"
)

func Test_validPalindrome(t *testing.T) {
	testCases := []struct {
		name    string
		testStr string
		want    bool
		wantErr error
	}{
		{
			name:    "valid palindrome 1",
			testStr: "davad",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "valid palindrome 2",
			testStr: "lettel",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "valid palindrome 3",
			testStr: "racecar",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "invalid palindrome",
			testStr: "david",
			want:    false,
			wantErr: ErrInvalidPalindrome,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidPalindrome(tt.testStr)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("validPalindrome() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("validPalindrome() got = %v, want %v", got, tt.want)
			}
		})
	}
}
