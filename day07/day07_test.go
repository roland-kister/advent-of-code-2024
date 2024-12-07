package day07

import (
	"reflect"
	"strings"
	"testing"
)

const example = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`

func TestLoadInput(t *testing.T) {
	d := Day07{}
	d.LoadInput(strings.NewReader(example))

	res := []uint64{190, 3267, 83, 156, 7290, 161011, 192, 21037, 292}
	parts := [][]uint64{
		{10, 19},
		{81, 40, 27},
		{17, 5},
		{15, 6},
		{6, 8, 6, 15},
		{16, 10, 13},
		{17, 8, 14},
		{9, 7, 18, 13},
		{11, 6, 16, 20},
	}

	if !reflect.DeepEqual(d.res, res) {
		t.Fatalf("LoadInput() -> res %v, want %v", d.res, res)
	}

	if !reflect.DeepEqual(d.parts, parts) {
		t.Fatalf("LoadInput() -> res %v, want %v", d.parts, parts)
	}
}

func TestPartOne(t *testing.T) {
	d := Day07{}
	d.LoadInput(strings.NewReader(example))

	out := 3749
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	d := Day07{}
	d.LoadInput(strings.NewReader(example))

	out := 11387
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
