package main

import "testing"

func TestNoOfTest(t *testing.T) {

	tests := []struct {
		name   string
		input  int
		output int
	}{
		{"n1", 1, 1},
		{"n2", 2, 2},
		{"n3", 3, 4},
		{"n4", 4, 7},
		{"n10", 10, 7},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := noofways(test.input)
			if out != test.output {
				t.Errorf("expected %d got %d", out, test.output)
			}
		})
	}

}
