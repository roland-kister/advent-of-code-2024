// https://adventofcode.com/2024/day/6

// extremely horrendous solution ... please don't judge

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
	tiles         [][]tileType
	guardStartX   int
	guardStartY   int
	guardStartDir tileType
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
				d.guardStartDir = tile
			}
		}
	}

	d.tiles[d.guardStartY][d.guardStartY] = empty

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func (d *Day06) PartOne() int {
	guardX := d.guardStartX
	guardY := d.guardStartY
	guardDir := d.guardStartDir

	if len(d.tiles) == 0 {
		return 0
	}

	d.reset()

	for !d.done(guardX, guardY) {
		guardX, guardY, guardDir = d.nextStep(guardX, guardY, guardDir)
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

func (d *Day06) nextStep(guardX, guardY int, guardDir tileType) (nextX, nextY int, nextDir tileType) {
	newGuardDir := guardDir

	for range 4 {
		dirX := 0
		dirY := 0

		switch newGuardDir {
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
			d.tiles[guardY][guardX] = guardToVisited(newGuardDir)
			return guardX + dirX, guardY + dirY, newGuardDir
		}

		if d.tiles[guardY+dirY][guardX+dirX] != obstacle && d.tiles[guardY+dirY][guardX+dirX] != newObstacle {
			d.tiles[guardY][guardX] = guardToVisited(newGuardDir)
			return guardX + dirX, guardY + dirY, newGuardDir
		}

		switch newGuardDir {
		case guardUp:
			newGuardDir = guardRight
		case guardRight:
			newGuardDir = guardDown
		case guardDown:
			newGuardDir = guardLeft
		case guardLeft:
			newGuardDir = guardUp
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
			if tile == visitedUp || tile == visitedRight || tile == visitedDown || tile == visitedLeft || tile == newObstacle {
				d.tiles[y][x] = empty
			}
		}
	}

	d.tiles[d.guardStartY][d.guardStartY] = empty
}

func (d *Day06) PartTwo() int {
	sum := 0

	d.reset()

	for y := range d.tiles {
		for x, tile := range d.tiles[y] {
			shadowGrid := make([][][]bool, 4)
			for i := range 4 {
				shadowGrid[i] = make([][]bool, len(d.tiles))
				for y := range len(d.tiles) {
					shadowGrid[i][y] = make([]bool, len(d.tiles[y]))
				}

			}

			if tile == empty {
				d.tiles[y][x] = newObstacle
			} else {
				continue
			}

			guardX := d.guardStartX
			guardY := d.guardStartY
			guardDir := d.guardStartDir

		Main:
			for !d.done(guardX, guardY) {
				switch guardDir {
				case guardUp:
					if shadowGrid[0][guardY][guardX] {
						sum++
						break Main
					}
					shadowGrid[0][guardY][guardX] = true
				case guardRight:
					if shadowGrid[1][guardY][guardX] {
						sum++
						break Main
					}
					shadowGrid[1][guardY][guardX] = true
				case guardDown:
					if shadowGrid[2][guardY][guardX] {
						sum++
						break Main
					}
					shadowGrid[2][guardY][guardX] = true
				case guardLeft:
					if shadowGrid[3][guardY][guardX] {
						sum++
						break Main
					}
					shadowGrid[3][guardY][guardX] = true
				}

				guardX, guardY, guardDir = d.nextStep(guardX, guardY, guardDir)
			}

			d.reset()
		}
	}

	return sum
}

func (d *Day06) visited(guardX, guardY int, guardDir tileType) bool {
	visited := guardToVisited(guardDir)

	fmt.Printf("%c - %c\n", visited, d.tiles[guardY][guardX])

	return d.tiles[guardY][guardX] == visited
}
