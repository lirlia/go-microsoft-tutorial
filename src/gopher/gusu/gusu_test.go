package main

import (
	"testing"
)

func TestGusu(t *testing.T) {
	cases := []struct {
		name     string
		input    int
		expected bool
	}{
		{name: "+odd", input: 5, expected: false},
		{name: "+even", input: 6, expected: true},
		{name: "-odd", input: -5, expected: false},
		{name: "-even", input: -6, expected: true},
		{name: "zero", input: 0, expected: true},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := gusuchk(c.input); c.expected != actual {
				t.Errorf("want IsOdd(%d) = %v, got %v", c.input, c.expected, actual)
			}
		})
	}
}
