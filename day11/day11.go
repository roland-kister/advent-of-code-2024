// https://adventofcode.com/2024/day/11

package day11

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"sync"
)

type Day11 struct {
	initStone *stone
}

type stone struct {
	val       uint64
	nextStone *stone
}

func (d *Day11) LoadInput(input io.Reader) {
	d.initStone = nil
	var prevStone *stone = nil

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	inputLine := scanner.Text()

	stonesStr := strings.Split(inputLine, " ")
	for _, stoneStr := range stonesStr {
		val, err := strconv.Atoi(stoneStr)
		if err != nil {
			panic(err)
		}

		currStone := &stone{val: uint64(val)}

		if prevStone == nil {
			d.initStone = currStone
			prevStone = currStone
			continue
		}

		prevStone.nextStone = currStone
		prevStone = currStone
	}
}

func (d *Day11) PartOne() int {
	return d.solve(25)
}

func (d *Day11) PartTwo() int {
	return d.solve(25)
}

func (d *Day11) solve(iterCount int) int {
	initStone := d.copyInitStone()

	initStoneCount := 0
	for s := initStone; s != nil; s = s.nextStone {
		initStoneCount++
	}

	sumChan := make(chan int, initStoneCount)
	wg := new(sync.WaitGroup)

	for s := initStone; s != nil; s = s.nextStone {
		wg.Add(1)
		go s.solveRoutine(iterCount, sumChan, wg)
	}

	wg.Wait()
	close(sumChan)

	sum := 0
	for subSum := range sumChan {
		sum += subSum
	}

	return sum
}

func (s *stone) solveRoutine(iterCount int, sumChan chan<- int, wg *sync.WaitGroup) {
	s.nextStone = nil

	for range iterCount {
		curr := s

		for curr != nil {
			switch {
			case curr.val == 0:
				curr.val = 1
				curr = curr.nextStone
			case curr.split():
				curr = curr.nextStone.nextStone
			default:
				curr.val *= 2024
				curr = curr.nextStone
			}
		}
	}

	sum := 0
	for curr := s; curr != nil; curr = curr.nextStone {
		sum++
	}

	sumChan <- sum
	wg.Done()
}

func (d *Day11) copyInitStone() *stone {
	var initStone *stone = nil

	var prev *stone = nil

	for s := d.initStone; s != nil; s = s.nextStone {
		curr := &stone{
			val:       s.val,
			nextStone: nil,
		}

		if prev == nil {
			initStone = curr
		} else {
			prev.nextStone = curr
		}

		prev = curr
	}

	return initStone
}

func (s *stone) divider() uint64 {
	n := s.val

	switch {
	case n == 0:
		return 0
	case n < 10:
		return 0
	case n < 100:
		return 10
	case n < 1_000:
		return 0
	case n < 10_000:
		return 100
	case n < 100_000:
		return 0
	case n < 1_000_000:
		return 1_000
	case n < 10_000_000:
		return 0
	case n < 100_000_000:
		return 10_000
	case n < 1_000_000_000:
		return 0
	case n < 10_000_000_000:
		return 100_000
	case n < 100_000_000_000:
		return 0
	case n < 1_000_000_000_000:
		return 1_000_000
	case n < 10_000_000_000_000:
		return 0
	case n < 100_000_000_000_000:
		return 10_000_000
	case n < 1_000_000_000_000_000:
		return 0
	case n < 10_000_000_000_000_000:
		return 100_000_000
	case n < 100_000_000_000_000_000:
		return 0
	case n < 1_000_000_000_000_000_000:
		return 1_000_000_000
	default:
		return 100_000
	}
}

func (s *stone) split() bool {
	divider := s.divider()

	if divider == 0 {
		return false
	}

	newVal := s.val % divider
	s.val = s.val / divider

	newStone := &stone{
		val:       newVal,
		nextStone: s.nextStone,
	}

	s.nextStone = newStone

	return true
}
