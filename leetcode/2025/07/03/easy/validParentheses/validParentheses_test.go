package validparentheses

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "positive",
			args: args{s: "()"},
			want: true,
		},
		{
			name: "positive",
			args: args{s: "()[]{}"},
			want: true,
		},
		{
			name: "negative",
			args: args{s: "(]"},
			want: false,
		},
		{
			name: "postive",
			args: args{s: "([])"},
			want: true,
		},
		{
			name: "postive",
			args: args{s: "()[]{}"},
			want: true,
		},
		{
			name: "empty",
			args: args{s: ""},
			want: true,
		},
		{
			name: "x",
			args: args{s: "([)]"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
