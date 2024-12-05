package day01

import (
	"reflect"
	"strings"
	"testing"
)

const example = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestLoadInput(t *testing.T) {
	d := Day01{}
	d.LoadInput(strings.NewReader(example))

	lefts := []int{3, 4, 2, 1, 3, 3}
	rights := []int{4, 3, 5, 3, 9, 3}

	if !reflect.DeepEqual(d.lefts, lefts) {
		t.Fatalf("LoadInput() -> lefts %v, want %v", d.lefts, lefts)
	}

	if !reflect.DeepEqual(d.rights, rights) {
		t.Fatalf("LoadInput() -> rights %v, want %v", d.rights, rights)
	}
}

func TestPartOne(t *testing.T) {
	d := Day01{}
	d.LoadInput(strings.NewReader(example))

	out := 11
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	d := Day01{}
	d.LoadInput(strings.NewReader(example))

	out := 31
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
