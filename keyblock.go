package caesar

import "github.com/stripedpajamas/caesar/runes"

type keyblock struct {
	block  [5][5]rune
	lookup map[rune]location
}

type location struct {
	row int
	col int
}

func newKeyblock(key string) *keyblock {
	kb := new(keyblock)
	kb.block = [5][5]rune{}
	kb.lookup = make(map[rune]location)

	// get unique letters from key
	seen := make(map[rune]bool)
	uniqueKey, keyIdx := [25]rune{}, 0
	for _, r := range key {
		if !runes.IsLetter(r) || r == 'j' || r == 'J' {
			continue
		}
		ur := runes.ToUpper(r)
		if seen[ur] {
			continue
		}
		uniqueKey[keyIdx] = ur
		seen[ur] = true
		keyIdx++
	}

	// add the rest of the alphabet (excepting already used letters)
	for _, r := range alphabet {
		if seen[r] {
			continue
		}
		uniqueKey[keyIdx] = r
		keyIdx++
	}

	keyIdx = 0
	// write the resulting key into the keyblock
	for row := 0; row < 5; row++ {
		kb.block[row] = [5]rune{}
		for col := 0; col < 5; col++ {
			kb.block[row][col] = uniqueKey[keyIdx]
			kb.lookup[uniqueKey[keyIdx]] = location{row, col}
			keyIdx++
		}
	}

	return kb
}

// Used by Playfair cipher
// - If the letters appear on the same row of your table,
//   replace them with the letters to their immediate right respectively
// - If the letters appear on the same column of your table,
//   replace them with the letters immediately below respectively
// - If the letters are not on the same row or column,
//   replace them with the letters on the same row respectively
//   but at the other pair of corners of the rectangle defined by the original pair.
//   The order is important – the first letter of the encrypted pair is the
//   one that lies on the same row as the first letter of the plaintext pair.
func (kb *keyblock) getCorrespondingPair(a, b rune, reverse bool) (x, y rune) {
	var diff int
	if reverse {
		diff = -1
	} else {
		diff = 1
	}
	aLoc, bLoc := kb.lookup[a], kb.lookup[b]
	var xLoc, yLoc location
	if aLoc.row == bLoc.row {
		xLoc = location{row: aLoc.row, col: (aLoc.col + diff + 5) % 5}
		yLoc = location{row: bLoc.row, col: (bLoc.col + diff + 5) % 5}
	} else if aLoc.col == bLoc.col {
		xLoc = location{row: (aLoc.row + diff + 5) % 5, col: aLoc.col}
		yLoc = location{row: (bLoc.row + diff + 5) % 5, col: bLoc.col}
	} else {
		xLoc = location{row: aLoc.row, col: bLoc.col}
		yLoc = location{row: bLoc.row, col: aLoc.col}
	}
	x = kb.block[xLoc.row][xLoc.col]
	y = kb.block[yLoc.row][yLoc.col]
	return
}
