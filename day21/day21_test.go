package day21

import (
	"strings"
	"testing"
)

const example = `029A
980A
179A
456A
379A`

func TestPartOne(t *testing.T) {
	d := Day21{}
	d.LoadInput(strings.NewReader(example))

	out := 126384
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}
