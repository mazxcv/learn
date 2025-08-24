package _83_Move_Zeroes

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "[0]",
			args: args{nums: []int{0}},
		},
		{
			name: "[1]",
			args: args{nums: []int{1}},
		},
		{
			name: "[]",
			args: args{nums: []int{}},
		},
		{
			name: "[1,2,3,4]",
			args: args{nums: []int{1, 2, 3, 4}},
		},
		{
			name: "[0, 0,0,0,]",
			args: args{nums: []int{0, 0, 0, 0}},
		},
		{
			name: "[0,0,0,1]",
			args: args{nums: []int{0, 0, 0, 1}},
		},
		{
			name: "[0,1,2,3,0,5,6,0,7,0]",
			args: args{nums: []int{0, 1, 2, 3, 0, 5, 6, 0, 7, 0}},
		},
		{
			name: "[0,1,2,3,0,5,6,0,7]",
			args: args{nums: []int{0, 1, 2, 3, 0, 5, 6, 0, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
		})
	}
}
