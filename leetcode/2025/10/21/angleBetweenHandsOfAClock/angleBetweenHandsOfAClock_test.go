package anglebetweenhandsofaclock

import "testing"

func Test_angleClock(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		hour    int
		minutes int
		want    float64
	}{
		{
			name:    "Example 1",
			hour:    12,
			minutes: 30,
			want:    165,
		},
		{
			name:    "Example 2",
			hour:    3,
			minutes: 30,
			want:    75,
		},
		{
			name:    "Example 3",
			hour:    3,
			minutes: 15,
			want:    7.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := angleClock(tt.hour, tt.minutes)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("angleClock() = %v, want %v", got, tt.want)
			}
		})
	}
}
