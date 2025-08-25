package leet

import "testing"

func Test_hammingWeight(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// n = 111
		{
			name: "1",
			args: args{
				n: 11,
			},
			want: 3,
		},
		// n = 128
		{
			name: "2",
			args: args{
				n: 128,
			},
			want: 1,
		},
		// n = 2147483645
		{
			name: "3",
			args: args{
				n: 2147483645,
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.n); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
