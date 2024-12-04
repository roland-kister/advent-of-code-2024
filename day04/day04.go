// https://adventofcode.com/2024/day/4

package day04

import (
	"bufio"
	"os"
	"sync"

	"github.com/roland-kister/advent-of-code-2024/internal"
)

// out of bounds
const oob = -1

type grid [][]rune

type day04 struct {
	g grid
}

func NewDay04() internal.Solver {
	return &day04{}
}

func (d *day04) LoadInput(inputPath string) {
	input, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	d.g = make([][]rune, 0)

	y := 0

	for scanner.Scan() {
		row := scanner.Text()

		d.g = append(d.g, make([]rune, len(row)))
		for x, runeVal := range row {
			d.g[y][x] = runeVal
		}

		y++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func (d *day04) PartOne() int {
	return d.sum(grid.countXmas)
}

func (d *day04) PartTwo() int {
	return d.sum(grid.countMas)
}

func (d *day04) sum(sumFn func(g grid, x, y int) int) int {
	yChan := make(chan int, len(d.g))
	sumChan := make(chan int, len(d.g))

	wg := new(sync.WaitGroup)

	for range 16 {
		wg.Add(1)

		go d.g.traverseByRow(sumFn, yChan, sumChan, wg)
	}

	for y := range len(d.g) {
		yChan <- y
	}

	close(yChan)

	wg.Wait()

	total := 0

	close(sumChan)

	for sum := range sumChan {
		total += sum
	}

	return total
}

func (g grid) traverseByRow(sumFn func(g grid, x, y int) int, yChan <-chan int, sumChan chan<- int, wg *sync.WaitGroup) {
	for y := range yChan {
		sum := 0

		for x := range len(g[y]) {
			sum += sumFn(g, x, y)
		}
		
		sumChan <- sum
	}

	wg.Done()
}

func (g grid) at(x, y int) rune {
	if y < 0 || x < 0 {
		return oob
	}

	if y >= len(g) || x >= len(g[y]) {
		return oob
	}

	return g[y][x]
}

func (g grid) countXmas(x, y int) int {
	if g.at(x, y) != 'X' {
		return 0
	}

	total := 0

	total += g.checkXmas(x, y, 0, 1)
	total += g.checkXmas(x, y, 1, 1)
	total += g.checkXmas(x, y, 1, 0)
	total += g.checkXmas(x, y, 1, -1)
	total += g.checkXmas(x, y, 0, -1)
	total += g.checkXmas(x, y, -1, -1)
	total += g.checkXmas(x, y, -1, 0)
	total += g.checkXmas(x, y, -1, 1)

	return total	
}

func (g grid) checkXmas(x, y, xDir, yDir int) int {
	if g.at(x + (xDir * 1), y + (yDir * 1)) != 'M' {
		return 0
	}

	if g.at(x + (xDir * 2), y + (yDir * 2)) != 'A' {
		return 0
	}

	if g.at(x + (xDir * 3), y + (yDir * 3)) != 'S' {
		return 0
	}

	return 1
}

func (g grid) countMas(x, y int) int {
	if g.at(x, y) != 'A' {
		return 0
	}

	rt := g.at(x+1, y+1)
	rb := g.at(x+1, y-1)
	lb := g.at(x-1, y-1)
	lt := g.at(x-1, y+1)

	aDiag := (rt == 'M' && lb == 'S') || (rt == 'S' && lb == 'M')
	bDiag := (lt == 'M' && rb == 'S') || (lt == 'S' && rb == 'M')

	if aDiag && bDiag {
		return 1
	}

	return 0
}
