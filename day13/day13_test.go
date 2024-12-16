package day13

import (
	"reflect"
	"strings"
	"testing"
)

const example = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

func TestLoadInput(t *testing.T) {
	d := Day13{}
	d.LoadInput(strings.NewReader(example))

	partOneMachs := []machine{
		{a: [2]int{94, 34}, b: [2]int{22, 67}, prize: [2]int{8400, 5400}},
		{a: [2]int{26, 66}, b: [2]int{67, 21}, prize: [2]int{12748, 12176}},
		{a: [2]int{17, 86}, b: [2]int{84, 37}, prize: [2]int{7870, 6450}},
		{a: [2]int{69, 23}, b: [2]int{27, 71}, prize: [2]int{18641, 10279}},
	}

	partTwoMachs := []machine{
		{a: [2]int{94, 34}, b: [2]int{22, 67}, prize: [2]int{8400 + p2Mul, 5400 + p2Mul}},
		{a: [2]int{26, 66}, b: [2]int{67, 21}, prize: [2]int{12748 + p2Mul, 12176 + p2Mul}},
		{a: [2]int{17, 86}, b: [2]int{84, 37}, prize: [2]int{7870 + p2Mul, 6450 + p2Mul}},
		{a: [2]int{69, 23}, b: [2]int{27, 71}, prize: [2]int{18641 + p2Mul, 10279 + p2Mul}},
	}

	if !reflect.DeepEqual(d.partOneMachs, partOneMachs) {
		t.Fatalf("LoadInput() -> partOneMachs %v, want %v", d.partOneMachs, partOneMachs)
	}

	if !reflect.DeepEqual(d.partTwoMachs, partTwoMachs) {
		t.Fatalf("LoadInput() -> partTwoMachs %v, want %v", d.partTwoMachs, partTwoMachs)
	}
}

func TestPartOne(t *testing.T) {
	d := Day13{}
	d.LoadInput(strings.NewReader(example))

	out := 480
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	d := Day13{}
	d.LoadInput(strings.NewReader(example))

	out := 875318608908
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
