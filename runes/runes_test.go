package runes_test

import (
	"fmt"
	"testing"

	"github.com/stripedpajamas/caesar/runes"
)

func TestIsLetter(t *testing.T) {
	testCases := []struct {
		input    rune
		expected bool
	}{
		{'a', true},
		{' ', false},
		{'B', true},
		{'$', false},
		{'0', false},
	}

	for _, tc := range testCases {
		if actual := runes.IsLetter(tc.input); actual != tc.expected {
			fmt.Printf("for input %v, wanted %v, got %v\n", tc.input, tc.expected, actual)
			t.Fail()
		}
	}
}

func TestIsUpper(t *testing.T) {
	testCases := []struct {
		input    rune
		expected bool
	}{
		{'a', false},
		{' ', false},
		{'B', true},
		{'$', false},
		{'0', false},
	}

	for _, tc := range testCases {
		if actual := runes.IsUpper(tc.input); actual != tc.expected {
			fmt.Printf("for input %v, wanted %v, got %v\n", tc.input, tc.expected, actual)
			t.Fail()
		}
	}
}

func TestIsLower(t *testing.T) {
	testCases := []struct {
		input    rune
		expected bool
	}{
		{'a', true},
		{' ', false},
		{'B', false},
		{'$', false},
		{'0', false},
	}

	for _, tc := range testCases {
		if actual := runes.IsLower(tc.input); actual != tc.expected {
			fmt.Printf("for input %v, wanted %v, got %v\n", tc.input, tc.expected, actual)
			t.Fail()
		}
	}
}

func TestToUpper(t *testing.T) {
	testCases := []struct {
		input    rune
		expected rune
	}{
		{'a', 'A'},
		{'B', 'B'},
		{'z', 'Z'},
		{'@', ' '},
	}

	for _, tc := range testCases {
		if actual := runes.ToUpper(tc.input); actual != tc.expected {
			fmt.Printf("for input %v, wanted %v, got %v\n", tc.input, tc.expected, actual)
			t.Fail()
		}
	}
}
