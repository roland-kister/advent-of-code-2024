// https://adventofcode.com/2024

package main

import (
	"fmt"
	"os"

	"github.com/roland-kister/advent-of-code-2024/day01"
	"github.com/roland-kister/advent-of-code-2024/day02"
	"github.com/roland-kister/advent-of-code-2024/day03"
	"github.com/roland-kister/advent-of-code-2024/day04"
	"github.com/roland-kister/advent-of-code-2024/day05"
	"github.com/roland-kister/advent-of-code-2024/day06"
	"github.com/roland-kister/advent-of-code-2024/day07"
	"github.com/roland-kister/advent-of-code-2024/day08"
	"github.com/roland-kister/advent-of-code-2024/day09"
	"github.com/roland-kister/advent-of-code-2024/internal"
)

func main() {
	solvers := []internal.Solver{
		&day01.Day01{},
		&day02.Day02{},
		&day03.Day03{},
		&day04.Day04{},
		&day05.Day05{},
		&day06.Day06{},
		&day07.Day07{},
		&day08.Day08{},
		&day09.Day09{},
	}

	for i, solver := range solvers {
		dayIndex := fmt.Sprintf("%02d", i+1)
		fmt.Printf("day %s:\n", dayIndex)

		input, err := os.Open(fmt.Sprintf("./inputs/%s.txt", dayIndex))
		if err != nil {
			panic(err)
		}

		solver.LoadInput(input)

		input.Close()

		partOne := solver.PartOne()
		fmt.Printf("\tpart one = %d\n", partOne)

		partTwo := solver.PartTwo()
		fmt.Printf("\tpart two = %d\n", partTwo)
	}
}
