package day01

import "testing"

func TestPartOne(t *testing.T) {
	lefts := []int{3, 4, 2, 1, 3, 3}
	rights := []int{4, 3, 5, 3, 9, 3}

	out := 11

	d := day01{
		lefts,
		rights,
	}

	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	lefts := []int{3, 4, 2, 1, 3, 3}
	rights := []int{4, 3, 5, 3, 9, 3}

	out := 31

	d := day01{
		lefts,
		rights,
	}

	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
