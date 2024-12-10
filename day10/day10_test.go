package day10

import (
	"reflect"
	"strings"
	"testing"
)

const example = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestLoadInput(t *testing.T) {
	d := Day10{}
	d.LoadInput(strings.NewReader(example))

	topoMap := [][]byte{
		{8, 7, 8, 9, 4, 3, 0, 1},
		{9, 8, 7, 6, 5, 2, 1, 0},
		{0, 1, 4, 5, 6, 0, 3, 4},
		{1, 2, 3, 4, 7, 1, 2, 5},
		{0, 1, 0, 9, 8, 9, 9, 6},
		{1, 8, 9, 8, 9, 0, 8, 7},
		{2, 7, 6, 7, 0, 1, 0, 3},
		{3, 4, 5, 4, 3, 2, 1, 2},
	}

	if !reflect.DeepEqual(d.topoMap, topoMap) {
		t.Fatalf("LoadInput() -> topoMap %v, want %v", d.topoMap, topoMap)
	}
}

func TestPartOne(t *testing.T) {
	d := Day10{}
	d.LoadInput(strings.NewReader(example))

	out := 36
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	d := Day10{}
	d.LoadInput(strings.NewReader(example))

	out := 81
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
