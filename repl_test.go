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

		if len(actual) != len(c.expected) {
			t.Errorf("length of tested slice is different than the expected slice")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("tested word is different than the expected word")
			}
		}
	}
}
