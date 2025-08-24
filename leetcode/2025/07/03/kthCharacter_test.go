package kthcharacter

import "testing"

func Test_kthCharacter(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "positive",
			args: args{
				k: 1,
			},
			want: 'a',
		},
		{
			name: "positive",
			args: args{
				k: 2,
			},
			want: 'b',
		},
		{
			name: "positive",
			args: args{
				k: 3,
			},
			want: 'b',
		},
		{
			name: "positive",
			args: args{
				k: 4,
			},
			want: 'c',
		},
		{
			name: "positive",
			args: args{
				k: 5,
			},
			want: 'b',
		},
		{
			name: "positive",
			args: args{
				k: 6,
			},
			want: 'c',
		},
		{
			name: "positive",
			args: args{
				k: 7,
			},
			want: 'c',
		},
		{
			name: "positive",
			args: args{
				k: 8,
			},
			want: 'd',
		},
		{
			name: "negative",
			args: args{
				k: 0,
			},
			want: 'a',
		},
		{
			name: "negative",
			args: args{
				k: -1,
			},
			want: 'a',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthCharacter(tt.args.k); got != tt.want {
				t.Errorf("kthCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
