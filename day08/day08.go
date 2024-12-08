// https://adventofcode.com/2024/day/8

package day08

import (
	"bufio"
	"io"
	"math"
)

type antenna struct {
	x    int
	y    int
	freq byte
}

type antinode struct {
	x int
	y int
}

type Day08 struct {
	ans  []antenna
	xMax int
	yMax int
}

func (d *Day08) LoadInput(input io.Reader) {
	d.ans = make([]antenna, 0)
	d.xMax = 0

	scanner := bufio.NewScanner(input)

	y := 0
	for ; scanner.Scan(); y++ {
		inLine := scanner.Bytes()

		if d.xMax == 0 {
			d.xMax = len(inLine) - 1
		}

		for x, freq := range inLine {
			if freq == '.' {
				continue
			}

			d.ans = append(d.ans, antenna{
				x,
				y,
				freq,
			})
		}
	}

	d.yMax = y - 1
}

func (d *Day08) PartOne() int {
	return d.getSum(1, 1)
}

func (d *Day08) PartTwo() int {
	return d.getSum(0, math.MaxInt)
}

func (d *Day08) anRoutine(an antenna, minIter, maxIter int, antisChan chan<- []antinode) {
	antis := make([]antinode, 0)

	for _, otherAn := range d.ans {
		if an.freq != otherAn.freq {
			continue
		}

		if an.x == otherAn.x && an.y == otherAn.y {
			continue
		}

		xDiff := an.x - otherAn.x
		yDiff := an.y - otherAn.y

		iter := minIter
		anti := (antinode{x: an.x + xDiff*iter, y: an.y + yDiff*iter})

		for ; !d.isAntiOut(anti) && iter <= maxIter; iter++ {
			antis = append(antis, anti)
			anti = (antinode{x: an.x + xDiff*iter, y: an.y + yDiff*iter})
		}
	}

	antisChan <- antis
}

func (d *Day08) isAntiOut(anti antinode) bool {
	return (anti.x > d.xMax || anti.x < 0 || anti.y > d.yMax || anti.y < 0)
}

func (d *Day08) getSum(minIter, maxIter int) int {
	antisChan := make(chan []antinode)

	for _, an := range d.ans {
		go d.anRoutine(an, minIter, maxIter, antisChan)

	}

	antis := make([]antinode, 0)
	for range len(d.ans) {
		antis = append(antis, <-antisChan...)
	}

	antisMap := make(map[antinode]bool)
	for _, anti := range antis {
		_, ok := antisMap[anti]
		if ok {
			continue
		}

		antisMap[anti] = true
	}

	return len(antisMap)
}
