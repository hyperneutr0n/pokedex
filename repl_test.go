package main

import (
	"fmt"
	"testing"
)

func printPass(input string, expected, actual []string) {
	fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, input, expected, actual)
}

func printFail(input string, expected, actual interface{}, reason string, t *testing.T) {
	t.Errorf(`---------------------------------
%s
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, reason, input, expected, actual)
}

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Go is AWESOME",
			expected: []string{"go", "is", "awesome"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
	}

	passCount := 0
	failCount := 0

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			failCount++
			printFail(c.input, len(c.expected), len(actual), "Slice length doesn't match", t)
			continue
		}

		fail := false
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				fail = true
				break
			}
		}

		if fail {
			failCount++
			printFail(c.input, c.expected, actual, "Word doesn't match", t)
			continue
		} else {
			passCount++
			printPass(c.input, c.expected, actual)
		}
	}
	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

