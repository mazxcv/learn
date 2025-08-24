package findsumpairs

import (
	"reflect"
	"testing"
)

func Test_findSumPairs(t *testing.T) {
	type args struct {
		init     [][]int
		commands []string
		values   [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			//  ["FindSumPairs","count","count"]
			// [[[1,2],[2,3]],[3],[5]]
			// [null,1, 1]
			name: "positive",
			args: args{
				init: [][]int{
					{1, 2},
					{2, 3},
				},
				commands: []string{"count", "count"},
				values: [][]int{
					{3},
					{5},
				},
			},
			want: []int{0, 1, 1},
		}, {
			// ["FindSumPairs","count","add","count","count","add","add","count"]
			// [[[1,1,2,2,2,3],[1,4,5,2,5,4]],[7],[3,2],[8],[4],[0,1],[1,1],[7]]
			// [null,8,null,2,1,null,null,11]
			name: "positive",
			args: args{
				init: [][]int{
					{1, 1, 2, 2, 2, 3},
					{1, 4, 5, 2, 5, 4},
				},
				commands: []string{"count", "add", "count", "count", "add", "add", "count"},
				values: [][]int{
					{7},
					{3, 2},
					{8},
					{4},
					{0, 1},
					{1, 1},
					{7},
				},
			},
			want: []int{0, 8, 0, 2, 1, 0, 0, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSumPairs(tt.args.init, tt.args.commands, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSumPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
