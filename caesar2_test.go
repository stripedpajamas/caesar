package caesar

import (
	"fmt"
	"testing"
)

func TestSomething(t *testing.T) {
	tb1 := newTranspositionBlock("CARGO")
	tb2 := newTranspositionBlock("CARGO")
	fmt.Println(tb1.transpose("abcdefg"))
	fmt.Println(tb2.detranspose("bgafdec"))
}
