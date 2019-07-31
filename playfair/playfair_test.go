package playfair_test

import (
	"fmt"
	"testing"

	"github.com/stripedpajamas/caesar/playfair"
)

func TestEncrypt(t *testing.T) {
	// example from https://en.wikipedia.org/wiki/Playfair_cipher#Example
	testCases := []struct {
		input    string
		key      string
		expected string
	}{
		{
			"abx",
			"playfair example",
			"PDGW",
		},
		{
			"abc",
			"playfair example",
			"PDGR",
		},
		{
			"Hide the gold in the tree stump",
			"playfair example",
			"BMODZBXDNABEKUDMUIXMMOUVIF",
		},
	}

	for idx, tc := range testCases {
		if actual := playfair.Encrypt(tc.input, tc.key); actual != tc.expected {
			fmt.Printf("(en) test %d failed: wanted %s, got %s\n", idx, tc.expected, actual)
			t.Fail()
		}
	}
}

func TestDecrypt(t *testing.T) {
	testCases := []struct {
		input    string
		key      string
		expected string
	}{
		{
			"BMODZBXDNABEKUDMUIXMMOUVIF",
			"playfair example",
			"HIDETHEGOLDINTHETREXESTUMP",
		},
	}

	for idx, tc := range testCases {
		if actual := playfair.Decrypt(tc.input, tc.key); actual != tc.expected {
			fmt.Printf("(de) test %d failed: wanted %s, got %s\n", idx, tc.expected, actual)
			t.Fail()
		}
	}
}
