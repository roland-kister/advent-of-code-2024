// https://adventofcode.com/2024/day/21

package day21

import (
	"bufio"
	"io"
	"math"
	"slices"
)

type Day21 struct {
	codes [][]byte
}

func (d *Day21) LoadInput(input io.Reader) {
	d.codes = make([][]byte, 0)

	scanner := bufio.NewScanner(input)
	for i := 0; scanner.Scan(); i++ {
		row := scanner.Bytes()

		d.codes = append(d.codes, make([]byte, len(row)))

		copy(d.codes[i], row)
	}

	fillDirMap()
	fillNumMap()
}

func (d *Day21) PartOne() int {
	return d.solve(3)
}

func (d *Day21) PartTwo() int {
	return d.solve(26)
}

func (d *Day21) solve(robotCount int) int {
	sum := 0

	for _, code := range d.codes {
		currBtns = [26]int{aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn, aBtn}
		total = 0

		num := 0
		for _, val := range code {
			switch val {
			case 'A':
				numRobot(aBtn, robotCount-1)
			case '0':
				numRobot(zeroBtn, robotCount-1)
				num *= 10
				num += 0
			case '1':
				numRobot(oneBtn, robotCount-1)
				num *= 10
				num += 1
			case '2':
				numRobot(twoBtn, robotCount-1)
				num *= 10
				num += 2
			case '3':
				numRobot(threeBtn, robotCount-1)
				num *= 10
				num += 3
			case '4':
				numRobot(fourBtn, robotCount-1)
				num *= 10
				num += 4
			case '5':
				numRobot(fiveBtn, robotCount-1)
				num *= 10
				num += 5
			case '6':
				numRobot(sixBtn, robotCount-1)
				num *= 10
				num += 6
			case '7':
				numRobot(sevenBtn, robotCount-1)
				num *= 10
				num += 7
			case '8':
				numRobot(eightBtn, robotCount-1)
				num *= 10
				num += 8
			case '9':
				numRobot(nineBtn, robotCount-1)
				num *= 10
				num += 9
			}
		}

		sum += total * num
	}

	return sum
}

var currBtns [26]int
var total = 0

func numRobot(btn int, id int) {
	currBtn := currBtns[id]
	nextCurrBtn := currBtns[id-1]

	path := numMap[(nextCurrBtn<<numDoubleShift)|(currBtn<<numShift)|btn]
	for path != 0 {
		dirRobot(path&dirMask, id-1)
		path >>= dirShift
	}

	currBtns[id] = btn
}

func dirRobot(btn int, id int) {
	currBtn := currBtns[id]

	nextCurrBtn := currBtn
	if id > 0 {
		nextCurrBtn = currBtns[id-1]
	}

	path := dirMap[(nextCurrBtn<<dirDoubleShift)|(currBtn<<dirShift)|btn]
	for path != 0 {
		if id == 0 {
			total++
		} else {
			dirRobot(path&dirMask, id-1)
		}
		path >>= dirShift
	}

	currBtns[id] = btn
}

const (
	dirShift       int = 3
	dirDoubleShift int = dirShift * 2
	dirMask        int = 0b111
	numShift       int = 4
	numDoubleShift int = numShift * 2
	numMask        int = 0b1111
)

var (
	dirMap [512]int  = [512]int{}
	numMap [2048]int = [2048]int{}
)

func fillDirMap() {
	startEnd := getDirStartEnd()

	maxRow := len(dirKeypad)
	maxCol := len(dirKeypad[0])

	for start := 0; start < maxRow*maxCol; start++ {
		startRow := start / maxCol
		startCol := start % maxCol
		startBtn := dirKeypad[startRow][startCol]

		if startBtn == gapBtn {
			continue
		}

		for end := 0; end < maxRow*maxCol; end++ {
			endRow := end / maxCol
			endCol := end % maxCol
			endBtn := dirKeypad[endRow][endCol]

			if endBtn == gapBtn {
				continue
			}

			paths := startEnd[(startBtn<<dirShift)|endBtn]

			for next := 0; next < maxRow*maxCol; next++ {
				nextRow := next / maxCol
				nextCol := next % maxCol
				nextBtn := dirKeypad[nextRow][nextCol]

				if nextBtn == gapBtn {
					continue
				}

				var (
					minPathLen = math.MaxInt
					minPath    []int
				)

				for _, path := range paths {
					testNextBtn := nextBtn

					pathLen := 0

					for _, btn := range path {
						nextPaths := startEnd[(testNextBtn<<dirShift)|btn]
						pathLen += len(nextPaths[0])

						testNextBtn = btn
					}

					if pathLen < minPathLen {
						minPathLen = pathLen
						minPath = path
					}
				}

				optimalPath := 0
				for i := len(minPath) - 1; i >= 0; i-- {
					optimalPath = (optimalPath << dirShift) | minPath[i]
				}

				nextStartEnd := (nextBtn << (dirDoubleShift)) | (startBtn << dirShift) | endBtn
				dirMap[nextStartEnd] = optimalPath
			}
		}
	}
}

func fillNumMap() {
	dirStartEnd := getDirStartEnd()

	maxRow := len(numKeypad)
	maxCol := len(numKeypad[0])

	startEnd := [256][][]int{}
	for start := 0; start < maxRow*maxCol; start++ {
		startRow := start / maxCol
		startCol := start % maxCol
		startBtn := numKeypad[startRow][startCol]

		if startBtn == gapBtn {
			continue
		}

		for end := 0; end < maxRow*maxCol; end++ {
			endRow := end / maxCol
			endCol := end % maxCol
			endBtn := numKeypad[endRow][endCol]

			if endBtn == gapBtn {
				continue
			}

			dijkstra := newDijkstra(numKeypad, startRow, startCol, endRow, endCol)
			paths := dijkstra.search()

			startEnd[(startBtn<<numShift)|endBtn] = paths
		}
	}

	for start := 0; start < maxRow*maxCol; start++ {
		startRow := start / maxCol
		startCol := start % maxCol
		startBtn := numKeypad[startRow][startCol]

		if startBtn == gapBtn {
			continue
		}

		for end := 0; end < maxRow*maxCol; end++ {
			endRow := end / maxCol
			endCol := end % maxCol
			endBtn := numKeypad[endRow][endCol]

			if endBtn == gapBtn {
				continue
			}

			paths := startEnd[(startBtn<<numShift)|endBtn]

			nextMaxRow := len(dirKeypad)
			nextMaxCol := len(dirKeypad[0])
			for next := 0; next < nextMaxRow*nextMaxCol; next++ {
				nextRow := next / nextMaxCol
				nextCol := next % nextMaxCol
				nextBtn := dirKeypad[nextRow][nextCol]

				if nextBtn == gapBtn {
					continue
				}

				var (
					minPathLen = math.MaxInt
					minPath    []int
				)

				for _, path := range paths {
					testNextBtn := nextBtn

					pathLen := 0

					for _, btn := range path {
						nextPaths := dirStartEnd[(testNextBtn<<dirShift)|btn]
						pathLen += len(nextPaths[0])

						testNextBtn = btn
					}

					if pathLen < minPathLen {
						minPathLen = pathLen
						minPath = path
					}
				}

				optimalPath := 0
				for i := len(minPath) - 1; i >= 0; i-- {
					optimalPath = (optimalPath << dirShift) | minPath[i]
				}

				nextStartEnd := (nextBtn << (numDoubleShift)) | (startBtn << numShift) | endBtn
				numMap[nextStartEnd] = optimalPath
			}
		}
	}
}

func getDirStartEnd() [64][][]int {
	maxRow := len(dirKeypad)
	maxCol := len(dirKeypad[0])

	startEnd := [64][][]int{}
	for start := 0; start < maxRow*maxCol; start++ {
		startRow := start / maxCol
		startCol := start % maxCol
		startBtn := dirKeypad[startRow][startCol]

		if startBtn == gapBtn {
			continue
		}

		for end := 0; end < maxRow*maxCol; end++ {
			endRow := end / maxCol
			endCol := end % maxCol
			endBtn := dirKeypad[endRow][endCol]

			if endBtn == gapBtn {
				continue
			}

			dijkstra := newDijkstra(dirKeypad, startRow, startCol, endRow, endCol)
			paths := dijkstra.search()

			startEnd[(startBtn<<dirShift)|endBtn] = paths
		}
	}

	return startEnd
}

const (
	aBtn     = 1
	upBtn    = 2
	rightBtn = 3
	downBtn  = 4
	leftBtn  = 5
	gapBtn   = math.MaxInt
	sevenBtn = 7
	eightBtn = 8
	nineBtn  = 9
	fourBtn  = 4
	fiveBtn  = 5
	sixBtn   = 6
	oneBtn   = 11
	twoBtn   = 2
	threeBtn = 3
	zeroBtn  = 10
)

var dirKeypad [][]int = [][]int{
	{gapBtn, upBtn, aBtn},
	{leftBtn, downBtn, rightBtn},
}

var numKeypad [][]int = [][]int{
	{sevenBtn, eightBtn, nineBtn},
	{fourBtn, fiveBtn, sixBtn},
	{oneBtn, twoBtn, threeBtn},
	{gapBtn, zeroBtn, aBtn},
}

type position struct {
	row int
	col int
}

type dijkstra struct {
	keypad    [][]int
	start     position
	end       position
	finalized map[position]int
	unvisited map[position]int
}

func newDijkstra(keypad [][]int, startRow, startCol, endRow, endCol int) dijkstra {
	return dijkstra{
		keypad:    keypad,
		start:     position{startRow, startCol},
		end:       position{endRow, endCol},
		finalized: make(map[position]int),
		unvisited: make(map[position]int),
	}
}

func (d dijkstra) search() [][]int {
	d.unvisited[d.start] = 0

	for {
		var (
			currPos  position
			currDist = math.MaxInt
		)

		for pos, dist := range d.unvisited {
			if currDist > dist {
				currPos = pos
				currDist = dist
			}
		}

		d.explore(position{currPos.row - 1, currPos.col}, currDist+1)
		d.explore(position{currPos.row, currPos.col - 1}, currDist+1)
		d.explore(position{currPos.row + 1, currPos.col}, currDist+1)
		d.explore(position{currPos.row, currPos.col + 1}, currDist+1)

		d.finalized[currPos] = currDist
		delete(d.unvisited, currPos)

		if currPos.row == d.end.row && currPos.col == d.end.col {
			break
		}
	}

	paths := d.findPaths(d.start)
	for i := range paths {
		slices.Reverse(paths[i])
		paths[i] = append(paths[i], aBtn)
	}

	pathToNum := func(path []int) int {
		num := 0
		for _, btn := range path {
			score := 0

			switch btn {
			case leftBtn:
				score = 1
			case downBtn:
				score = 2
			case upBtn:
				score = 3
			case rightBtn:
				score = 4
			case aBtn:
				score = 5
			}

			if num%10 == score {
				continue
			}

			num *= 10
			num += score
		}

		return num
	}

	slices.SortFunc(paths, func(a, b []int) int {
		return pathToNum(a) - pathToNum(b)
	})

	return paths
}

func (d dijkstra) explore(pos position, dist int) {
	if pos.row < 0 || pos.row >= len(d.keypad) || pos.col < 0 || pos.col >= len(d.keypad[pos.row]) {
		return
	}

	if d.keypad[pos.row][pos.col] == gapBtn {
		return
	}

	_, ok := d.finalized[pos]
	if ok {
		return
	}

	oldDist, ok := d.unvisited[pos]
	if ok && oldDist < dist {
		return
	}

	d.unvisited[pos] = dist
}

func (d dijkstra) findPaths(pos position) [][]int {
	if pos.row == d.end.row && pos.col == d.end.col {
		return [][]int{{}}
	}

	dist := d.finalized[pos]

	res := d.findPathsSub(position{pos.row - 1, pos.col}, upBtn, dist)
	res = append(res, d.findPathsSub(position{pos.row, pos.col + 1}, rightBtn, dist)...)
	res = append(res, d.findPathsSub(position{pos.row + 1, pos.col}, downBtn, dist)...)
	res = append(res, d.findPathsSub(position{pos.row, pos.col - 1}, leftBtn, dist)...)

	return res

}

func (d dijkstra) findPathsSub(pos position, dir int, currDist int) [][]int {
	if pos.row < 0 || pos.row >= len(d.keypad) || pos.col < 0 || pos.col >= len(d.keypad[pos.row]) {
		return [][]int{}
	}

	if d.keypad[pos.row][pos.col] == gapBtn {
		return [][]int{}
	}

	nextDist, ok := d.finalized[pos]
	if !ok || nextDist <= currDist {
		return [][]int{}
	}

	paths := d.findPaths(pos)
	for i := range paths {
		paths[i] = append(paths[i], dir)
	}

	return paths
}
