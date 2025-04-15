package main

import (
	"strings"
	"testing"
	"time"

	"github.com/Ciobi0212/pokedex/internal/pokecache"
	"github.com/Ciobi0212/pokedex/internal/utils"
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
		funcOutput := utils.CleanInput(testCase.input)
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

func TestReadLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
