package roman2integer

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "bul",
			args: args{s: ""},
			want: 0,
		},
		{
			name: "III",
			args: args{s: "III"},
			want: 3,
		},
		{
			name: "LVIII",
			args: args{s: "LVIII"},
			want: 58,
		},
		{
			name: "MCMXCIV",
			args: args{s: "MCMXCIV"},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
