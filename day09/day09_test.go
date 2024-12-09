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

	diskMap := []uint{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2}
	fragments := []fragment{
		{id: 0, length: 2},
		{id: empty, length: 3},
		{id: 1, length: 3},
		{id: empty, length: 3},
		{id: 2, length: 1},
		{id: empty, length: 3},
		{id: 3, length: 3},
		{id: empty, length: 1},
		{id: 4, length: 2},
		{id: empty, length: 1},
		{id: 5, length: 4},
		{id: empty, length: 1},
		{id: 6, length: 4},
		{id: empty, length: 1},
		{id: 7, length: 3},
		{id: empty, length: 1},
		{id: 8, length: 4},
		{id: empty, length: 0},
		{id: 9, length: 2},
	}

	if !reflect.DeepEqual(d.diskMap, diskMap) {
		t.Fatalf("LoadInput() -> diskMap %v, want %v", d.diskMap, diskMap)
	}

	if !reflect.DeepEqual(d.fragments, fragments) {
		t.Fatalf("LoadInput() -> fragments %v, want %v", d.fragments, fragments)
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
