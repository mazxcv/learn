package randomwait

import (
	"fmt"
	"testing"
)

func Test_rnd(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		{
			name:  "positive test",
			want:  6,
			want1: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := rnd()
			fmt.Println(got, got1)
			fmt.Println(tt.want, tt.want1)
			if got > tt.want {
				t.Errorf("rnd() got = %v, want %v", got, tt.want)
			}
			if got1 > tt.want1 {
				t.Errorf("rnd() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_chRnd(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		{
			name:  "positive test",
			want:  6,
			want1: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := chRnd()
			fmt.Println(got, got1)
			fmt.Println(tt.want, tt.want1)
			if got > tt.want {
				t.Errorf("chRnd() got = %v, want %v", got, tt.want)
			}
			if got1 > tt.want1 {
				t.Errorf("chRnd() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
