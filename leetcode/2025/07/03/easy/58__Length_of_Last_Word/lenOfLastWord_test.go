package _8__Length_of_Last_Word

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantCnt int
	}{
		{
			name: "positive",
			args: args{
				s: "Hello World",
			},
			wantCnt: 5,
		},
		{
			name: "moon",
			args: args{
				s: "   fly me   to   the moon  ",
			},
			wantCnt: 4,
		},
		{
			name: "joyboy",
			args: args{
				s: "luffy is still joyboy",
			},
			wantCnt: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCnt := lengthOfLastWord(tt.args.s); gotCnt != tt.wantCnt {
				t.Errorf("lengthOfLastWord() = %v, want %v", gotCnt, tt.wantCnt)
			}
		})
	}
}
