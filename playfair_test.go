package caesar_test

import (
	"fmt"
	"testing"

	"github.com/stripedpajamas/caesar"
)

func TestPlayfairEncrypt(t *testing.T) {
	pf := caesar.Playfair{}
	// example from https://en.wikipedia.org/wiki/Playfair_cipher#Example
	testCases := []struct {
		input       string
		key         string
		expected    string
		expectedErr bool
	}{
		{
			"abx",
			"playfair example",
			"PDGW",
			false,
		},
		{
			"abc",
			"playfair example",
			"PDGR",
			false,
		},
		{
			"Hide the gold in the tree stump",
			"playfair example",
			"BMODZBXDNABEKUDMUIXMMOUVIF",
			false,
		},
	}

	for idx, tc := range testCases {
		actual, err := pf.Encrypt(tc.input, tc.key)
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

func TestPlayfairDecrypt(t *testing.T) {
	pf := caesar.Playfair{}
	testCases := []struct {
		input       string
		key         string
		expected    string
		expectedErr bool
	}{
		{
			"PDGW",
			"playfair example",
			"ABXQ",
			false,
		},
		{
			"PDGR",
			"playfair example",
			"ABCX",
			false,
		},
		{
			"BMODZBXDNABEKUDMUIXMMOUVIF",
			"playfair example",
			"HIDETHEGOLDINTHETREXESTUMP",
			false,
		},
	}

	for idx, tc := range testCases {
		actual, err := pf.Decrypt(tc.input, tc.key)
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
