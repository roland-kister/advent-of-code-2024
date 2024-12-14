// https://adventofcode.com/2024/day/12

package day12

import (
	"bufio"
	"io"
	"math"
)

type farmMap [][]byte

const oob = math.MaxUint8
const empty = 0
const (
	visited     byte = 0x1
	topFence    byte = 0x2
	rightFence  byte = 0x4
	bottomFence byte = 0x8
	leftFence   byte = 0x10
)

type Day12 struct {
	region farmMap
	buff   farmMap
}

func (d *Day12) LoadInput(input io.Reader) {
	d.region = make(farmMap, 0)

	scanner := bufio.NewScanner(input)

	for y := 0; scanner.Scan(); y++ {
		row := scanner.Bytes()

		d.region = append(d.region, make([]byte, len(row)))
		copy(d.region[y], row)
	}
}

func (d *Day12) PartOne() int {
	d.makeBuffer()
	return d.solve(true)

}

func (d *Day12) PartTwo() int {
	d.makeBuffer()
	return d.solve(false)
}

func (d *Day12) makeBuffer() {
	d.buff = make(farmMap, len(d.region))

	for y, row := range d.region {
		d.buff[y] = make([]byte, len(row))

		for x := range row {
			d.buff[y][x] = empty
		}
	}
}

func (d *Day12) solve(partOne bool) int {
	cost := 0

	for y := range len(d.buff) {
		for x := range len(d.buff[y]) {
			if d.buff[y][x]&visited != 0 {
				continue
			}

			area, perim := d.calculate(y, x, partOne)

			cost += area * perim
		}
	}

	return cost
}

func (d *Day12) calculate(y, x int, partOne bool) (area, perim int) {
	area, perim = 0, 0

	if d.buff.at(y, x) == oob || d.buff.at(y, x)&visited != 0 {
		return
	}

	d.buff[y][x] |= visited
	area = 1

	nextArea, nextPerim := d.nextArea(y, x, -1, 0, partOne)
	area, perim = area+nextArea, perim+nextPerim

	nextArea, nextPerim = d.nextArea(y, x, 0, 1, partOne)
	area, perim = area+nextArea, perim+nextPerim

	nextArea, nextPerim = d.nextArea(y, x, 1, 0, partOne)
	area, perim = area+nextArea, perim+nextPerim

	nextArea, nextPerim = d.nextArea(y, x, 0, -1, partOne)
	area, perim = area+nextArea, perim+nextPerim

	return
}

func (d *Day12) nextArea(y, x, yDir, xDir int, partOne bool) (area, perim int) {
	area, perim = 0, 0

	if d.region.at(y, x) == d.region.at(y+yDir, x+xDir) {
		nextArea, nextPerim := d.calculate(y+yDir, x+xDir, partOne)
		area += nextArea
		perim += nextPerim
	} else {
		if partOne || !d.handleFence(y, x, yDir, xDir) {
			perim++
		}
	}

	return
}

func (d *Day12) handleFence(y, x, yDir, xDir int) bool {
	ySide, xSide := xDir, yDir

	fenceMask := topFence
	if xDir == 1 {
		fenceMask = rightFence
	} else if yDir == 1 {
		fenceMask = bottomFence
	} else if xDir == -1 {
		fenceMask = leftFence
	}

	if d.buff[y][x]&fenceMask != 0 {
		return true
	}

	d.buff[y][x] |= fenceMask

	char := d.region.at(y, x)

	for yNext, xNext := y+ySide, x+xSide; ; yNext, xNext = yNext+ySide, xNext+xSide {
		charNext := d.region.at(yNext, xNext)
		if charNext != char {
			break
		}

		if charNext == d.region.at(yNext+yDir, xNext+xDir) {
			break
		}

		d.buff[yNext][xNext] |= fenceMask
	}

	for yNext, xNext := y+ySide*-1, x+xSide*-1; ; yNext, xNext = yNext+ySide*-1, xNext+xSide*-1 {
		charNext := d.region.at(yNext, xNext)
		if charNext != char {
			break
		}

		if charNext == d.region.at(yNext+yDir, xNext+xDir) {
			break
		}

		d.buff[yNext][xNext] |= fenceMask
	}

	return false
}

func (fMap farmMap) at(y, x int) byte {
	if y < 0 || y >= len(fMap) {
		return oob
	}

	if x < 0 || x >= len(fMap[y]) {
		return oob
	}

	return fMap[y][x]
}
