package breakingbestandworstrecord

import (
	"reflect"
	"testing"
)

func Test_breakingRecords(t *testing.T) {
	type args struct {
		scores []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "positive",
			args: args{
				scores: []int32{10, 5, 20, 20, 4, 5, 2, 25, 1},
			},
			want: []int32{2, 4},
		},
		{
			name: "positive",
			args: args{
				scores: []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42},
			},
			want: []int32{4, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := breakingRecords(tt.args.scores); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("breakingRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
