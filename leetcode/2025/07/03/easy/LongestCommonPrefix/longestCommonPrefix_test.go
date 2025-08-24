package longestcommonprefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "negative",
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			name: "positive",
			args: args{strs: []string{"flow", "flenn", "flamingo"}},
			want: "fl",
		},
		{
			name: "including",
			args: args{strs: []string{"ab", "a"}},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
