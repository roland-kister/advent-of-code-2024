package internal

import "io"

type Solver interface {
	LoadInput(io.Reader)
	PartOne() int
	PartTwo() int
}
