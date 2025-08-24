package palindromenumber

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isPalindrome",
			args: args{121},
			want: true,
		},
		{
			name: "isPalindrome",
			args: args{-121},
			want: false,
		},
		{
			name: "isNotPalindrome",
			args: args{224},
			want: false,
		},
		{
			name: "ten",
			args: args{10},
			want: false,
		},
		{
			name: "digits",
			args: args{8},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindromeNaive(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isPalindrome",
			args: args{121},
			want: true,
		},
		{
			name: "isPalindrome",
			args: args{-121},
			want: false,
		},
		{
			name: "isNotPalindrome",
			args: args{224},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindromeNaive(tt.args.x); got != tt.want {
				t.Errorf("IsPalindromeNaive() = %v, want %v", got, tt.want)
			}
		})
	}
}
