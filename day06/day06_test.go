package day06

import (
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
