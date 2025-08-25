package leet

import "testing"

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// [2,2,1]
		{
			name: "1",
			args: args{
				nums: []int{2, 2, 1},
			},
			want: 1,
		},
		// [4,1,2,1,2]
		{
			name: "2",
			args: args{
				nums: []int{4, 1, 2, 1, 2},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
