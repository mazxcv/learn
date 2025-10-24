package longestpalindrome

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive test",
			args: args{
				"abccccdd",
			},
			want: 7,
		},
		{
			name: "positive test",
			args: args{
				"a",
			},
			want: 1,
		},
		{
			name: "positive test",
			args: args{
				"bb",
			},
			want: 2,
		},
		{
			name: "positive test",
			args: args{
				"ccc",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
