package migratorybirds

import "testing"

func Test_migratoryBirds(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name        string
		args        args
		wantMinType int32
	}{
		{
			name: "positive",
			args: args{
				arr: []int32{1, 4, 4, 4, 5, 3},
			},
			wantMinType: 4,
		},
		{
			name: "positive",
			args: args{
				arr: []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4},
			},
			wantMinType: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMinType := migratoryBirds(tt.args.arr); gotMinType != tt.wantMinType {
				t.Errorf("migratoryBirds() = %v, want %v", gotMinType, tt.wantMinType)
			}
		})
	}
}
