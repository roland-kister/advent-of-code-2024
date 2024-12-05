package day03

import (
	"reflect"
	"strings"
	"testing"
)

const exampleOne = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
const exampleTwo = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestLoadInput(t *testing.T) {
	d := Day03{}
	d.LoadInput(strings.NewReader(strings.Join([]string{exampleOne, exampleTwo}, "\n")))

	insts := []string{
		"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
	}

	if !reflect.DeepEqual(d.insts, insts) {
		t.Fatalf("LoadInput() -> insts %v, want %v", d.insts, insts)
	}
}

func TestPartOne(t *testing.T) {
	d := Day03{}
	d.LoadInput(strings.NewReader(exampleOne))

	out := 161
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	d := Day03{}
	d.LoadInput(strings.NewReader(exampleTwo))

	out := 48
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
