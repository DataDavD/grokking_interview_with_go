package two_pointers

import "testing"

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name    string
		testStr string
		want    string
	}{
		{
			name:    "valid palindrome 1",
			testStr: "Hello World",
			want:    "World Hello",
		},
		{
			name:    "valid palindrome 2",
			testStr: "We love Golang",
			want:    "Golang love We",
		},
		{
			name:    "valid palindrome 3",
			testStr: "Hello World",
			want:    "World Hello",
		},
		{
			name:    "valid palindrome 4",
			testStr: "To be or not to be",
			want:    "be to not or be To",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.testStr); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
