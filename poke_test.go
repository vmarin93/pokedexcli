package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "leading and trailing spaces",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "no separator",
			input:    "happydinosaur",
			expected: []string{"happydinosaur"},
		},
		{
			name:     "uppercasing in words",
			input:    "ARMAGHEDON google",
			expected: []string{"armaghedon", "google"},
		},
	}

	for _, tc := range cases {
		actual := cleanInput(tc.input)
		if len(actual) != len(tc.expected) {
			t.Errorf("test %s: len of actual %d doesn't match len of expected %d", tc.name, len(actual), len(tc.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := tc.expected[i]
			if word != expectedWord {
				t.Errorf("test %s: word in actual %s doesn't match word in expected %s", tc.name, word, expectedWord)
			}
		}
	}
}
