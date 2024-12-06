// https://adventofcode.com/2024/day/6

package day06

import (
	"bufio"
	"fmt"
	"io"
)

type tileType rune

const (
	empty        tileType = '.'
	obstacle     tileType = '#'
	newObstacle  tileType = 'O'
	guardUp      tileType = '^'
	guardRight   tileType = '>'
	guardLeft    tileType = '<'
	guardDown    tileType = 'v'
	visitedUp    tileType = 'U'
	visitedRight tileType = 'R'
	visitedDown  tileType = 'D'
	visitedLeft  tileType = 'L'
)

type Day06 struct {
	tiles       [][]tileType
	guardStartX int
	guardStartY int
}

func (d *Day06) LoadInput(input io.Reader) {
	d.tiles = make([][]tileType, 0)
	d.guardStartX = 0
	d.guardStartY = 0

	scanner := bufio.NewScanner(input)

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Bytes()

		d.tiles = append(d.tiles, make([]tileType, len(line)))

		for x, tileRaw := range line {
			tile := tileType(tileRaw)

			d.tiles[y][x] = tile

			switch tile {
			case guardUp, guardRight, guardDown, guardLeft:
				d.guardStartX = x
				d.guardStartY = y
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func (d *Day06) PartOne() int {
	guardX := d.guardStartX
	guardY := d.guardStartY

	if len(d.tiles) == 0 {
		return 0
	}

	for !d.done(guardX, guardY) {
		guardX, guardY = d.nextStep(guardX, guardY)
	}

	sum := 0

	for y := range d.tiles {
		for _, tile := range d.tiles[y] {
			if tile == visitedUp || tile == visitedRight || tile == visitedDown || tile == visitedLeft {
				sum++
			}
		}
	}

	return sum
}

func (d *Day06) done(guardX, guardY int) bool {
	return guardX < 0 || guardY < 0 || guardX >= len(d.tiles[0]) || guardY >= len(d.tiles)
}

func (d *Day06) nextStep(guardX, guardY int) (nextX, nextY int) {
	for range 4 {
		dirX := 0
		dirY := 0

		switch d.tiles[guardY][guardX] {
		case guardUp:
			dirY = -1
		case guardRight:
			dirX = 1
		case guardDown:
			dirY = 1
		case guardLeft:
			dirX = -1
		}

		if d.done(guardX+dirX, guardY+dirY) {
			d.tiles[guardY][guardX] = guardToVisited(d.tiles[guardY][guardX])
			return guardX + dirX, guardY + dirY
		}

		if d.tiles[guardY+dirY][guardX+dirX] != obstacle && d.tiles[guardY+dirY][guardX+dirX] != newObstacle {
			d.tiles[guardY+dirY][guardX+dirX] = d.tiles[guardY][guardX]
			d.tiles[guardY][guardX] = guardToVisited(d.tiles[guardY][guardX])
			return guardX + dirX, guardY + dirY
		}

		switch d.tiles[guardY][guardX] {
		case guardUp:
			d.tiles[guardY][guardX] = guardRight
		case guardRight:
			d.tiles[guardY][guardX] = guardDown
		case guardDown:
			d.tiles[guardY][guardX] = guardLeft
		case guardLeft:
			d.tiles[guardY][guardX] = guardUp
		}
	}

	panic("can't continue with guard")
}

func guardToVisited(guard tileType) tileType {
	switch guard {
	case guardUp:
		return visitedUp
	case guardRight:
		return visitedRight
	case guardDown:
		return visitedDown
	case guardLeft:
		return visitedLeft
	default:
		panic(fmt.Sprintf("can't convert guard %d to visited", guard))
	}
}

func (d *Day06) reset() {
	for y := range d.tiles {
		for x, tile := range d.tiles[y] {
			if tile == visitedUp || tile == visitedRight || tile == visitedDown || tile == visitedLeft {
				d.tiles[y][x] = empty
			}
		}
	}

	d.tiles[d.guardStartY][d.guardStartY] = guardUp
}

func (d *Day06) PartTwo() int {
	for y := range d.tiles {
		for x, tile := range d.tiles[y] {
			d.reset()

			if tile == empty {
				d.tiles[y][x] = empty
			}
		}
	}

	return 0
}
