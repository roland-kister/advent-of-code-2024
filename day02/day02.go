// https://adventofcode.com/2024/day/2

package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/roland-kister/advent-of-code-2024/internal"
)

type day02 struct {
	reps [][]int
}

func NewDay02() internal.Solver {
	return &day02{}
}

func (d *day02) LoadInput(inputPath string) {
	d.reps = make([][]int, 0)

	input, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lvls := strings.Split(scanner.Text(), " ")

		rep := make([]int, len(lvls))

		for j, lvl := range lvls {
			lvlNum, err := strconv.Atoi(lvl)
			if err != nil {
				panic(err)
			}

			rep[j] = lvlNum
		}

		d.reps = append(d.reps, rep)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func (d *day02) PartOne() int {
	safeCount := 0

	for _, rep := range d.reps {
		if isRepSafe(rep) {
			safeCount++
		}
	}

	return safeCount
}

func (d *day02) PartTwo() int {
	safeTolChan := make(chan bool, len(d.reps))

	wg := new(sync.WaitGroup)

	for _, rep := range d.reps {
		wg.Add(1)
		go isRepSafeTol(rep, safeTolChan, wg)
	}

	wg.Wait()

	safeTolCount := 0

	for range len(d.reps) {
		if <-safeTolChan {
			safeTolCount++
		}
	}

	close(safeTolChan)

	return safeTolCount
}

func isRepSafe(rep []int) bool {
	if len(rep) < 2 {
		panic(fmt.Sprintf("report has less than 2 levels (%d)", len(rep)))
	}

	inc := rep[0] > rep[1]

	cmp := func(a, b int) bool {
		return a < b && b-a <= 3
	}

	if inc {
		cmp = func(a, b int) bool {
			return a > b && a-b <= 3
		}
	}

	for i := 0; i < len(rep)-1; i++ {
		if !cmp(rep[i], rep[i+1]) {
			return false
		}
	}

	return true
}

func isRepSafeTol(rep []int, resChan chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()

	if isRepSafe(rep) {
		resChan <- true
		return
	}

	for i := range len(rep) {
		toleradRep := removeEl(rep, i)

		if isRepSafe(toleradRep) {
			resChan <- true
			return
		}
	}

	resChan <- false
}

func removeEl(slice []int, index int) []int {
	res := make([]int, len(slice)-1)

	shift := 0

	for i, el := range slice {
		if i == index {
			shift--
			continue
		}

		res[i+shift] = el
	}

	return res
}
