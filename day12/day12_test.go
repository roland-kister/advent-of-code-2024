package day12

import (
	"reflect"
	"strings"
	"testing"
)

const example = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func TestLoadInput(t *testing.T) {
	d := Day12{}
	d.LoadInput(strings.NewReader(example))

	regionMap := farmMap{
		{'R', 'R', 'R', 'R', 'I', 'I', 'C', 'C', 'F', 'F'},
		{'R', 'R', 'R', 'R', 'I', 'I', 'C', 'C', 'C', 'F'},
		{'V', 'V', 'R', 'R', 'R', 'C', 'C', 'F', 'F', 'F'},
		{'V', 'V', 'R', 'C', 'C', 'C', 'J', 'F', 'F', 'F'},
		{'V', 'V', 'V', 'V', 'C', 'J', 'J', 'C', 'F', 'E'},
		{'V', 'V', 'I', 'V', 'C', 'C', 'J', 'J', 'E', 'E'},
		{'V', 'V', 'I', 'I', 'I', 'C', 'J', 'J', 'E', 'E'},
		{'M', 'I', 'I', 'I', 'I', 'I', 'J', 'J', 'E', 'E'},
		{'M', 'I', 'I', 'I', 'S', 'I', 'J', 'E', 'E', 'E'},
		{'M', 'M', 'M', 'I', 'S', 'S', 'J', 'E', 'E', 'E'},
	}

	if !reflect.DeepEqual(d.region, regionMap) {
		t.Fatalf("LoadInput() -> regionMap %v, want %v", d.region, regionMap)
	}
}

func TestPartOne(t *testing.T) {
	d := Day12{}
	d.LoadInput(strings.NewReader(example))

	out := 1930
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	d := Day12{}
	d.LoadInput(strings.NewReader(example))

	out := 1206
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
