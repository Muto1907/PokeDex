package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "\t EYoO\nWHaT's       Up",
			expected: []string{"eyoo", "what's", "up"},
		},
		{
			input: `1
								I dOn32T
			KNowwwww


			}[€     ]`,
			expected: []string{"1", "i", "don32t", "knowwwww", "}[€", "]"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "               \n  \n\t\t",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Error unexpected length of inputslice. Expected: %d got: %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Error unexpected word in slice. Expected: %s got: %s", expectedWord, word)
			}
		}
	}
}
