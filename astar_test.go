package main

import (
	"testing"
)

func TestGetBackRoute(t *testing.T) {
	start := newPoint(0, 0, 0, nil)
	first := newPoint(1, 1, 0, &start)
	second := newPoint(2, 2, 0, &first)
	third := newPoint(3, 3, 0, &second)
	fourth := newPoint(2, 3, 0, &second)

	var tests = []struct{
		start *Point
		end *Point
		route int
	}{
		{
			&start,
			&start,
			0,
		},
		{
			&start,
			&third,
			42,
		},
		{
			&start,
			&fourth,
			38,
		},
	}

	for _, test := range tests {
		got := getBackRoute(test.start, test.end)
		if got != test.route {
			t.Errorf("Wrong back route: got %v, want: %v", got, test.route)
		}
	}
}
