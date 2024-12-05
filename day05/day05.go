// https://adventofcode.com/2024/day/5

package day05

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type page struct {
	num       int
	preceding []*page
}

type update []int

type Day05 struct {
	pgMap map[int]page
	ups   []update
}

func (d *Day05) LoadInput(input io.Reader) {
	d.pgMap = make(map[int]page)
	d.ups = make([]update, 0)

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		switch line := scanner.Text(); {
		case strings.Contains(line, "|"):
			d.parsePgOrder(line)
		case len(line) > 0:
			d.parseUpdate(line)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func (d *Day05) PartOne() int {
	return 0
}

func (d *Day05) PartTwo() int {
	return 0
}

func (d *Day05) parsePgOrder(line string) {
	pgNumsStr := strings.Split(line, "|")

	pgNum, err := strconv.Atoi(pgNumsStr[0])
	if err != nil {
		panic(err)
	}

	prePgNum, err := strconv.Atoi(pgNumsStr[1])
	if err != nil {
		panic(err)
	}

	prePg, ok := d.pgMap[prePgNum]
	if !ok {
		prePg = page{
			num:       prePgNum,
			preceding: make([]*page, 0),
		}

		d.pgMap[prePgNum] = prePg
	}

	pg, ok := d.pgMap[pgNum]
	if !ok {
		pg = page{
			num:       pgNum,
			preceding: make([]*page, 1),
		}

		d.pgMap[pgNum] = pg
	}

	pg.preceding = append(pg.preceding, &prePg)
}

func (d *Day05) parseUpdate(line string) {
	pgNumsStr := strings.Split(line, ",")

	up := make(update, len(pgNumsStr))

	for i, pgNumStr := range pgNumsStr {
		pgNum, err := strconv.Atoi(pgNumStr)
		if err != nil {
			panic(err)
		}

		up[i] = pgNum
	}

	d.ups = append(d.ups, up)
}
