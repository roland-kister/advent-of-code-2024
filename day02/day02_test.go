package day02

import (
	"reflect"
	"strings"
	"testing"
)

const example = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestLoadInput(t *testing.T) {
	d := Day02{}
	d.LoadInput(strings.NewReader(example))

	reps := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	if !reflect.DeepEqual(d.reps, reps) {
		t.Fatalf("LoadInput() -> reps %v, want %v", d.reps, reps)
	}
}

func TestPartOne(t *testing.T) {
	d := Day02{}
	d.LoadInput(strings.NewReader(example))

	out := 2
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	d := Day02{}
	d.LoadInput(strings.NewReader(example))

	out := 4
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
