package two_pointers

import (
	"errors"
	"testing"
)

func Test_validPalindromeAdv(t *testing.T) {
	testCases := []struct {
		name    string
		testStr string
		want    bool
		wantErr error
	}{
		{
			name:    "valid palindrome #1",
			testStr: "davad",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "valid palindrome #2",
			testStr: "lettel",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "valid palindrome #3",
			testStr: "RACEACAR",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "invalid palindrome #1",
			testStr: "davzziad",
			want:    false,
			wantErr: ErrInvalidPalindrome,
		},
		{
			name:    "inavlid palindrome #2",
			testStr: "abdca",
			want:    false,
			wantErr: ErrInvalidPalindrome,
		},
		{
			name:    "invalid palindrome #3",
			testStr: "eeccccbebaeddeabebccceea",
			want:    false,
			wantErr: ErrInvalidPalindrome,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validPalindromeAdv(tt.testStr)
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
