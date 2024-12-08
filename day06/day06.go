// https://adventofcode.com/2024/day/6

package day06

import (
	"bufio"
	"io"
	"slices"
	"sync"
)

const (
	obsTile    byte = '#'
	gUpTile    byte = '^'
	gRightTile byte = '>'
	gDownTile  byte = 'v'
	gLeftTile  byte = '<'
)

type position struct {
	y int
	x int
}

func (o position) getY() int {
	return o.y
}

func (o position) getX() int {
	return o.x
}

type obstacle struct {
	pos       position
	hitTop    bool
	hitRight  bool
	hitBottom bool
	hitLeft   bool
}

type obstacleDim []*obstacle

type guard struct {
	pos position
	dir position
}

type Day06 struct {
	yMap       []obstacleDim
	xMap       []obstacleDim
	guardStart guard
}

func (d *Day06) LoadInput(input io.Reader) {
	d.guardStart = guard{}
	d.yMap = make([]obstacleDim, 0)

	scanner := bufio.NewScanner(input)

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Bytes()

		d.yMap = append(d.yMap, make(obstacleDim, 0))

		if y == 0 {
			d.xMap = make([]obstacleDim, len(line))
		}

		for x, tile := range line {
			if tile == obsTile {
				obs := &obstacle{
					pos: position{y, x},
				}

				d.yMap[y] = append(d.yMap[y], obs)
				d.xMap[x] = append(d.xMap[x], obs)
				continue
			}

			if tile != gUpTile && tile != gRightTile && tile != gDownTile && tile != gLeftTile {
				continue
			}

			d.guardStart = guard{
				pos: position{y, x},
				dir: position{0, 0},
			}

			switch tile {
			case gUpTile:
				d.guardStart.dir.y = -1
			case gRightTile:
				d.guardStart.dir.x = 1
			case gDownTile:
				d.guardStart.dir.y = 1
			case gLeftTile:
				d.guardStart.dir.x = -1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for y := range d.yMap {
		slices.SortFunc(d.yMap[y], func(a, b *obstacle) int {
			return a.pos.x - b.pos.x
		})
	}

	for x := range d.xMap {
		slices.SortFunc(d.xMap[x], func(a, b *obstacle) int {
			return a.pos.y - b.pos.y
		})
	}
}

func (d *Day06) PartOne() int {
	visited := d.getVisited()
	return len(visited)
}

func (d *Day06) PartTwo() int {
	visited := d.getVisited()

	resChan := make(chan bool, len(visited)-1)
	wg := new(sync.WaitGroup)

	for _, newObsPos := range visited {
		if newObsPos.y == d.guardStart.pos.y && newObsPos.x == d.guardStart.pos.x {
			continue
		}

		wg.Add(1)
		go d.verifyLoop(newObsPos, resChan, wg)
	}

	wg.Wait()

	sum := 0

	for range len(visited) - 1 {
		loop := <-resChan
		if loop {
			sum++
		}
	}
	return sum
}

func (d *Day06) verifyLoop(newObsPos position, resChan chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()

	yMap, xMap := d.getMapCopies()
	g := d.guardStart

	newObs := &obstacle{
		pos: newObsPos,
	}

	yMap[newObsPos.y] = append(yMap[newObsPos.y], newObs)
	slices.SortFunc(yMap[newObsPos.y], func(a, b *obstacle) int {
		return a.pos.x - b.pos.x
	})

	xMap[newObsPos.x] = append(xMap[newObsPos.x], newObs)
	slices.SortFunc(xMap[newObsPos.x], func(a, b *obstacle) int {
		return a.pos.y - b.pos.y
	})

	var nextOb *obstacle
	for {
		if g.dir.y != 0 {
			nextOb = xMap[g.pos.x].nextObstacle(position.getY, g)
		} else {
			nextOb = yMap[g.pos.y].nextObstacle(position.getX, g)
		}

		if nextOb == nil {
			resChan <- false
			break
		}

		if g.moveAndRotate(nextOb) {
			resChan <- true
			break
		}
	}
}

func (d *Day06) getMapCopies() (yMapCopy, xMapCopy []obstacleDim) {
	yMapCopy = make([]obstacleDim, len(d.yMap))
	for y := range yMapCopy {
		yMapCopy[y] = make(obstacleDim, len(d.yMap[y]))
		for x := range yMapCopy[y] {
			yMapCopy[y][x] = &obstacle{}
			*yMapCopy[y][x] = *d.yMap[y][x]
		}
	}

	xMapCopy = make([]obstacleDim, len(d.xMap))
	for x := range xMapCopy {
		xMapCopy[x] = make(obstacleDim, len(d.xMap[x]))
		for y := range xMapCopy[x] {
			xMapCopy[x][y] = &obstacle{}
			*xMapCopy[x][y] = *d.xMap[x][y]
		}
	}

	return yMapCopy, xMapCopy
}

func (d *Day06) getVisited() []position {
	yMap, xMap := d.getMapCopies()
	g := d.guardStart

	gSteps := make([][]bool, len(yMap))
	for y := range yMap {
		gSteps[y] = make([]bool, len(xMap))
	}

	var nextOb *obstacle
	for {
		if g.dir.y != 0 {
			nextOb = xMap[g.pos.x].nextObstacle(position.getY, g)
		} else {
			nextOb = yMap[g.pos.y].nextObstacle(position.getX, g)
		}

		if nextOb == nil {
			break
		}

		if g.dir.y != 0 {
			for y := g.pos.y; y != nextOb.pos.y; y += g.dir.y {
				gSteps[y][g.pos.x] = true
			}
		} else {
			for x := g.pos.x; x != nextOb.pos.x; x += g.dir.x {
				gSteps[g.pos.y][x] = true
			}
		}

		g.moveAndRotate(nextOb)
	}

	// walk guard out of map
	if g.dir.y != 0 {
		for y := g.pos.y; y > 0 && y < len(yMap); y += g.dir.y {
			gSteps[y][g.pos.x] = true
		}
	} else {
		for x := g.pos.x; x > 0 && x < len(xMap); x += g.dir.x {
			gSteps[g.pos.y][x] = true
		}
	}

	visited := make([]position, 0)
	for y := range gSteps {
		for x := range gSteps[y] {
			if gSteps[y][x] {
				visited = append(visited, position{y, x})
			}
		}
	}

	return visited
}

func (oDim obstacleDim) nextObstacle(posFunc func(position) int, g guard) *obstacle {
	if len(oDim) == 0 {
		return nil
	}

	oIndex := oDim.binSearch(posFunc, g, 0, len(oDim)-1)

	if oIndex < 0 || oIndex >= len(oDim) {
		return nil
	}

	return oDim[oIndex]
}

func (oDim obstacleDim) binSearch(posFunc func(position) int, g guard, low, high int) int {
	if high < low {
		if posFunc(g.dir) < 0 {
			return high
		} else {
			return low
		}
	}

	mid := low + (high-low)/2

	if posFunc(oDim[mid].pos) > posFunc(g.pos) {
		return oDim.binSearch(posFunc, g, low, mid-1)
	}

	return oDim.binSearch(posFunc, g, mid+1, high)
}

func (g *guard) moveAndRotate(ob *obstacle) bool {
	g.pos.x = ob.pos.x - g.dir.x
	g.pos.y = ob.pos.y - g.dir.y

	loop := false

	if g.dir.y == -1 {
		if ob.hitBottom {
			loop = true
		}
		ob.hitBottom = true

		g.dir.y = 0
		g.dir.x = 1
	} else if g.dir.x == 1 {
		if ob.hitLeft {
			loop = true
		}
		ob.hitLeft = true

		g.dir.y = 1
		g.dir.x = 0
	} else if g.dir.y == 1 {
		if ob.hitTop {
			loop = true
		}
		ob.hitTop = true

		g.dir.y = 0
		g.dir.x = -1
	} else if g.dir.x == -1 {
		if ob.hitRight {
			loop = true
		}
		ob.hitRight = true

		g.dir.y = -1
		g.dir.x = 0
	} else {
		panic("guard got into unknown direction")
	}

	return loop
}
