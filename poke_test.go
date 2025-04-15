package main

import (
	"strings"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{
			input:  "  hello world  ",
			output: []string{"hello", "world"},
		},

		{
			input:  "i don't know    what to type  ",
			output: []string{"i", "don't", "know", "what", "to", "type"},
		},

		{
			input:  " ",
			output: []string{},
		},
	}

	for _, testCase := range cases {
		funcOutput := cleanInput(testCase.input)
		if len(funcOutput) != len(testCase.output) {
			t.Errorf("Sizes don't match: %v vs %v", len(funcOutput), len(testCase.output))
		}

		for i := range funcOutput {
			if funcOutput[i] != testCase.output[i] {
				t.Errorf("Expected %s - Actual: %s", strings.Join(testCase.output, ","), strings.Join(funcOutput, ","))
			}
		}
	}

}
