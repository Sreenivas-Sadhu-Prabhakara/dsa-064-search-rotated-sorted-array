package main

import "testing"

func TestSearchRotated(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 1, 0},
		{[]int{1}, 2, -1},
		{[]int{1, 2, 3, 4, 5}, 4, 3}, // not rotated at all
		{[]int{5, 1, 3}, 5, 0},
		{[]int{3, 1}, 1, 1},
		{[]int{}, 7, -1},
	}
	for _, c := range cases {
		if got := SearchRotated(c.in, c.target); got != c.want {
			t.Errorf("SearchRotated(%v,%d) = %d, want %d", c.in, c.target, got, c.want)
		}
	}
}
