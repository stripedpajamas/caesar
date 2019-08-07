package caesar

import (
	"sort"
	"strings"
)

type transpositionBlock struct {
	data []string
	seq  []int
}

type element struct {
	value rune
	index int
}

type byRune []element

func (a byRune) Len() int           { return len(a) }
func (a byRune) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byRune) Less(i, j int) bool { return a[i].value < a[j].value }

func newTranspositionBlock(text, key string) *transpositionBlock {
	tb := new(transpositionBlock)
	tb.data = make([]string, len(key))
	tb.seq = computeKeySeq(key)

	chunk := (len(text) / len(key)) + 1

	for i := 0; i < len(key); i++ {
		var column strings.Builder
		for j := i; j < len(text); j += chunk {
			column.WriteByte(text[j])
		}
		tb.data[i] = column.String()
	}

	return tb
}

func (tb *transpositionBlock) columnsInorder() string {
	var out strings.Builder
	for _, colIdx := range tb.seq {
		out.WriteString(tb.data[colIdx])
	}
	return out.String()
}

func computeKeySeq(key string) []int {
	seq := make([]int, len(key))
	elements := make([]element, len(key))
	for idx, r := range key {
		elements[idx] = element{value: r, index: idx}
	}
	sort.Sort(byRune(elements))
	for idx, el := range elements {
		seq[idx] = el.index
	}

	return seq
}
