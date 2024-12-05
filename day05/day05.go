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
	preceding []int
}

type update []int

type Day05 struct {
	pgMap map[int]*page
	ups   []update
}

func (d *Day05) LoadInput(input io.Reader) {
	d.pgMap = make(map[int]*page)
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

	for key := range d.pgMap {
		preceding := d.pgMap[key].preceding
		quicksort(preceding, 0, len(preceding)-1)
	}
}

func (d *Day05) PartOne() int {
	total := 0

	for _, up := range d.ups {
		if d.validUp(up) {
			total += up[len(up)/2]
		}
	}

	return total
}

func (d *Day05) PartTwo() int {
	total := 0

	for _, up := range d.ups {
		if d.validUp(up) {
			continue
		}

		d.quickfix(up, 0, len(up)-1)
		total += up[len(up)/2]
	}

	return total
}

func (d *Day05) parsePgOrder(line string) {
	pgNumsStr := strings.Split(line, "|")

	prePgNum, err := strconv.Atoi(pgNumsStr[0])
	if err != nil {
		panic(err)
	}

	pgNum, err := strconv.Atoi(pgNumsStr[1])
	if err != nil {
		panic(err)
	}

	_, ok := d.pgMap[prePgNum]
	if !ok {
		d.pgMap[prePgNum] = &page{
			num:       prePgNum,
			preceding: make([]int, 0),
		}
	}

	pg, ok := d.pgMap[pgNum]
	if !ok {
		pg = &page{
			num:       pgNum,
			preceding: make([]int, 0),
		}

		d.pgMap[pgNum] = pg
	}

	pg.preceding = append(pg.preceding, prePgNum)
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

func quicksort(slice []int, low, high int) {
	if low >= high || low < 0 {
		return
	}

	p := partition(slice, low, high)

	quicksort(slice, low, p-1)
	quicksort(slice, p+1, high)
}

func partition(slice []int, low, high int) int {
	pivot := slice[high]

	i := low

	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			slice[i], slice[j] = slice[j], slice[i]

			i++
		}
	}

	slice[i], slice[high] = slice[high], slice[i]

	return i
}

func (d *Day05) validUp(up update) bool {
	for i, pgNum := range up {
		pg := d.pgMap[pgNum]

		if pg.precedes(up[i+1:]) {
			return false
		}
	}

	return true
}

func (p page) precedes(prePgNums []int) bool {
	for _, prePgNum := range prePgNums {
		if p.binSearch(0, len(p.preceding)-1, prePgNum) != -1 {
			return true
		}
	}

	return false
}

func (p page) binSearch(low, high, prePgNum int) int {
	if high < low {
		return -1
	}

	mid := low + (high-low)/2

	if p.preceding[mid] == prePgNum {
		return mid
	}

	if p.preceding[mid] > prePgNum {
		return p.binSearch(low, mid-1, prePgNum)
	}

	return p.binSearch(mid+1, high, prePgNum)
}

func (d *Day05) quickfix(up update, low, high int) {
	if low >= high || low < 0 {
		return
	}

	p := d.partitionFix(up, low, high)

	d.quickfix(up, low, p-1)
	d.quickfix(up, p+1, high)
}

func (d *Day05) partitionFix(up update, low, high int) int {
	pivot := up[high]

	i := low

	for j := low; j < high; j++ {
		pg, ok := d.pgMap[up[j]]
		if !ok {
			continue
		}

		if pg.precedes([]int{pivot}) {
			up[i], up[j] = up[j], up[i]

			i++
		}
	}

	up[i], up[high] = up[high], up[i]

	return i
}
