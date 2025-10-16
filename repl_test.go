package main 

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input     string 
		expected  []string 
	}{
		{
			input: "  ",
			expected: []string{},
		},
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: " bulbasaur charmander ",
			expected: []string{"bulbasaur", "charmander"},
		},
		{
			input: " suicune raikou entei ",
			expected: []string{"suicune", "raikou", "entei"},
		},
		{
			input: " zapdos moltres articuno palkia dialga ",
			expected: []string{"zapdos", "moltres", "articuno", "palkia", "dialga"},
		},
	}
	
	for _, c := range cases {
		actual := cleanInput(c.input)
		// check the length of the actual slice against the expected slice 
		// if they don't match, use t.Errorf to print an error message 
		// and fail the test 
		if len(actual) != len(c.expected) {
			t.Errorf("length of tested slice is different than the expected slice")
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// check each word in the slice 
			// if they don't match, use t.Errorf to print an error message 
			// and fail the test 
			if word != expectedWord {
				t.Errorf("tested word is different than the expected word")
				t.Fail()
			}
		}
	}
}
