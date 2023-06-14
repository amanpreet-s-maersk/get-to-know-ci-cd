package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 1},
		{1, 2, 1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d,%d", tt.a, tt.b), func(t *testing.T) {
			ans := Sum(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
