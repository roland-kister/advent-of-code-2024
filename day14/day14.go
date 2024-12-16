// https://adventofcode.com/2024/day/14

package day14

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

type Day14 struct {
	robs []robot
	xMax int
	yMax int
}

type robot struct {
	p [2]int
	v [2]int
}

func (d *Day14) LoadInput(input io.Reader) {
	d.robs = make([]robot, 0)
	d.xMax = 0
	d.yMax = 0

	re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		match := re.FindStringSubmatch(scanner.Text())
		pXstr := match[1]
		pYstr := match[2]
		vXstr := match[3]
		vYstr := match[4]

		pX, err := strconv.Atoi(pXstr)
		if err != nil {
			panic(err)
		}

		pY, err := strconv.Atoi(pYstr)
		if err != nil {
			panic(err)
		}

		vX, err := strconv.Atoi(vXstr)
		if err != nil {
			panic(err)
		}

		vY, err := strconv.Atoi(vYstr)
		if err != nil {
			panic(err)
		}

		d.robs = append(d.robs, robot{
			p: [2]int{pX, pY},
			v: [2]int{vX, vY},
		})

		if pX+1 > d.xMax {
			d.xMax = pX + 1
		}

		if pY+1 > d.yMax {
			d.yMax = pY + 1
		}
	}
}

func (d *Day14) PartOne() int {
	midX := d.xMax / 2
	midY := d.yMax / 2

	quarters := [2][2]int{{0, 0}, {0, 0}}

	for _, rob := range d.robs {
		pX := (rob.p[0] + rob.v[0]*100) % d.xMax
		if pX < 0 {
			pX = d.xMax + pX
		}

		pY := (rob.p[1] + rob.v[1]*100) % d.yMax
		if pY < 0 {
			pY = d.yMax + pY
		}

		if pX == midX || pY == midY {
			continue
		}

		xQ := 0
		if pX > midX {
			xQ = 1
		}

		yQ := 0
		if pY > midY {
			yQ = 1
		}

		quarters[xQ][yQ]++
	}

	return quarters[0][0] * quarters[0][1] * quarters[1][0] * quarters[1][1]
}

func (d *Day14) PartTwo() int {

	deb := make([][]byte, d.yMax)
	for y := range d.yMax {
		deb[y] = make([]byte, d.xMax)

		for x := range d.xMax {
			deb[y][x] = 0
		}
	}

	i := 0
Main:
	for ; ; i++ {
		for y := range d.yMax {
			for x := range d.xMax {
				deb[y][x] = 0
			}
		}

		for _, rob := range d.robs {
			pX := (rob.p[0] + rob.v[0]*i) % d.xMax
			if pX < 0 {
				pX = d.xMax + pX
			}

			pY := (rob.p[1] + rob.v[1]*i) % d.yMax
			if pY < 0 {
				pY = d.yMax + pY
			}

			if deb[pY][pX] != 0 {
				continue Main
			}

			deb[pY][pX]++
		}
		break
	}

	return i
}
