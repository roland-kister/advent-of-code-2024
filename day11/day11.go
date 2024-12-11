// https://adventofcode.com/2024/day/11

package day11

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Day11 struct {
	initStone *stone
}

type stone struct {
	val       int
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

		currStone := &stone{val: val}

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
	return d.solve(75)
}

func (d *Day11) solve(iterCount int) int {
	initStone := d.copyInitStone()

	for range iterCount {
		curr := initStone

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
	for s := initStone; s != nil; s = s.nextStone {
		sum++
	}

	return sum
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

func (s *stone) countDigits() int {
	tmp := s.val
	count := 0

	for tmp != 0 {
		tmp /= 10
		count++
	}

	return count
}

func (s *stone) split() bool {
	digitCount := s.countDigits()

	if digitCount%2 != 0 {
		return false
	}

	newStone := &stone{
		val:       0,
		nextStone: s.nextStone,
	}

	s.nextStone = newStone

	multiplier := 1
	for range digitCount / 2 {
		part := (s.val % 10) * multiplier
		newStone.val += part
		multiplier *= 10

		s.val /= 10
	}

	return true
}
