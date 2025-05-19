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
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
	}

	passCount := 0
	failCount := 0

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			failCount++
			t.Errorf(`---------------------------------
Slice length doesn't match
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, c.input, len(c.expected), len(actual))

			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				failCount++
				t.Errorf(`---------------------------------
Word doesn't match
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, c.input, c.expected, actual)

				continue
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, c.input, c.expected, actual)
			}
		}
	}
	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
