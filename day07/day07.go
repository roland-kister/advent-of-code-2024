// https://adventofcode.com/2024/day/7

package day07

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"sync"
)

type operator byte

const (
	add operator = 0
	mul operator = 1
	cat operator = 2
)

type Day07 struct {
	res   []uint64
	parts [][]uint64
}

func (d *Day07) LoadInput(input io.Reader) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Split(line, ": ")
		partsStr := strings.Split(words[1], " ")

		num, err := strconv.ParseUint(words[0], 10, 64)
		if err != nil {
			panic(err)
		}

		d.res = append(d.res, num)
		parts := make([]uint64, len(partsStr))

		for i, partStr := range partsStr {
			num, err := strconv.ParseUint(partStr, 10, 64)
			if err != nil {
				panic(err)
			}

			parts[i] = num
		}

		d.parts = append(d.parts, parts)
	}
}

func (d *Day07) PartOne() int {
	return d.getSum(false)
}

func (d *Day07) PartTwo() int {
	return d.getSum(true)
}

func (d *Day07) getSum(useCat bool) int {
	wg := new(sync.WaitGroup)
	resChan := make(chan uint64, len(d.res))

	for i := range d.res {
		wg.Add(1)

		go sumRoutine(d.res[i], d.parts[i], useCat, resChan, wg)
	}

	wg.Wait()
	close(resChan)

	sum := 0

	for res := range resChan {
		sum += int(res)
	}

	return sum
}

func sumRoutine(res uint64, parts []uint64, useCat bool, resChan chan<- uint64, wg *sync.WaitGroup) {
	defer wg.Done()

	opCombs := genOpCombs([]operator{}, len(parts)-1, useCat)

	for _, opComb := range opCombs {
		total := parts[0]

		for i, op := range opComb {
			switch op {
			case add:
				total += parts[i+1]
			case mul:
				total *= parts[i+1]
			case cat:
				total = concat(total, parts[i+1])
			}

			if total > res {
				break
			}
		}

		if total == res {
			resChan <- res
			break
		}
	}
}

func genOpCombs(curr []operator, length int, useCat bool) [][]operator {
	if len(curr) == length {
		return [][]operator{curr}
	}

	currAdd := make([]operator, len(curr))
	copy(currAdd, curr)

	currMul := make([]operator, len(curr))
	copy(currMul, curr)

	res := append(
		genOpCombs(append(currAdd, add), length, useCat),
		genOpCombs(append(currMul, mul), length, useCat)...,
	)

	if !useCat {
		return res
	}

	currCat := make([]operator, len(curr))
	copy(currCat, curr)

	return append(
		res,
		genOpCombs(append(currCat, cat), length, useCat)...,
	)
}

func concat(a, b uint64) uint64 {
	tmp := b
	shift := uint64(1)

	for tmp != 0 {
		shift *= 10
		tmp /= 10
	}

	return a*shift + b
}
