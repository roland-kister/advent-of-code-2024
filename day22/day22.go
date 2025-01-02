// https://adventofcode.com/2024/day/22

package day22

import (
	"bufio"
	"io"
	"strconv"
)

type Day22 struct {
	nums []int
}

func (d *Day22) LoadInput(input io.Reader) {
	d.nums = make([]int, 0)

	scanner := bufio.NewScanner(input)
	for i := 0; scanner.Scan(); i++ {
		numStr := scanner.Text()

		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		d.nums = append(d.nums, num)
	}
}

func (d *Day22) PartOne() int {
	sum := 0

	for _, num := range d.nums {
		num := num

		for range 2000 {
			num = next(num)
		}

		sum += num
	}

	return int(sum)
}

func (d *Day22) PartTwo() int {
	sharedMem := make(map[[4]int]int)

	for _, num := range d.nums {
		var seq [4]int

		mod := num % 10

		for i := range 3 {
			num = next(num)
			seq[i] = (num % 10) - mod
			mod = num % 10
		}

		mem := make(map[[4]int]int)

		for range 1996 {
			num = next(num)
			seq[3] = (num % 10) - mod
			mod = num % 10

			_, ok := mem[seq]
			if !ok {
				mem[seq] = mod
			}

			seq = shift(seq)
		}

		for memSeq, memMod := range mem {
			sharedMem[memSeq] += memMod
		}

	}

	max := 0
	for _, total := range sharedMem {
		if max < total {
			max = total
		}
	}

	return max
}

const prune int = 0b111111111111111111111111

func next(num int) int {
	num = ((num << 6) ^ num) & prune
	num = ((num >> 5) ^ num) & prune
	num = ((num << 11) ^ num) & prune

	return num
}

func shift(seq [4]int) [4]int {
	for i, j := 0, 1; j < 4; i, j = i+1, j+1 {
		seq[i] = seq[j]
	}

	return seq
}
