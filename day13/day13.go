// https://adventofcode.com/2024/day/13

package day13

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"
)

type machine struct {
	a     [2]int
	b     [2]int
	prize [2]int
}

const p2Mul = 10_000_000_000_000

type Day13 struct {
	partOneMachs []machine
	partTwoMachs []machine
}

func (d *Day13) LoadInput(input io.Reader) {
	d.partOneMachs = make([]machine, 0)
	d.partTwoMachs = make([]machine, 0)

	scanner := bufio.NewScanner(input)

	for i := 0; scanner.Scan(); i++ {
		aButton := scanner.Text()

		scanner.Scan()
		bButton := scanner.Text()

		scanner.Scan()
		prize := scanner.Text()

		machOne := machine{
			a:     parseXandY(aButton),
			b:     parseXandY(bButton),
			prize: parseXandY(prize),
		}
		d.partOneMachs = append(d.partOneMachs, machOne)

		machTwo := machine{
			a:     [2]int{machOne.a[0], machOne.a[1]},
			b:     [2]int{machOne.b[0], machOne.b[1]},
			prize: [2]int{machOne.prize[0] + p2Mul, machOne.prize[1] + p2Mul},
		}
		d.partTwoMachs = append(d.partTwoMachs, machTwo)

		scanner.Scan()
	}
}

func (d *Day13) PartOne() int {
	total := 0

	for _, mach := range d.partOneMachs {
		total += mach.solve()
	}

	return total
}

func (d *Day13) PartTwo() int {
	total := 0

	for _, mach := range d.partTwoMachs {
		total += mach.solve()
	}

	return total
}

func (m machine) solve() int {
	maxIter := int(math.Ceil(math.Min(float64(m.prize[0])/float64(m.a[0]), float64(m.prize[1])/float64(m.a[1]))))

	startX, stepX, err := m.findStartAndStep(0)
	if err != nil {
		return 0
	}

	startY, stepY, err := m.findStartAndStep(1)
	if err != nil {
		return 0
	}

	if stepX > stepY {
		return m.findMinTokens(startX, stepX, maxIter)
	}

	return m.findMinTokens(startY, stepY, maxIter)
}

func (m machine) findStartAndStep(idx int) (start, step int, err error) {
	start, step = 0, 1
	bMods := make(map[int]int)

	if m.a[0]%m.b[0] == 0 || m.b[0]%m.a[0] == 0 {
		return
	}

	aPrMax := m.prize[idx]/m.a[idx] + 1

	aPr := 0
	if m.prize[idx]%m.b[idx] == 0 {
		aPr++
	}

	for ; aPr < aPrMax; aPr++ {
		bMod := (m.prize[idx] - m.a[idx]*aPr) % m.b[idx]

		if bMod == 0 {
			break
		}

		_, ok := bMods[bMod]
		if ok {
			return 0, 0, errors.New("not winnable")
		}

		bMods[bMod] = aPr
	}

	start, step = aPr, 1
	firstMod := m.prize[idx] % m.b[idx]

	for ; aPr < aPrMax; aPr++ {
		if (m.prize[idx]-m.a[idx]*aPr)%m.b[idx] == firstMod {
			step = aPr
			break
		}
	}

	return
}

func (m machine) findMinTokens(start, step, maxIter int) int {
	min := 0

	for aPr := start; aPr < maxIter; aPr += step {
		rem0 := m.prize[0] - m.a[0]*aPr
		rem1 := m.prize[1] - m.a[1]*aPr

		if rem0/m.b[0] != rem1/m.b[1] || rem0%m.b[0] != 0 || rem1%m.b[1] != 0 {
			continue
		}

		solution := aPr*3 + (rem0 / m.b[0])

		if solution < min || min == 0 {
			min = solution
		} else if solution > min {
			fmt.Println("break")
			break
		}

	}

	return min
}

func parseXandY(text string) [2]int {
	re := regexp.MustCompile(`(?m)X.(\d+), Y.(\d+)`)

	match := re.FindStringSubmatch(text)
	xStr := match[1]
	yStr := match[2]

	x, err := strconv.Atoi(xStr)
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(yStr)
	if err != nil {
		panic(err)
	}

	return [2]int{x, y}
}
