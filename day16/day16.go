// https://adventofcode.com/2024/day/16

package day16

import (
	"bufio"
	"fmt"
	"io"
	"slices"
)

type Day16 struct {
}

type edge struct {
	x         int
	y         int
	connected []*edge
}

func (d *Day16) LoadInput(input io.Reader) {
	scanner := bufio.NewScanner(input)

	maze := make([][]byte, 0)

	for y := 0; scanner.Scan(); y++ {
		row := scanner.Bytes()

		maze = append(maze, make([]byte, len(row)))
		copy(maze[y], row)
	}

	edges := make([]*edge, 0)
	var start *edge
	var end *edge

	for y, row := range maze {
		for x, val := range row {
			if val == '#' {
				continue
			}

			vert := maze[y-1][x] == '.' || maze[y+1][x] == '.'
			horiz := maze[y][x-1] == '.' || maze[y][x+1] == '.'

			if (vert && horiz) || val == 'S' || val == 'E' {
				edges = append(edges, &edge{x: x, y: y, connected: make([]*edge, 0)})
			}

			if val == 'S' {
				start = edges[len(edges)-1]
			} else if val == 'E' {
				end = edges[len(edges)-1]
			}

		}
	}

	for i1, i2 := 0, 1; i2 < len(edges); i1, i2 = i1+1, i2+1 {
		if edges[i1].y != edges[i2].y {
			continue
		}

		connected := true
		for x := edges[i1].x; x < edges[i2].x; x++ {
			if maze[edges[i2].y][x] == '#' {
				connected = false
				break
			}
		}

		if !connected {
			continue
		}

		edges[i1].connected = append(edges[i1].connected, edges[i2])
		edges[i2].connected = append(edges[i2].connected, edges[i1])
	}

	slices.SortFunc(edges, func(a, b *edge) int {
		if a.x != b.x {
			return a.x - b.x
		}

		return a.y - b.y

	})

	for i1, i2 := 0, 1; i2 < len(edges); i1, i2 = i1+1, i2+1 {
		if edges[i1].x != edges[i2].x {
			continue
		}

		connected := true
		for y := edges[i1].y; y < edges[i2].y; y++ {
			if maze[y][edges[i2].x] == '#' {
				connected = false
				break
			}
		}

		if !connected {
			continue
		}

		edges[i1].connected = append(edges[i1].connected, edges[i2])
		edges[i2].connected = append(edges[i2].connected, edges[i1])
	}

	for _, edg := range edges {
		fmt.Printf("x: %d, y: %d, conn: %d\n", edg.x, edg.y, len(edg.connected))
	}

	fmt.Println(start)
	fmt.Println(end)
}

func (d *Day16) PartOne() int {
	return 0
}

func (d *Day16) PartTwo() int {
	return 0
}
