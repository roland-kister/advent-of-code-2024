package internal

type Solver interface {
	LoadInput(inputPath string)
	PartOne() int
	PartTwo() int
}
