// https://adventofcode.com/2024/day/15

package day15

import (
	"bufio"
	"io"
)

type warehouse [][]byte

type Day15 struct {
	wh    warehouse
	wh2   warehouse
	dirs  []byte
	rPos  [2]int
	rPos2 [2]int
}

func (d *Day15) LoadInput(input io.Reader) {
	d.wh = make([][]byte, 0)
	d.dirs = make([]byte, 0)

	scanner := bufio.NewScanner(input)

	for y := 0; scanner.Scan(); y++ {
		row := scanner.Bytes()
		if len(row) == 0 {
			break
		}

		d.wh = append(d.wh, make([]byte, len(row)))
		copy(d.wh[y], row)

		for x, val := range row {
			if val == '@' {
				d.rPos = [2]int{y, x}
			}
		}

		d.wh2 = append(d.wh2, make([]byte, len(row)*2))
		for x, i := 0, 0; i < len(row); x, i = x+2, i+1 {
			switch row[i] {
			case '#':
				d.wh2[y][x] = '#'
				d.wh2[y][x+1] = '#'
			case 'O':
				d.wh2[y][x] = '['
				d.wh2[y][x+1] = ']'
			case '.':
				d.wh2[y][x] = '.'
				d.wh2[y][x+1] = '.'
			case '@':
				d.wh2[y][x] = '@'
				d.wh2[y][x+1] = '.'
				d.rPos2 = [2]int{y, x}
			}
		}
	}

	for scanner.Scan() {
		row := scanner.Bytes()
		d.dirs = append(d.dirs, row...)
	}
}

func (d *Day15) PartOne() int {
	for _, dir := range d.dirs {
		d.wh.partOneMove(&d.rPos, dir)
	}

	total := 0
	for y, row := range d.wh {
		for x, val := range row {
			if val == 'O' {
				total += y*100 + x
			}
		}
	}

	return total
}

func (wh warehouse) partOneMove(rPos *[2]int, dir byte) {
	yDir, xDir := 0, 0
	switch dir {
	case '^':
		yDir = -1
	case '>':
		xDir = 1
	case 'v':
		yDir = 1
	case '<':
		xDir = -1
	}

	switch wh[rPos[0]+yDir][rPos[1]+xDir] {
	case '#':
		break
	case '.':
		wh[rPos[0]][rPos[1]] = '.'
		rPos[0], rPos[1] = rPos[0]+yDir, rPos[1]+xDir
		wh[rPos[0]][rPos[1]] = '@'
	case 'O':
		wh.partOneBox(rPos, yDir, xDir)
	}
}

func (wh warehouse) partOneBox(rPos *[2]int, yDir, xDir int) {
	y, x := rPos[0]+yDir, rPos[1]+xDir

	for ; wh[y][x] != '#'; y, x = y+yDir, x+xDir {
		if wh[y][x] != '.' {
			continue
		}

		wh[y][x] = 'O'
		wh[rPos[0]][rPos[1]] = '.'
		rPos[0], rPos[1] = rPos[0]+yDir, rPos[1]+xDir
		wh[rPos[0]][rPos[1]] = '@'
		break
	}
}

func (d *Day15) PartTwo() int {
	for _, dir := range d.dirs {
		d.wh2.partTwoMove(&d.rPos2, dir)

	}

	total := 0
	for y, row := range d.wh2 {
		for x, val := range row {
			if val == '[' {
				total += y*100 + x
			}
		}
	}

	return total
}

func (wh warehouse) partTwoMove(rPos *[2]int, dir byte) {
	yDir, xDir := 0, 0
	switch dir {
	case '^':
		yDir = -1
	case '>':
		xDir = 1
	case 'v':
		yDir = 1
	case '<':
		xDir = -1
	}

	switch wh[rPos[0]+yDir][rPos[1]+xDir] {
	case '#':
		break
	case '.':
		wh[rPos[0]][rPos[1]] = '.'
		rPos[0], rPos[1] = rPos[0]+yDir, rPos[1]+xDir
		wh[rPos[0]][rPos[1]] = '@'
	case '[', ']':
		if xDir != 0 {
			wh.partTwoBoxHorizontal(rPos, xDir)
		} else {
			wh.partTwoBoxVertical(rPos, yDir)
		}
	}
}

func (wh warehouse) partTwoBoxHorizontal(rPos *[2]int, xDir int) {
	y, x := rPos[0], rPos[1]+xDir

	for ; wh[rPos[0]][x] != '#'; x = x + xDir {
		if wh[y][x] == '.' {
			break
		}
	}

	if wh[y][x] != '.' {
		return
	}

	char := byte('[')
	if xDir > 0 {
		char = ']'
	}

	for i := x; i != rPos[1]+xDir; i += (xDir * -1) {
		wh[y][i] = char

		if char == '[' {
			char = ']'
		} else {
			char = '['
		}

	}

	wh[rPos[0]][rPos[1]] = '.'
	rPos[1] += xDir
	wh[rPos[0]][rPos[1]] = '@'
}

func (wh warehouse) partTwoBoxVertical(rPos *[2]int, yDir int) {
	xOff := 0

	if wh[rPos[0]+yDir][rPos[1]] == ']' {
		xOff--
	}

	if !wh.partTwoVertPossible(rPos[0]+yDir, rPos[1]+xOff, yDir) {
		return
	}

	wh.partTwoVertMove(rPos[0]+yDir, rPos[1]+xOff, yDir)

	wh[rPos[0]][rPos[1]] = '.'
	rPos[0] += yDir
	wh[rPos[0]][rPos[1]] = '@'
}

func (wh warehouse) partTwoVertPossible(y, x, yDir int) bool {
	if wh[y+yDir][x] == '#' || wh[y+yDir][x+1] == '#' {
		return false
	}

	if wh[y+yDir][x] == '.' && wh[y+yDir][x+1] == '.' {
		return true
	}

	if wh[y+yDir][x] == '[' {
		return wh.partTwoVertPossible(y+yDir, x, yDir)
	}

	possible := true

	if wh[y+yDir][x] == ']' {
		possible = possible && wh.partTwoVertPossible(y+yDir, x-1, yDir)
	}

	if wh[y+yDir][x+1] == '[' {
		possible = possible && wh.partTwoVertPossible(y+yDir, x+1, yDir)
	}

	return possible
}

func (wh warehouse) partTwoVertMove(y, x, yDir int) {
	if wh[y+yDir][x] == '[' {
		wh.partTwoVertMove(y+yDir, x, yDir)
	}

	if wh[y+yDir][x] == ']' {
		wh.partTwoVertMove(y+yDir, x-1, yDir)
	}

	if wh[y+yDir][x+1] == '[' {
		wh.partTwoVertMove(y+yDir, x+1, yDir)
	}

	wh[y+yDir][x], wh[y+yDir][x+1] = '[', ']'
	wh[y][x], wh[y][x+1] = '.', '.'
}
