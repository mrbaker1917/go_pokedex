package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "how are you today?  ",
			expected: []string{"how", "are", "you", "today?"},
		},
		{
			input:    " Charmander   Bulbasaur PIKACHU  ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual and expected have different lengths: FAIL!")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf(" the actual word %s and the expected word %s do not match: FAIL!", word, expectedWord)
			}
		}
	}
}
