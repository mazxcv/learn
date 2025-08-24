package kthcharacter

import "testing"

func Test_kthCharacter(t *testing.T) {
	type args struct {
		k          int64
		operations []int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "positive",
			args: args{
				k:          1,
				operations: []int{0},
			},
			want: 'a',
		},
		{
			name: "positive",
			args: args{
				k:          2,
				operations: []int{0},
			},
			want: 'a',
		},
		{
			name: "positive",
			args: args{
				k:          2,
				operations: []int{0, 0, 0},
			},
			want: 'a',
		},
		{
			name: "positive",
			args: args{
				k:          3,
				operations: []int{0, 1, 0},
			},
			want: 'b',
		},
		{
			name: "positive",
			args: args{
				k:          4,
				operations: []int{0, 1, 0, 1},
			},
			want: 'b',
		},
		{
			name: "positive",
			args: args{
				k:          5,
				operations: []int{0, 1, 0, 1, 1},
			},
			want: 'c',
		},
		{
			name: "positive",
			args: args{
				k:          30654210,
				operations: []int{1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0},
			},
			want: 'k',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthCharacter(tt.args.k, tt.args.operations); got != tt.want {
				t.Errorf("kthCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
