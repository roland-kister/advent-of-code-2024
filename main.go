// https://adventofcode.com/2024

package main

import (
	"fmt"

	"github.com/roland-kister/advent-of-code-2024/day01"
	"github.com/roland-kister/advent-of-code-2024/day02"
	"github.com/roland-kister/advent-of-code-2024/internal"
)

func main() {
	solvers := []internal.Solver{
		day01.NewDay01(),
		day02.NewDay02(),
	}

	for i, solver := range solvers {
		dayIndex := fmt.Sprintf("%02d", i+1)

		fmt.Printf("day %s:\n", dayIndex)

		solver.LoadInput(fmt.Sprintf("./inputs/%s.txt", dayIndex))

		partOne := solver.PartOne()
		fmt.Printf("\t part one = %d\n", partOne)

		partTwo := solver.PartTwo()
		fmt.Printf("\t part one = %d\n", partTwo)
	}
}
