package caesar_test

import (
	"fmt"
	"testing"

	"github.com/stripedpajamas/caesar"
)

type testCase struct {
	input       string
	key         string
	expected    string
	expectedErr bool
}

func TestCaesar(t *testing.T) {
	encryptionCases := []testCase{
		{"abc", "1", "BCD", false},
		{"zzz", "1", "AAA", false},
		{"aZf", "1", "BAG", false},
		{"123", "1", "", false},
		{"abc", "26", "ABC", false},
		{"abc", "27", "BCD", false},
		{"abc", "$", "", true},
		{"abc", "", "", true},
	}
	decryptionCases := []testCase{
		{"bcd", "1", "ABC", false},
		{"zzz", "1", "YYY", false},
		{"aZf", "1", "ZYE", false},
		{"123", "1", "", false},
		{"abc", "26", "ABC", false},
		{"abc", "27", "ZAB", false},
		{"abc", "$", "", true},
		{"abc", "", "", true},
	}
	runTests(t, caesar.Caesar{}, encryptionCases, decryptionCases)
}

func TestPlayfair(t *testing.T) {
	encryptionCases := []testCase{
		{"abx", "playfair example", "PDGW", false},
		{"abc", "playfair example", "PDGR", false},
		{
			"Hide the gold in the tree stump",
			"playfair example",
			"BMODZBXDNABEKUDMUIXMMOUVIF",
			false,
		},
	}

	decryptionCases := []testCase{
		{"PDGW", "playfair example", "ABXQ", false},
		{"PDGR", "playfair example", "ABCX", false},
		{
			"BMODZBXDNABEKUDMUIXMMOUVIF",
			"playfair example",
			"HIDETHEGOLDINTHETREXESTUMP",
			false,
		},
		{
			"BMODZBXDNABEKUDMUIXMMOUVI", // odd length ciphertext
			"playfair example",
			"",
			true,
		},
	}

	runTests(t, caesar.Playfair{}, encryptionCases, decryptionCases)
}

func TestVigenere(t *testing.T) {
	encryptionCases := []testCase{
		{"attack at dawn", "lemon", "LXFOPVEFRNHR", false},
		{"attack at dawn!!!", "lemon", "LXFOPVEFRNHR", false},
		{"CRYPTO IS SHORT FOR CRYPTOGRAPHY", "ABCD", "CSASTPKVSIQUTGQUCSASTPIUAQJB", false},
		{"asdf", "abc4", "", true},
	}

	decryptionCases := []testCase{
		{"LXFOPVEFRNHR", "lemon", "ATTACKATDAWN", false},
		{"CSASTPKVSIQUTGQUCSASTPIUAQJB", "ABCD", "CRYPTOISSHORTFORCRYPTOGRAPHY", false},
		{"asdf", "abc4", "", true},
	}

	runTests(t, caesar.Vigenere{}, encryptionCases, decryptionCases)
}

func TestADFGX(t *testing.T) {
	encryptionCases := []testCase{
		{"attack at once", "btalpdhozkqfvsngicuxmrewy,cargo", "FAXDFADDDGDGFFFAFAXAFAFX", false},
		{"attack at once!!!", "btalpdhozkqfvsngicuxmrewy,cargo", "FAXDFADDDGDGFFFAFAXAFAFX", false},
		{"hello world", "apple,book", "DAFFAGFDDDXFXXFAAXGD", false},
		{"asdf", "onlyonekey", "", true},
		{"asdf", "onlyonekey,", "", true},
		{"asdf", "good,666", "", true},
	}

	decryptionCases := []testCase{
		{"FAXDFADDDGDGFFFAFAXAFAFX", "btalpdhozkqfvsngicuxmrewy,cargo", "ATTACKATONCE", false},
		{"DAFFAGFDDDXFXXFAAXGD", "apple,book", "HELLOWORLD", false},
		{"GAAFA FXGDF GGAFG FGAGA GG", "help,me", "THISISATEST", false},
		{"asdf", "onlyonekey", "", true},
		{"asdf", "onlyonekey;", "", true},
		{"asdf", "good,666", "", true},
	}

	runTests(t, caesar.ADFGX{}, encryptionCases, decryptionCases)
}

func TestBifid(t *testing.T) {
	encryptionCases := []testCase{
		{"FLEE AT ONCE", "BGWKZQPNDSIOAXEFCLUMTHYVR", "UAEOLWRINS", false},
		{"flEE aT once!!!", "BGWKZQPNDSIOAXEFCLUMTHYVR", "UAEOLWRINS", false},
	}

	decryptionCases := []testCase{
		{"UAEOLWRINS", "BGWKZQPNDSIOAXEFCLUMTHYVR", "FLEEATONCE", false},
	}

	runTests(t, caesar.Bifid{}, encryptionCases, decryptionCases)
}

func runTests(t *testing.T, c caesar.Cipher, encryptionCases, decryptionCases []testCase) {
	for idx, tc := range encryptionCases {
		actual, err := c.Encrypt(tc.input, tc.key)
		if tc.expectedErr && err == nil {
			fmt.Printf("(encrypt) test %d failed: wanted error, got success\n", idx)
			t.Fail()
		}
		if !tc.expectedErr && err != nil {
			fmt.Printf("(encrypt) test %d failed: %s\n", idx, err.Error())
		}
		if actual != tc.expected {
			fmt.Printf("(encrypt) test %d failed: wanted %s, got %s\n", idx, tc.expected, actual)
			t.Fail()
		}
	}

	for idx, tc := range decryptionCases {
		actual, err := c.Decrypt(tc.input, tc.key)
		if tc.expectedErr && err == nil {
			fmt.Printf("(decrypt) test %d failed: wanted error, got success\n", idx)
			t.Fail()
		}
		if !tc.expectedErr && err != nil {
			fmt.Printf("(decrypt) test %d failed: %s\n", idx, err.Error())
		}
		if actual != tc.expected {
			fmt.Printf("(decrypt) test %d failed: wanted %s, got %s\n", idx, tc.expected, actual)
			t.Fail()
		}
	}
}
