// https://adventofcode.com/2024/day/9

package day09

import (
	"bufio"
	"io"
	"math"
)

type diskContent []uint

type Day09 struct {
	disk diskContent
}

const empty uint = math.MaxUint

func (d *Day09) LoadInput(input io.Reader) {
	d.disk = make(diskContent, 0)
	diskMap := make([]uint, 0)

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		inBytes := scanner.Bytes()

		for _, inByte := range inBytes {
			diskMap = append(diskMap, uint(inByte)-'0')
		}
	}

	id := uint(0)
	for i := 0; i < len(diskMap); i++ {
		val := empty
		if i%2 == 0 {
			val = id
			id++
		}

		file := make([]uint, diskMap[i])
		for j := range diskMap[i] {
			file[j] = val
		}

		d.disk = append(d.disk, file...)
	}
}

func (d *Day09) PartOne() int {
	disk := make(diskContent, len(d.disk))
	copy(disk, d.disk)

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

	return disk.getSum()
}

func (d *Day09) PartTwo() int {
	disk := make(diskContent, len(d.disk))
	copy(disk, d.disk)

	fragmentLen := 0

	for i := len(disk) - 1; i > 0; i -= fragmentLen {
		fragmentLen = disk.fragmentLengthRev(i)
		if disk[i] == empty {
			continue
		}

		freeSpaceStart, found := disk.nextFreeSpace(fragmentLen, i)
		if !found {
			continue
		}

		for j := 0; j < fragmentLen; j++ {
			disk[freeSpaceStart+j], disk[i-j] = disk[i-j], disk[freeSpaceStart+j]
		}
	}

	return disk.getSum()
}

func (d diskContent) fragmentLengthRev(index int) int {
	block := d[index]

	length := 0
	for index-length >= 0 && d[index-length] == block {
		length++
	}

	return length
}

func (d diskContent) nextFreeSpace(length, maxOffset int) (int, bool) {
	currLength := 0

	for i := 0; i <= maxOffset; i++ {
		if currLength == length {
			return i - length, true
		}

		if d[i] == empty {
			currLength++
		} else {
			currLength = 0
		}
	}

	return -1, false
}

func (d diskContent) getSum() int {
	sum := 0

	for i := 0; i < len(d); i++ {
		if d[i] == empty {
			continue
		}

		sum += int(d[i]) * i
	}

	return sum
}
