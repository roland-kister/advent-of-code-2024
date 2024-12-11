package day11

import (
	"reflect"
	"strings"
	"testing"
)

const example = `125 17`

func TestLoadInput(t *testing.T) {
	d := Day11{}
	d.LoadInput(strings.NewReader(example))

	stones := stoneMap{
		125: 1,
		17:  1,
	}

	if !reflect.DeepEqual(d.stones, stones) {
		t.Fatalf("LoadInput() -> stones %v, want %v", d.stones, stones)
	}
}

func TestPartOne(t *testing.T) {
	d := Day11{}
	d.LoadInput(strings.NewReader(example))

	out := 55312
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	d := Day11{}
	d.LoadInput(strings.NewReader(example))

	out := 65601038650482
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
