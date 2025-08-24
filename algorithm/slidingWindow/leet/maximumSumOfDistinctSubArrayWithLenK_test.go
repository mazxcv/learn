package leet

import "testing"

func Test_maximumSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// nums = [1,5,4,2,9,9,9], k = 3.
		{
			name: "1",
			args: args{
				nums: []int{1, 5, 4, 2, 9, 9, 9},
				k:    3,
			},
			want: 15,
		},
		// [4,4,4], k = 3
		{
			name: "2",
			args: args{
				nums: []int{4, 4, 4},
				k:    3,
			},
			want: 0,
		},
		// [1,1,1,7,8,9], k = 3
		{
			name: "3",
			args: args{
				nums: []int{1, 1, 1, 7, 8, 9},
				k:    3,
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
