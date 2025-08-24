package removeelement

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive test",
			args: args{
				[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
				0,
			},
			want: 8,
		},
		{
			name: "positive test",
			args: args{
				[]int{0, 0},
				0,
			},
			want: 0,
		},
		{
			name: "positive test",
			args: args{
				[]int{0, 1},
				0,
			},
			want: 1,
		},
		{
			name: "positive test",
			args: args{
				[]int{1, 0},
				0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
