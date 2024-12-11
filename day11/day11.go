// https://adventofcode.com/2024/day/11

package day11

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type stoneMap map[uint64]uint64

type Day11 struct {
	stones stoneMap
}

func (d *Day11) LoadInput(input io.Reader) {
	d.stones = make(stoneMap)

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	inputLine := scanner.Text()

	stonesStr := strings.Split(inputLine, " ")
	for _, stoneStr := range stonesStr {
		val, err := strconv.Atoi(stoneStr)
		if err != nil {
			panic(err)
		}

		d.stones[uint64(val)] = 1
	}
}

func (d *Day11) PartOne() int {
	return d.solve(25)
}

func (d *Day11) PartTwo() int {
	return d.solve(75)
}

func (d *Day11) solve(iterCount int) int {
	stones := make(stoneMap)

	for k, v := range d.stones {
		stones[k] = v
	}

	for range iterCount {
		stones = blink(stones)
	}

	sum := 0

	for _, count := range stones {
		sum += int(count)
	}

	return sum
}

func blink(stones stoneMap) stoneMap {
	newStones := make(stoneMap)
	for stone, count := range stones {
		if stone == 0 {
			newStones[1] += count
			continue
		}

		divider := divider(stone)

		if divider == 0 {
			newStones[stone*2024] += count
			continue
		}

		newStones[stone/divider] += count
		newStones[stone%divider] += count
	}

	return newStones
}

func divider(n uint64) uint64 {
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
