package day14

import (
	"reflect"
	"strings"
	"testing"
)

const example = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

func TestLoadInput(t *testing.T) {
	d := Day14{}
	d.LoadInput(strings.NewReader(example))

	robs := []robot{
		{[2]int{0, 4}, [2]int{3, -3}},
		{[2]int{6, 3}, [2]int{-1, -3}},
		{[2]int{10, 3}, [2]int{-1, 2}},
		{[2]int{2, 0}, [2]int{2, -1}},
		{[2]int{0, 0}, [2]int{1, 3}},
		{[2]int{3, 0}, [2]int{-2, -2}},
		{[2]int{7, 6}, [2]int{-1, -3}},
		{[2]int{3, 0}, [2]int{-1, -2}},
		{[2]int{9, 3}, [2]int{2, 3}},
		{[2]int{7, 3}, [2]int{-1, 2}},
		{[2]int{2, 4}, [2]int{2, -3}},
		{[2]int{9, 5}, [2]int{-3, -3}},
	}

	xMax := 11
	yMax := 7

	if !reflect.DeepEqual(d.robs, robs) {
		t.Fatalf("LoadInput() -> robs %v, want %v", d.robs, robs)
	}

	if d.xMax != xMax {
		t.Fatalf("LoadInput() -> xMax %d, want %d", d.xMax, xMax)
	}

	if d.yMax != yMax {
		t.Fatalf("LoadInput() -> yMax %d, want %d", yMax, yMax)
	}
}

func TestPartOne(t *testing.T) {
	d := Day14{}
	d.LoadInput(strings.NewReader(example))

	out := 12
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	d := Day14{}
	d.LoadInput(strings.NewReader(example))

	out := 0
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
