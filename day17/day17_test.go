package day17

import (
	"strings"
	"testing"
)

const example = `Register A: 729
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

func TestPartOne(t *testing.T) {
	d := Day17{}
	d.LoadInput(strings.NewReader(example))

	out := 3310
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

//func TestPartTwo(t *testing.T) {
//	d := Day17{}
//	d.LoadInput(strings.NewReader(example))
//
//	out := 117440
//	actOut := d.PartTwo()
//
//	if out != actOut {
//		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
//	}
//}
