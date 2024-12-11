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

	stoneTwo := &stone{
		val:       17,
		nextStone: nil,
	}

	initStone := &stone{
		val:       125,
		nextStone: stoneTwo,
	}

	if !reflect.DeepEqual(d.initStone, initStone) {
		t.Fatalf("LoadInput() -> initStone %v, want %v", d.initStone, initStone)
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

	out := 55312
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
