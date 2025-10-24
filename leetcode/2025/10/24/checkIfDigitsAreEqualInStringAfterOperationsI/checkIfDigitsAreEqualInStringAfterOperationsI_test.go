package checkifdigitsareequalinstringafteroperationsi

import "testing"

func Test_hasSameDigits(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "3902",
			want: true,
		},
		{
			name: "Example 2",
			s:    "34789",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasSameDigits(tt.s)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("hasSameDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
