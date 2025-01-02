package day22

import (
	"strings"
	"testing"
)

const example = `1
2
3
2024`

func TestPartOne(t *testing.T) {
	d := Day22{}
	d.LoadInput(strings.NewReader(example))

	out := 37327623
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	d := Day22{}
	d.LoadInput(strings.NewReader(example))

	out := 23
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
