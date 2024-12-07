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

	xMap := []obstacleDim{
		[]*obstacle{{pos: position{8, 0}}},
		[]*obstacle{{pos: position{6, 1}}},
		[]*obstacle{{pos: position{3, 2}}},
		[]*obstacle{},
		[]*obstacle{{pos: position{0, 4}}},
		[]*obstacle{},
		[]*obstacle{{pos: position{9, 6}}},
		[]*obstacle{{pos: position{4, 7}}},
		[]*obstacle{{pos: position{7, 8}}},
		[]*obstacle{{pos: position{1, 9}}},
	}

	yMap := []obstacleDim{
		[]*obstacle{{pos: position{0, 4}}},
		[]*obstacle{{pos: position{1, 9}}},
		[]*obstacle{},
		[]*obstacle{{pos: position{3, 2}}},
		[]*obstacle{{pos: position{4, 7}}},
		[]*obstacle{},
		[]*obstacle{{pos: position{6, 1}}},
		[]*obstacle{{pos: position{7, 8}}},
		[]*obstacle{{pos: position{8, 0}}},
		[]*obstacle{{pos: position{9, 6}}},
	}

	guardStart := guard{
		pos: position{y: 6, x: 4},
		dir: position{y: -1, x: 0},
	}

	for i, yEntry := range d.yMap {
		if len(yEntry) != len(yMap[i]) {
			t.Fatalf("LoadInput() -> len(yMap[%d]) %d, want %d", i, len(yEntry), len(yMap[i]))
		}

		if len(yEntry) == 0 {
			continue
		}

		if yEntry[0].pos.y != yMap[i][0].pos.y || yEntry[0].pos.x != yMap[i][0].pos.x {
			t.Fatalf("LoadInput() -> yMap[%d][0] position %v, want %v", i, yEntry[0].pos, yMap[i][0].pos)
		}
	}

	for i, xEntry := range d.xMap {
		if len(xEntry) != len(xMap[i]) {
			t.Fatalf("LoadInput() -> len(xMap[%d]) %d, want %d", i, len(xEntry), len(xMap[i]))
		}

		if len(xEntry) == 0 {
			continue
		}

		if xEntry[0].pos.y != xMap[i][0].pos.y || xEntry[0].pos.x != xMap[i][0].pos.x {
			t.Fatalf("LoadInput() -> xMap[%d][0] position %v, want %v", i, xEntry[0].pos, xMap[i][0].pos)
		}
	}

	if !reflect.DeepEqual(d.guardStart, guardStart) {
		t.Fatalf("LoadInput() -> guardStart %v, want %v", d.guardStart, guardStart)
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
