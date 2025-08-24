package leet

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// s1 = "ab", s2 = "eidbaooo"
		{
			name: "1",
			args: args{
				s1: "ab",
				s2: "eidbaooo",
			},
			want: true,
		},
		// s1 = "ab", s2 = "eidboaoo"
		{
			name: "2",
			args: args{
				s1: "ab",
				s2: "eidboaoo",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
