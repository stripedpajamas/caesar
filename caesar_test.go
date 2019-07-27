package caesar_test

import (
	"fmt"
	"testing"

	"github.com/stripedpajamas/caesar"
)

func TestEncrypt(t *testing.T) {
	testCases := []struct {
		input    string
		key      int
		expected string
	}{
		{"abc", 1, "bcd"},
		{"zzz", 1, "aaa"},
		{"aZf", 1, "bAg"},
		{"123", 1, "123"},
		{"abc", 26, "abc"},
		{"abc", 27, "bcd"},
	}

	for idx, tc := range testCases {
		if actual := caesar.Encrypt(tc.input, tc.key); actual != tc.expected {
			fmt.Printf("(en) test %d failed: wanted %s, got %s\n", idx, tc.expected, actual)
			t.Fail()
		}
	}
}

func TestDecrypt(t *testing.T) {
	testCases := []struct {
		input    string
		key      int
		expected string
	}{
		{"bcd", 1, "abc"},
		{"zzz", 1, "yyy"},
		{"aZf", 1, "zYe"},
		{"123", 1, "123"},
		{"abc", 26, "abc"},
		{"abc", 27, "zab"},
	}

	for idx, tc := range testCases {
		if actual := caesar.Decrypt(tc.input, tc.key); actual != tc.expected {
			fmt.Printf("(de) test %d failed: wanted %s, got %s\n", idx, tc.expected, actual)
			t.Fail()
		}
	}
}
