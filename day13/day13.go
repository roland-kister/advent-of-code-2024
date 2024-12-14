// https://adventofcode.com/2024/day/13

package day13

import (
	"bufio"
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

	// for _, mach := range d.partTwoMachs {
	// 	total += mach.solve()
	// }

	return total
}

func (m machine) solve() int {
	aMaxIter := int(math.Ceil(math.Min(float64(m.prize[0])/float64(m.a[0]), float64(m.prize[1])/float64(m.a[1]))))
	bMaxIter := int(math.Ceil(math.Min(float64(m.prize[0])/float64(m.b[0]), float64(m.prize[1])/float64(m.b[1]))))

	if aMaxIter < bMaxIter {
		return m.aSolvePress(aMaxIter)
	}

	return m.bSolvePress(bMaxIter)
}

func (m machine) aSolvePress(maxIter int) int {
	min := 0

	for aPr := range maxIter {
		rem0 := m.prize[0] - m.a[0]*aPr
		rem1 := m.prize[1] - m.a[1]*aPr

		if rem0/m.b[0] != rem1/m.b[1] || rem0%m.b[0] != 0 || rem1%m.b[1] != 0 {
			continue
		}

		solution := aPr*3 + (rem0 / m.b[0])

		if solution < min || min == 0 {
			min = solution
		}

	}

	return min
}

func (m machine) bSolvePress(maxIter int) int {
	min := 0

	for bPr := range maxIter {
		rem0 := m.prize[0] - m.b[0]*bPr
		rem1 := m.prize[1] - m.b[1]*bPr

		if rem0/m.a[0] != rem1/m.a[1] || rem0%m.a[0] != 0 || rem1%m.a[1] != 0 {
			continue
		}

		solution := bPr + (rem0/m.a[0])*3

		if solution < min || min == 0 {
			min = solution
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
