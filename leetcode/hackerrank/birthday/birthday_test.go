package birthday

import "testing"

func Test_birthday(t *testing.T) {
	type args struct {
		s []int32
		d int32
		m int32
	}
	tests := []struct {
		name    string
		args    args
		wantRes int32
	}{
		{
			name: "enumerates 1",
			args: args{
				s: []int32{1, 2, 1, 3, 2},
				d: 3,
				m: 2,
			},
			wantRes: 2,
		},
		{
			name: "enumerates 1",
			args: args{
				s: []int32{1, 1, 1, 1, 1, 1},
				d: 3,
				m: 2,
			},
			wantRes: 0,
		},
		{
			name: "enumerates 1",
			args: args{
				s: []int32{4},
				d: 4,
				m: 1,
			},
			wantRes: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := birthday(tt.args.s, tt.args.d, tt.args.m); gotRes != tt.wantRes {
				t.Errorf("birthday() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
