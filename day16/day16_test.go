package day16

import (
	"strings"
	"testing"
)

const example = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

func TestLoadInput(t *testing.T) {
	d := Day16{}
	d.LoadInput(strings.NewReader(example))

	// if !reflect.DeepEqual(d.wh, wh) {
	// 	t.Fatalf("LoadInput() -> wh %v, want %v", d.wh, wh)
	// }
}

func TestPartOne(t *testing.T) {
	d := Day16{}
	d.LoadInput(strings.NewReader(example))

	out := 10092
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	d := Day16{}
	d.LoadInput(strings.NewReader(example))

	out := 9021
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
