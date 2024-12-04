package day04

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	grid := [][]rune{
		{'M','M','M','S','X','X','M','A','S','M'},
		{'M','S','A','M','X','M','S','M','S','A'},
		{'A','M','X','S','X','M','A','A','M','M'},
		{'M','S','A','M','A','S','M','S','M','X'},
		{'X','M','A','S','A','M','X','A','M','M'},
		{'X','X','A','M','M','X','X','A','M','A'},
		{'S','M','S','M','S','A','S','X','S','S'},
		{'S','A','X','A','M','A','S','A','A','A'},
		{'M','A','M','M','M','X','M','M','M','M'},
		{'M','X','M','X','A','X','M','A','S','X'},
	}

	out := 18

	d := day04{
		grid,
	}

	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	grid := [][]rune{
		{'M','M','M','S','X','X','M','A','S','M'},
		{'M','S','A','M','X','M','S','M','S','A'},
		{'A','M','X','S','X','M','A','A','M','M'},
		{'M','S','A','M','A','S','M','S','M','X'},
		{'X','M','A','S','A','M','X','A','M','M'},
		{'X','X','A','M','M','X','X','A','M','A'},
		{'S','M','S','M','S','A','S','X','S','S'},
		{'S','A','X','A','M','A','S','A','A','A'},
		{'M','A','M','M','M','X','M','M','M','M'},
		{'M','X','M','X','A','X','M','A','S','X'},
	}

	out := 9

	d := day04{
		grid,
	}

	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
