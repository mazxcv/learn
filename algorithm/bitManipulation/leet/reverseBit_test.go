package leet

import "testing"

func Test_reverseBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//  n = 43261596
		{
			name: "1",
			args: args{
				n: 43261596,
			},
			want: 964176192,
		},
		// n = 2147483644
		{
			name: "2",
			args: args{
				n: 2147483644,
			},
			want: 1073741822,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.n); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
