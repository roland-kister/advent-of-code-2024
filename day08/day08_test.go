package day08

import (
	"reflect"
	"strings"
	"testing"
)

const example = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestLoadInput(t *testing.T) {
	d := Day08{}
	d.LoadInput(strings.NewReader(example))

	ans := []antenna{
		{x: 8, y: 1, freq: '0'},
		{x: 5, y: 2, freq: '0'},
		{x: 7, y: 3, freq: '0'},
		{x: 4, y: 4, freq: '0'},
		{x: 6, y: 5, freq: 'A'},
		{x: 8, y: 8, freq: 'A'},
		{x: 9, y: 9, freq: 'A'},
	}

	xMax := 11
	yMax := 11

	if !reflect.DeepEqual(d.ans, ans) {
		t.Fatalf("LoadInput() -> ans %v, want %v", d.ans, ans)
	}

	if d.xMax != xMax {
		t.Fatalf("LoadInput() -> ans %d, want %d", xMax, xMax)
	}

	if d.yMax != yMax {
		t.Fatalf("LoadInput() -> ans %d, want %d", yMax, yMax)
	}
}

func TestPartOne(t *testing.T) {
	d := Day08{}
	d.LoadInput(strings.NewReader(example))

	out := 14
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	d := Day08{}
	d.LoadInput(strings.NewReader(example))

	out := 34
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
