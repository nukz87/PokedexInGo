package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"", nil},
		{"  ", nil},
		{"  hello  ", []string{"hello"}},
		{"  hello  world  ", []string{"hello", "world"}},
	}

	for i, c := range cases {
		actual := cleanInput(c.input)
		fmt.Printf("Test %v: ", i)
		if len(actual) != len(c.expected) {
			t.Errorf("not enough words\n")
			continue
		}
		ch := 1
		for i, _ := range actual {
			if actual[i] != c.expected[i] {
				ch = 0
				break
			}
		}
		if ch == 1 {
			fmt.Printf("AC\n")
		} else {
			t.Error("WA")
			t.Error("expect: ", c.expected)
			t.Error("actual: ", actual)
		}
	}
}
