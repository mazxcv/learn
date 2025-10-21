package majorityelement

import "testing"

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			name: "Example 2",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
