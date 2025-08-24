package leet

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// s = "cbaebabacd", p = "abc"
		{
			name: "1",
			args: args{
				s: "cbaebabacd",
				p: "abc",
			},
			want: []int{0, 6},
		},
		// s = "abab", p = "ab"
		{
			name: "2",
			args: args{
				s: "abab",
				p: "ab",
			},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
