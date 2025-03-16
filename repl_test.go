package main

import "testing"

func TestCleanInput(t *testing.T) {
    cases := []struct{
        input string
        expected []string
    }{
        {
            input: "  hello  world  ",
            expected: []string{"hello", "world"},
        },
        {
            input: "Charmander Bulbasaur PIKACHU",
            expected: []string{"charmander", "bulbasaur", "pikachu"},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("slice lengths don't match. actual: %s expected: %s", actual, c.expected)
            continue
        }
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("%d words don't match. actual: %s expected: %s", i, word, expectedWord)
            }
        }
    }
}
