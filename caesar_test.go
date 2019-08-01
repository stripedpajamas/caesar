package caesar_test

import (
	"fmt"
	"testing"

	"github.com/stripedpajamas/caesar"
)

func TestCaesarEncrypt(t *testing.T) {
	c := caesar.Caesar{}
	testCases := []struct {
		input       string
		key         string
		expected    string
		expectedErr bool
	}{
		{"abc", "1", "bcd", false},
		{"zzz", "1", "aaa", false},
		{"aZf", "1", "bAg", false},
		{"123", "1", "123", false},
		{"abc", "26", "abc", false},
		{"abc", "27", "bcd", false},
		{"abc", "$", "", true},
		{"abc", "", "", true},
	}

	for idx, tc := range testCases {
		actual, err := c.Encrypt(tc.input, tc.key)
		if tc.expectedErr && err == nil {
			fmt.Printf("(en) test %d failed: wanted error, got success", idx)
			t.Fail()
		}
		if actual != tc.expected {
			fmt.Printf("(en) test %d failed: wanted %s, got %s\n", idx, tc.expected, actual)
			t.Fail()
		}
	}
}

func TestCaesarDecrypt(t *testing.T) {
	c := caesar.Caesar{}
	testCases := []struct {
		input       string
		key         string
		expected    string
		expectedErr bool
	}{
		{"bcd", "1", "abc", false},
		{"zzz", "1", "yyy", false},
		{"aZf", "1", "zYe", false},
		{"123", "1", "123", false},
		{"abc", "26", "abc", false},
		{"abc", "27", "zab", false},
		{"abc", "$", "", true},
		{"abc", "", "", true},
	}

	for idx, tc := range testCases {
		actual, err := c.Decrypt(tc.input, tc.key)
		if tc.expectedErr && err == nil {
			fmt.Printf("(de) test %d failed: wanted error, got success", idx)
			t.Fail()
		}
		if actual != tc.expected {
			fmt.Printf("(de) test %d failed: wanted %s, got %s\n", idx, tc.expected, actual)
			t.Fail()
		}
	}
}
