// https://adventofcode.com/2024/day/21

package day21

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"regexp"
	"slices"
	"strconv"
)

type Day21 struct {
	codes  [][]byte
	dirMap map[currFromTo][]byte
	numMap map[currFromTo][]byte
}

func (d *Day21) LoadInput(input io.Reader) {
	d.codes = make([][]byte, 0)

	scanner := bufio.NewScanner(input)
	for i := 0; scanner.Scan(); i++ {
		row := scanner.Bytes()

		d.codes = append(d.codes, make([]byte, len(row)))

		copy(d.codes[i], row)
	}

	d.dirMap = dirPad.buildCurrMap(dirPad)
	d.numMap = numPad.buildCurrMap(dirPad)
}

func (d *Day21) PartOne() int {
	return d.solve(4)
}

func (d *Day21) PartTwo() int {
	return d.solve(20)
}

type fromTo struct {
	from byte
	to   byte
}

type currFromTo struct {
	nextCurr byte
	from     byte
	to       byte
}

type keypad [][]byte

const gap = ' '

var numPad = keypad{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{gap, '0', 'A'},
}

var dirPad = keypad{
	{gap, '^', 'A'},
	{'<', 'v', '>'},
}

type position struct {
	y int
	x int
}

type robot struct {
	pathMap  map[currFromTo][]byte
	moves    []byte
	currBtn  byte
	next     *robot
	previous *robot
	kp       keypad
	id       int
}

func (kp keypad) buildCurrMap(nextKp keypad) map[currFromTo][]byte {
	kpMap := kp.buildMap()
	nextKpMap := nextKp.buildMap()

	currFromToMap := make(map[currFromTo][]byte)
	for from := 0; from < len(kp)*len(kp[0]); from++ {
		for to := 0; to < len(kp)*len(kp[0]); to++ {
			for nextCurr := 0; nextCurr < len(nextKp)*len(nextKp[0]); nextCurr++ {
				fromBtn := kp[from/len(kp[0])][from%len(kp[0])]
				toBtn := kp[to/len(kp[0])][to%len(kp[0])]
				nextCurrBtn := nextKp[nextCurr/len(nextKp[0])][nextCurr%len(nextKp[0])]

				if fromBtn == gap || toBtn == gap || nextCurrBtn == gap {
					continue
				}

				paths := kpMap[fromTo{fromBtn, toBtn}]

				var (
					minPathLen = math.MaxInt
					minPath    []byte
				)

				for _, path := range paths {
					testNextCurrBtn := nextCurrBtn

					pathLen := 0
					for _, btn := range path {
						paths := nextKpMap[fromTo{testNextCurrBtn, btn}]
						pathLen += len(paths[0])
						testNextCurrBtn = btn
					}

					if pathLen < minPathLen {
						minPathLen = pathLen
						minPath = path
					}
				}

				currFromToMap[currFromTo{nextCurrBtn, fromBtn, toBtn}] = minPath
			}
		}
	}

	return currFromToMap
}

func (kp keypad) buildMap() map[fromTo][][]byte {
	fromToMap := make(map[fromTo][][]byte)
	for from := 0; from < len(kp)*len(kp[0]); from++ {
		for to := 0; to < len(kp)*len(kp[0]); to++ {
			fromBtn := kp[from/len(kp[0])][from%len(kp[0])]
			toBtn := kp[to/len(kp[0])][to%len(kp[0])]

			if fromBtn == gap || toBtn == gap {
				continue
			}

			paths := kp.dijkstra(position{from / len(kp[0]), from % len(kp[0])}, position{to / len(kp[0]), to % len(kp[0])})
			fromToMap[fromTo{fromBtn, toBtn}] = paths
		}
	}

	return fromToMap
}

func (kp keypad) dijkstra(start, end position) [][]byte {
	finalized := make(map[position]int)

	unvisited := make(map[position]int)
	unvisited[start] = 0

	for {
		var (
			currPos  position
			currDist = math.MaxInt
		)

		for pos, dist := range unvisited {
			if currDist > dist {
				currDist = dist
				currPos = pos
			}
		}

		kp.explore(position{currPos.y - 1, currPos.x}, currDist+1, finalized, unvisited)
		kp.explore(position{currPos.y, currPos.x - 1}, currDist+1, finalized, unvisited)
		kp.explore(position{currPos.y + 1, currPos.x}, currDist+1, finalized, unvisited)
		kp.explore(position{currPos.y, currPos.x + 1}, currDist+1, finalized, unvisited)

		finalized[currPos] = currDist
		delete(unvisited, currPos)

		if currPos.y == end.y && currPos.x == end.x {
			break
		}
	}

	pathToNum := func(path []byte) int {
		num := 0
		for _, btn := range path {
			score := 0

			switch btn {
			case '<':
				score = 1
			case 'v':
				score = 2
			case '^':
				score = 3
			case '>':
				score = 4
			case 'A':
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

	paths := kp.findPaths(start, end, finalized)
	for i := range paths {
		slices.Reverse(paths[i])
		paths[i] = append(paths[i], 'A')
	}

	slices.SortFunc(paths, func(a, b []byte) int {
		return pathToNum(a) - pathToNum(b)
	})

	return paths
}

func (d *Day21) solve(numOfRobots int) int {
	re := regexp.MustCompile(`0?(\d+)A`)

	total := 0
	for _, code := range d.codes {
		numOfPresses := d.enterCode(numOfRobots, code)

		match := re.FindStringSubmatch(string(code))
		numStr := match[1]

		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%d * %d = %d\n", numOfPresses, num, numOfPresses*num)

		total += numOfPresses * num
	}

	return total
}

func (kp keypad) explore(pos position, dist int, finalized, unvisited map[position]int) {
	if pos.y < 0 || pos.y >= len(kp) || pos.x < 0 || pos.x >= len(kp[pos.y]) || kp[pos.y][pos.x] == gap {
		return
	}

	_, ok := finalized[pos]
	if ok {
		return
	}

	oldDist, ok := unvisited[pos]
	if ok && oldDist < dist {
		return
	}

	unvisited[pos] = dist
}

func (kp keypad) findPaths(pos, endPos position, finalized map[position]int) [][]byte {
	if pos.y == endPos.y && pos.x == endPos.x {
		return [][]byte{{}}
	}

	currDist := finalized[pos]

	res := kp.findPathsSub(position{pos.y - 1, pos.x}, endPos, '^', currDist, finalized)
	res = append(res, kp.findPathsSub(position{pos.y, pos.x + 1}, endPos, '>', currDist, finalized)...)
	res = append(res, kp.findPathsSub(position{pos.y + 1, pos.x}, endPos, 'v', currDist, finalized)...)
	res = append(res, kp.findPathsSub(position{pos.y, pos.x - 1}, endPos, '<', currDist, finalized)...)

	return res

}

func (kp keypad) findPathsSub(pos, endPos position, direction byte, currDist int, finalized map[position]int) [][]byte {
	if pos.y < 0 || pos.y >= len(kp) || pos.x < 0 || pos.x >= len(kp[pos.y]) || kp[pos.y][pos.x] == gap {
		return [][]byte{}
	}

	nextDist, ok := finalized[pos]
	if !ok || nextDist <= currDist {
		return [][]byte{}
	}

	paths := kp.findPaths(pos, endPos, finalized)
	for i := range paths {
		paths[i] = append(paths[i], direction)
	}

	return paths
}

func (d *Day21) enterCode(numOfRobots int, code []byte) int {
	robots := make([]*robot, 0)

	numRobot := &robot{
		pathMap: d.numMap,
		moves:   make([]byte, 0),
		currBtn: 'A',
		kp:      numPad,
	}

	robots = append(robots, numRobot)

	for range numOfRobots - 2 {
		robots = append(robots, &robot{
			pathMap: d.dirMap,
			moves:   make([]byte, 0),
			currBtn: 'A',
			kp:      dirPad,
		})
	}

	for i := range numOfRobots - 2 {
		robots[i].next = robots[i+1]
		robots[i].id = i
	}
	robots[len(robots)-1].id = len(robots) - 1

	for _, btn := range code {
		fmt.Printf("%c\n", btn)
		numRobot.press(btn)
	}

	return len(robots[len(robots)-1].moves)
}

func (r *robot) press(nextBtn byte) {
	nextCurr := byte('A')
	if r.next != nil {
		nextCurr = r.next.currBtn
	}

	path := r.pathMap[currFromTo{nextCurr, r.currBtn, nextBtn}]
	r.moves = append(r.moves, path...)

	r.currBtn = nextBtn

	if r.next == nil {
		return
	}

	for _, btn := range path {
		r.next.press(btn)
	}
}
