package parser_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/kevinschoon/bankocr/pkg/parser"
)

func maybe(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func equal(first, second [9]int) bool {
	for i := 0; i < 8; i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

var sample = [][9]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{1, 1, 1, 1, 1, 1, 1, 1, 1},
	{2, 2, 2, 2, 2, 2, 2, 2, 2},
	{3, 3, 3, 3, 3, 3, 3, 3, 3},
	{4, 4, 4, 4, 4, 4, 4, 4, 4},
	{5, 5, 5, 5, 5, 5, 5, 5, 5},
	{6, 6, 6, 6, 6, 6, 6, 6, 6},
	{7, 7, 7, 7, 7, 7, 7, 7, 7},
	{8, 8, 8, 8, 8, 8, 8, 8, 8},
	{9, 9, 9, 9, 9, 9, 9, 9, 9},
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
}

func TestParser(t *testing.T) {
	fp, err := os.Open("../../examples/sample.txt")
	maybe(t, err)
	p := parser.New()
	results, err := p.ReadAll(fp)
	maybe(t, err)
	if ! equal(sample[0], results[0]) {
		maybe(t, fmt.Errorf("%T != %T", sample[0], results[1]))
	}
}
