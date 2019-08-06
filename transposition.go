package caesar

import "sort"

type transpositionBlock struct {
	data [][]rune
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
	tb.data = make([][]rune, (len(text)/len(key))+1)
	tb.seq = computeKeySeq(key)

	return tb
}

func computeKeySeq(key string) []int {
	seq := make([]int, len(key))
	elements := make([]element, len(key))
	for idx, r := range key {
		elements[idx] = element{value: r, index: idx}
	}
	sort.Sort(byRune(elements))
	for idx, el := range elements {
		seq[el.index] = idx
	}

	return seq
}
