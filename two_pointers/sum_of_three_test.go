package two_pointers

import "testing"

func Test_threeSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true: only positive numbers",
			args: args{
				nums:   []int{3, 7, 1, 2, 8, 4, 6},
				target: 21,
			},
			want: true,
		},
		{
			name: "true: with negative numbers",
			args: args{
				nums:   []int{-1, 2, 1, -4, 5, -3},
				target: -8,
			},
			want: true,
		},
		{
			name: "false: only positive numbers",
			args: args{
				nums:   []int{3, 5, 1, -1, 9},
				target: 30,
			},
			want: false,
		},
		{
			name: "false: with negative numbers",
			args: args{
				nums:   []int{-1, 1, 3},
				target: 0,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findSumOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
