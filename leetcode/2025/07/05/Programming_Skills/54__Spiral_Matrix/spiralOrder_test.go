package _4__Spiral_Matrix

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "null",
			args: args{
				matrix: [][]int{},
			},
			want: []int{},
		},
		{
			name: "zero",
			args: args{
				matrix: [][]int{{0}},
			},
			want: []int{0},
		},
		{
			name: "zero1",
			args: args{
				matrix: [][]int{{0}, {0}},
			},
			want: []int{0, 0},
		},
		{
			name: "zero2",
			args: args{
				matrix: [][]int{{0, 0}},
			},
			want: []int{0, 0},
		},
		{
			name: "zero2",
			args: args{
				matrix: [][]int{{0, 0}, {0, 0}},
			},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "zero2",
			args: args{
				matrix: [][]int{{0, 0}, {0, 0}, {0, 0}},
			},
			want: []int{0, 0, 0, 0, 0, 0},
		},
		{
			name: "zero2",
			args: args{
				matrix: [][]int{{0, 0, 0}, {0, 0, 0}},
			},
			want: []int{0, 0, 0, 0, 0, 0},
		},
		{
			name: "one",
			args: args{
				matrix: [][]int{{1}},
			},
			want: []int{1},
		},
		{
			name: "12",
			args: args{
				matrix: [][]int{{1}, {2}},
			},
			want: []int{1, 2},
		},
		{
			name: "12",
			args: args{
				matrix: [][]int{{1, 2}},
			},
			want: []int{1, 2},
		},
		{
			name: "1234",
			args: args{
				matrix: [][]int{{1, 2}, {4, 3}},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "123456789",
			args: args{
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "123456789101112",
			args: args{
				matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
