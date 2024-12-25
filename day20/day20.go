// https://adventofcode.com/2024/day/20

package day20

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

type Day20 struct {
	savedPicoS int
	initRt     *racetrack
	start      position
	end        position
}

type position struct {
	x int
	y int
}

type racetrack struct {
	grid      [][]int
	unvisited map[position]int
}

const (
	obs   = math.MinInt
	empty = math.MinInt / 2
	oob   = -1
)

func (rt *racetrack) get(x, y int) int {
	if y < 0 || y >= len(rt.grid) {
		return oob
	}

	if x < 0 || x >= len(rt.grid[y]) {
		return oob
	}

	return rt.grid[y][x]
}

func (rt *racetrack) set(x, y, v int) {
	rt.grid[y][x] = v
}

func (rt *racetrack) appendRow(rowLen int) {
	rt.grid = append(rt.grid, make([]int, rowLen))
}

func (rt *racetrack) copy() *racetrack {
	newRt := &racetrack{
		grid:      make([][]int, len(rt.grid)),
		unvisited: make(map[position]int),
	}

	for y := range newRt.grid {
		newRt.grid[y] = make([]int, len(rt.grid[y]))
		copy(newRt.grid[y], rt.grid[y])
	}

	return newRt
}

func (rt *racetrack) debug() {
	for y := range rt.grid {
		for x := range rt.grid[y] {
			if rt.grid[y][x] == obs {
				fmt.Print("  # ")
			} else if rt.grid[y][x] == empty {
				fmt.Print("    ")
			} else {
				fmt.Printf("%3d ", rt.grid[y][x])
			}
		}
		fmt.Print("\n")
	}
}

func (rt *racetrack) dijkstra(start position) {
	if rt.get(start.x, start.y-1) == empty {
		rt.unvisited[position{start.x, start.y - 1}] = 1
	}
	if rt.get(start.x+1, start.y) == empty {
		rt.unvisited[position{start.x + 1, start.y}] = 1
	}
	if rt.get(start.x, start.y+1) == empty {
		rt.unvisited[position{start.x, start.y + 1}] = 1
	}
	if rt.get(start.x-1, start.y) == empty {
		rt.unvisited[position{start.x - 1, start.y}] = 1
	}

	rt.set(start.x, start.y, 0)

	for len(rt.unvisited) > 0 {
		var (
			currPos  position
			currDist = math.MaxInt
		)

		for pos, dist := range rt.unvisited {
			if dist < currDist {
				currPos = pos
				currDist = dist
			}
		}

		rt.explore(currPos, currDist, 0, -1)
		rt.explore(currPos, currDist, 1, 0)
		rt.explore(currPos, currDist, 0, 1)
		rt.explore(currPos, currDist, -1, 0)

		rt.set(currPos.x, currPos.y, currDist)
		delete(rt.unvisited, currPos)
	}
}

func (rt *racetrack) explore(currPos position, currDist int, xDir, yDir int) {
	if rt.get(currPos.x+xDir, currPos.y+yDir) != empty {
		return
	}

	next := position{currPos.x + xDir, currPos.y + yDir}

	oldDist, ok := rt.unvisited[next]
	if ok && oldDist < currDist+1 {
		return
	}

	rt.unvisited[next] = currDist + 1
}

func (d *Day20) LoadInput(input io.Reader) {
	d.savedPicoS = 100
	d.initRt = &racetrack{}

	scanner := bufio.NewScanner(input)
	for y := 0; scanner.Scan(); y++ {
		row := scanner.Bytes()

		d.initRt.appendRow(len(row))

		for x, v := range row {
			switch v {
			case '#':
				d.initRt.set(x, y, obs)
			case '.':
				d.initRt.set(x, y, empty)
			case 'S':
				d.initRt.set(x, y, empty)
				d.start = position{x, y}
			case 'E':
				d.initRt.set(x, y, empty)
				d.end = position{x, y}
			}
		}
	}
}

func (d *Day20) PartOne() int {
	startRt := d.initRt.copy()
	startRt.dijkstra(d.start)

	endRt := d.initRt.copy()
	endRt.dijkstra(d.end)

	//for y := range startRt.grid {
	//	for x := range startRt.grid[y] {
	//		startVal := startRt.get(x, y)
	//
	//		topVal := endRt.get(x, y-2)
	//		if topVal == oob || topVal == obs {}
	//
	//	}
	//}

	return 0
}

func (d *Day20) PartTwo() int {

	return 0
}
