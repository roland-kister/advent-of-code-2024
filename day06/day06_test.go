package day06

import (
	"reflect"
	"strings"
	"testing"
)

const example = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

func TestLoadInput(t *testing.T) {
	d := Day06{}
	d.LoadInput(strings.NewReader(example))

	tiles := [][]tileType{
		{empty, empty, empty, empty, obstacle, empty, empty, empty, empty, empty},
		{empty, empty, empty, empty, empty, empty, empty, empty, empty, obstacle},
		{empty, empty, empty, empty, empty, empty, empty, empty, empty, empty},
		{empty, empty, obstacle, empty, empty, empty, empty, empty, empty, empty},
		{empty, empty, empty, empty, empty, empty, empty, obstacle, empty, empty},
		{empty, empty, empty, empty, empty, empty, empty, empty, empty, empty},
		{empty, obstacle, empty, empty, guardUp, empty, empty, empty, empty, empty},
		{empty, empty, empty, empty, empty, empty, empty, empty, obstacle, empty},
		{obstacle, empty, empty, empty, empty, empty, empty, empty, empty, empty},
		{empty, empty, empty, empty, empty, empty, obstacle, empty, empty, empty},
	}

	guardStartX := 4
	guardStartY := 6

	if !reflect.DeepEqual(d.tiles, tiles) {
		t.Fatalf("LoadInput() -> tiles %v, want %v", d.tiles, tiles)
	}

	if d.guardStartX != guardStartX {
		t.Fatalf("LoadInput() -> guardStartX %v, want %v", d.guardStartX, guardStartX)
	}

	if d.guardStartY != guardStartY {
		t.Fatalf("LoadInput() -> guardStartY %v, want %v", d.guardStartY, guardStartY)
	}
}

func TestPartOne(t *testing.T) {
	d := Day06{}
	d.LoadInput(strings.NewReader(example))

	out := 41
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	d := Day06{}
	d.LoadInput(strings.NewReader(example))

	out := 6
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
