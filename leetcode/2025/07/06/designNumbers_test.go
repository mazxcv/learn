package designnumbers

import (
	"reflect"
	"testing"
)

func Test_designnumbers(t *testing.T) {
	type args struct {
		command []string
		values  [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "positive",
			args: args{
				command: []string{"NumberContainers", "find", "change", "change", "change", "change", "find", "change", "find"},
				values:  [][]int{{}, {10}, {2, 10}, {3, 10}, {1, 10}, {5, 10}, {10}, {1, 20}, {10}},
			},
			want: []int{0, -1, 0, 0, 0, 0, 1, 0, 2},
		},
		{
			name: "positive",
			args: args{
				command: []string{"NumberContainers", "change", "find", "change", "find", "find", "find"},
				values:  [][]int{{}, {1, 10}, {10}, {1, 20}, {10}, {20}, {30}},
			},
			want: []int{0, 0, 1, 0, -1, 1, -1},
		},
		{
			name: "positive",
			args: args{
				command: []string{"NumberContainers", "change", "find", "find", "find", "change", "change", "change", "change", "change", "change", "change", "change", "change", "find", "change", "find", "find", "change", "find", "change"},
				// [[],[75,40],[14],[41],[40],[27,40],[22,14],[85,14],[22,40],[18,34],[92,41],[22,40],[75,40],[59,34],[40],[27,14],[34],[14],[13,34],[40],[18,41]]
				values: [][]int{{}, {75, 40}, {14}, {41}, {40}, {27, 40}, {22, 14}, {85, 14}, {22, 40}, {18, 34}, {92, 41}, {22, 40}, {75, 40}, {59, 34}, {40}, {27, 14}, {34}, {14}, {13, 34}, {40}, {18, 41}},
			},
			// [null,null,-1,-1,75,null,null,null,null,null,null,null,null,null,22,null,18,27,null,22,null]
			want: []int{0, 0, -1, -1, 75, 0, 0, 0, 0, 0, 0, 0, 0, 0, 22, 0, 18, 27, 0, 22, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := designnumbers(tt.args.command, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("designnumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
