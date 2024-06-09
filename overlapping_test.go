package main

import "testing"

func TestOverlapping(t *testing.T) {
	xs := []Tuple[int, int]{{3, 4}, {1, 6}, {0, 7}}
	ys := []Tuple[int, int]{{4, 5}, {0, 3}, {1, 9}, {7, 8}, {2, 6}}

	cases := []struct {
		in   []Tuple[int, int]
		want int
	}{
		{xs, 3},
		{ys, 2},
	}
	for _, c := range cases {
		got := Overlapping(c.in)
		if got != c.want {
			t.Errorf("Overlapping(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
