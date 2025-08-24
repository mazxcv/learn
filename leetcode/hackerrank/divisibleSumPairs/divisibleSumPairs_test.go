package divisiblesumpairs

import "testing"

func Test_divisibleSumPairs(t *testing.T) {
	type args struct {
		n  int32
		k  int32
		ar []int32
	}
	tests := []struct {
		name    string
		args    args
		wantRes int32
	}{
		{
			name: "positive",
			args: args{
				n:  6,
				k:  3,
				ar: []int32{1, 3, 2, 6, 1, 2},
			},
			wantRes: 5,
		},
		{
			name: "with negative",
			args: args{
				n:  6,
				k:  3,
				ar: []int32{-1, -3, -2, -6, -1, -2},
			},
			wantRes: 5,
		},
		{
			name: "positive and negative",
			args: args{
				n:  6,
				k:  3,
				ar: []int32{-1, 3, -2, 6, -1, 2},
			},
			wantRes: 4,
		},
		{
			name: "zero",
			args: args{
				n:  1,
				k:  3,
				ar: []int32{0},
			},
			wantRes: 0,
		},
		{
			name: "empty",
			args: args{
				n:  0,
				k:  3,
				ar: []int32{},
			},
			wantRes: 0,
		},
		{
			name: "zero k",
			args: args{
				n:  0,
				k:  0,
				ar: []int32{},
			},
			wantRes: 0,
		},
		{
			name: "zero k",
			args: args{
				n:  1,
				k:  0,
				ar: []int32{0},
			},
			wantRes: 0,
		},
		{
			name: "zero k zero value",
			args: args{
				n:  2,
				k:  0,
				ar: []int32{0, 0},
			},
			wantRes: 0,
		},
		{
			name: "normal k zero value",
			args: args{
				n:  2,
				k:  2,
				ar: []int32{0, 0},
			},
			wantRes: 1,
		},
		{
			name: "normal k zero value",
			args: args{
				n:  4,
				k:  4,
				ar: []int32{0, 0, 0, 0},
			},
			wantRes: 6,
		},
		{
			name: "k = 2 1",
			args: args{
				n:  4,
				k:  2,
				ar: []int32{1, 1, 1, 1},
			},
			wantRes: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := divisibleSumPairs(tt.args.n, tt.args.k, tt.args.ar); gotRes != tt.wantRes {
				t.Errorf("divisibleSumPairs() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
