// https://adventofcode.com/2024/day/10

package day10

import (
	"bufio"
	"io"
)

type Day10 struct {
	topoMap [][]byte
}

func (d *Day10) LoadInput(input io.Reader) {
	scanner := bufio.NewScanner(input)

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Bytes()

		if y == 0 {
			d.topoMap = make([][]byte, len(line))

			for x := range len(line) {
				d.topoMap[x] = make([]byte, 0)
			}
		}

		for x, value := range line {
			d.topoMap[x] = append(d.topoMap[x], value-'0')
		}
	}
}

func (d *Day10) PartOne() int {
	sum := 0

	for x := range d.topoMap {
		for y, value := range d.topoMap[x] {
			if value != 0 {
				continue
			}

			peaks := d.traverse(x, y)

			uniquePeaks := make(map[[2]int]bool)
			for _, peak := range peaks {
				uniquePeaks[peak] = true
			}

			sum += len(uniquePeaks)
		}
	}

	return sum
}

func (d *Day10) PartTwo() int {
	sum := 0

	for x := range d.topoMap {
		for y, value := range d.topoMap[x] {
			if value != 0 {
				continue
			}

			peaks := d.traverse(x, y)

			sum += len(peaks)
		}
	}

	return sum
}

func (d *Day10) traverse(x, y int) [][2]int {
	if d.topoMap[x][y] == 9 {
		return [][2]int{{x, y}}
	}

	res := make([][2]int, 0)

	next := d.topoMap[x][y] + 1

	if y > 0 && d.topoMap[x][y-1] == next {
		res = append(res, d.traverse(x, y-1)...)
	}

	if x < len(d.topoMap)-1 && d.topoMap[x+1][y] == next {
		res = append(res, d.traverse(x+1, y)...)
	}

	if y < len(d.topoMap[x])-1 && d.topoMap[x][y+1] == next {
		res = append(res, d.traverse(x, y+1)...)
	}

	if x > 0 && d.topoMap[x-1][y] == next {
		res = append(res, d.traverse(x-1, y)...)
	}

	return res
}
