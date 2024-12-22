// https://adventofcode.com/2024/day/18

package day18

import (
	"bufio"
	"errors"
	"io"
	"math"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

type Day18 struct {
	bytesPos  []position
	xMax      int
	yMax      int
	obstCount int
}

func (d *Day18) LoadInput(input io.Reader) {
	d.bytesPos = make([]position, 0)
	d.xMax = 0
	d.yMax = 0
	d.obstCount = 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		splitStr := strings.Split(scanner.Text(), ",")

		x, err := strconv.Atoi(splitStr[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(splitStr[1])
		if err != nil {
			panic(err)
		}

		d.bytesPos = append(d.bytesPos, position{x, y})

		if x > d.xMax {
			d.xMax = x
		}

		if y > d.yMax {
			d.yMax = y
		}
	}
}

func (d *Day18) PartOne() int {
	if d.obstCount == 0 {
		d.obstCount = 1024
	}

	dist, _ := d.solve()

	return dist
}

func (d *Day18) PartTwo() int {
	res := 0

	for i := d.obstCount; i < len(d.bytesPos); i++ {
		d.obstCount = i
		_, err := d.solve()
		if err == nil {
			continue
		}

		xStr := strconv.Itoa(d.bytesPos[i-1].x)
		yStr := strconv.Itoa(d.bytesPos[i-1].y)

		res, _ = strconv.Atoi(strings.Join([]string{xStr, yStr}, ""))
		break
	}

	return res
}

func (d *Day18) solve() (int, error) {
	start := position{0, 0}

	obstacles := make(map[position]bool)
	for i := 0; i < d.obstCount; i++ {
		obstacles[d.bytesPos[i]] = true
	}

	finalized := make(map[position]int)
	finalized[start] = 0

	unvisited := make(map[position]int)
	_, ok := obstacles[position{0, 1}]
	if !ok {
		unvisited[position{0, 1}] = 1
	}

	_, ok = obstacles[position{1, 0}]
	if !ok {
		unvisited[position{1, 0}] = 1
	}

	for {
		var curr position
		currDist := math.MaxInt
		for position, dist := range unvisited {
			if currDist > dist {
				currDist = dist
				curr = position
			}
		}

		top := position{x: curr.x, y: curr.y - 1}
		if isOk(top, finalized, obstacles) && curr.y > 0 {
			setDist(top, currDist+1, unvisited)
		}

		right := position{x: curr.x + 1, y: curr.y}
		if isOk(right, finalized, obstacles) && curr.x < d.xMax {
			setDist(right, currDist+1, unvisited)
		}

		bottom := position{x: curr.x, y: curr.y + 1}
		if isOk(bottom, finalized, obstacles) && curr.y < d.yMax {
			setDist(bottom, currDist+1, unvisited)
		}

		left := position{x: curr.x - 1, y: curr.y}
		if isOk(left, finalized, obstacles) && curr.x > 0 {
			setDist(left, currDist+1, unvisited)
		}

		finalized[curr] = currDist
		delete(unvisited, curr)

		if curr.x == d.xMax && curr.y == d.yMax {
			return currDist, nil
		}

		if len(unvisited) == 0 {
			return 0, errors.New("end not reachable")
		}
	}

}

func isOk(pos position, finalized map[position]int, obstacles map[position]bool) bool {
	_, ok := finalized[pos]
	if ok {
		return false
	}

	_, ok = obstacles[pos]
	return !ok
}

func setDist(pos position, dist int, unvisited map[position]int) {
	existingDist, ok := unvisited[pos]
	if ok && existingDist < dist {
		return
	}

	unvisited[pos] = dist
}
