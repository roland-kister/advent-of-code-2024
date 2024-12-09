// https://adventofcode.com/2024/day/9

package day09

import (
	"bufio"
	"io"
	"math"
)

type Day09 struct {
	diskMap   []uint
	fragments []fragment
}

type fragment struct {
	id     uint
	length uint
}

const empty uint = math.MaxUint

func (d *Day09) LoadInput(input io.Reader) {
	d.diskMap = make([]uint, 0)
	d.fragments = make([]fragment, 0)

	scanner := bufio.NewScanner(input)

	id := uint(0)
	file := true

	for scanner.Scan() {
		inBytes := scanner.Bytes()

		for _, inByte := range inBytes {
			d.diskMap = append(d.diskMap, uint(inByte)-'0')

			currFragment := fragment{
				id:     id,
				length: uint(inByte) - '0',
			}

			if !file {
				currFragment.id = empty
				id++
			}

			d.fragments = append(d.fragments, currFragment)

			file = !file
		}
	}
}

func (d *Day09) PartOne() int {
	disk := make([]uint, 0)

	id := uint(0)
	for i := 0; i < len(d.diskMap); i++ {
		val := empty
		if i%2 == 0 {
			val = id
			id++
		}

		file := make([]uint, d.diskMap[i])
		for j := range d.diskMap[i] {
			file[j] = val
		}

		disk = append(disk, file...)
	}

	last := len(disk) - 1
	for i := 0; i < len(disk) && i < last; i++ {
		if disk[i] != empty {
			continue
		}

		disk[i], disk[last] = disk[last], disk[i]
		last--

		for last > i && disk[last] == empty {
			last--
		}

	}

	sum := 0

	for i := 0; i < len(disk); i++ {
		if disk[i] == empty {
			break
		}

		sum += int(disk[i]) * i
	}

	return sum
}

func (d *Day09) PartTwo() int {
	return 0
}
