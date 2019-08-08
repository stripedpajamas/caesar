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

type colinfo struct {
	size  int
	idx   int
	value int
}

type byRune []element

func (a byRune) Len() int           { return len(a) }
func (a byRune) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byRune) Less(i, j int) bool { return a[i].value < a[j].value }

type byValue []colinfo

func (a byValue) Len() int           { return len(a) }
func (a byValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byValue) Less(i, j int) bool { return a[i].value < a[j].value }

func newTranspositionBlock(key string) *transpositionBlock {
	tb := new(transpositionBlock)
	tb.data = make([]string, len(key))
	tb.seq = computeKeySeq(key)

	return tb
}

// loads text into block and then returns text as if it were read
// by column in alphabetic order of the original keyword
func (tb *transpositionBlock) transpose(text string) string {
	for i := 0; i < len(tb.seq); i++ {
		var column strings.Builder
		for j := i; j < len(text); j += len(tb.seq) {
			column.WriteByte(text[j])
		}
		tb.data[i] = column.String()
	}

	var out strings.Builder
	for _, colIdx := range tb.seq {
		out.WriteString(tb.data[colIdx])
	}
	return out.String()
}

// loads already transformed text into the block and returns
// text as if it were read by column in original order of key letters
func (tb *transpositionBlock) detranspose(text string) string {
	chunk := (len(text) / len(tb.seq)) + 1
	incompleteColumns := (chunk * len(tb.seq)) - len(text)
	fullColumns := len(tb.seq) - incompleteColumns

	// associate seq with size of destination columns
	colSizes := make([]int, len(tb.seq))
	for i, seq := range tb.seq {
		if seq < fullColumns {
			colSizes[i] = chunk
		} else {
			colSizes[i] = chunk - 1
		}
	}

	// iterate through text, adding letters to appropriate columns
	// based on sequence number
	textIdx := 0
	for i, colIdx := range tb.seq {
		var column strings.Builder
		start, end := textIdx, textIdx+colSizes[i]
		textIdx = end
		for start < end {
			column.WriteByte(text[start])
			start++
		}
		tb.data[colIdx] = column.String()
	}
	// return value is data read off by row
	var out strings.Builder
	for row := 0; row < chunk; row++ {
		for _, col := range tb.data {
			if len(col) <= row {
				continue
			}
			out.WriteByte(col[row])
		}
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
