// https://adventofcode.com/2024/day/12

package day12

import (
	"bufio"
	"fmt"
	"io"
)

type farmMap [][]byte

type Day12 struct {
	fMap farmMap
}

func (d *Day12) LoadInput(input io.Reader) {
	d.fMap = make(farmMap, 0)

	scanner := bufio.NewScanner(input)

	for y := 0; scanner.Scan(); y++ {
		row := scanner.Bytes()

		d.fMap = append(d.fMap, row)
	}
}

func (d *Day12) PartOne() int {
	fMap := d.fMap.copy()

	for y, row := range fMap {
		for x, val := range row {
			if val == 0 {
				continue
			}

			fmt.Println(fMap[y][x])
		}
	}

	return 0
}

func (d *Day12) PartTwo() int {
	return 0
}

func (fMap farmMap) copy() farmMap {
	newMap := make(farmMap, len(fMap))

	for y, row := range fMap {
		newMap[y] = make([]byte, len(row))
		copy(newMap[y], row)
	}

	return newMap
}
