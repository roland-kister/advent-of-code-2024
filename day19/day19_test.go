package day19

import (
	"strings"
	"testing"
)

const example = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

func TestPartOne(t *testing.T) {
	d := Day19{}
	d.LoadInput(strings.NewReader(example))

	out := 6
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	d := Day19{}
	d.LoadInput(strings.NewReader(example))

	out := 16
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
