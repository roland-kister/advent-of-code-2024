package day09

import (
	"reflect"
	"strings"
	"testing"
)

const example = `2333133121414131402`

func TestLoadInput(t *testing.T) {
	d := Day09{}
	d.LoadInput(strings.NewReader(example))

	disk := diskContent{0, 0, empty, empty, empty, 1, 1, 1, empty, empty, empty, 2, empty, empty, empty, 3, 3, 3, empty, 4, 4, empty, 5, 5, 5, 5, empty, 6, 6, 6, 6, empty, 7, 7, 7, empty, 8, 8, 8, 8, 9, 9}

	if !reflect.DeepEqual(d.disk, disk) {
		t.Fatalf("LoadInput() -> disk %v, want %v", d.disk, disk)
	}
}

func TestPartOne(t *testing.T) {
	d := Day09{}
	d.LoadInput(strings.NewReader(example))

	out := 1928
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	d := Day09{}
	d.LoadInput(strings.NewReader(example))

	out := 2858
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
