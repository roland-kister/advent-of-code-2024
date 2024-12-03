// https://adventofcode.com/2024/day/1

package day01

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/roland-kister/advent-of-code-2024/internal"
)

type day01 struct {
	lefts  []int
	rights []int
}

func NewDay01() internal.Solver {
	return &day01{}
}

func (d *day01) LoadInput(inputPath string) {
	d.lefts = make([]int, 0)
	d.rights = make([]int, 0)

	input, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), "   ")

		left, err := strconv.Atoi(words[0])
		if err != nil {
			panic(err)
		}

		d.lefts = append(d.lefts, left)

		right, err := strconv.Atoi(words[1])
		if err != nil {
			panic(err)
		}

		d.rights = append(d.rights, right)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func (d day01) PartOne() int {
	quicksort(d.lefts, 0, len(d.lefts)-1)
	quicksort(d.rights, 0, len(d.rights)-1)

	diff := 0

	for i := 0; i < len(d.lefts); i++ {
		diff += abs(d.lefts[i] - d.rights[i])
	}

	return diff
}

func (d day01) PartTwo() int {
	quicksort(d.lefts, 0, len(d.lefts)-1)
	quicksort(d.rights, 0, len(d.rights)-1)

	occurMapLen := d.lefts[len(d.lefts)-1] + 1
	if d.lefts[len(d.lefts)-1] < d.rights[len(d.rights)-1] {
		occurMapLen = d.rights[len(d.rights)-1] + 1
	}

	occurMap := make([]int, occurMapLen)

	for _, right := range d.rights {
		occurMap[right]++
	}

	simScore := 0

	for _, left := range d.lefts {
		simScore += occurMap[left] * left
	}

	return simScore
}

func quicksort(slice []int, low int, high int) {
	if low >= high || low < 0 {
		return
	}

	p := partition(slice, low, high)

	quicksort(slice, low, p-1)
	quicksort(slice, p+1, high)
}

func partition(slice []int, low int, high int) int {
	pivot := slice[high]

	i := low

	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			slice[i], slice[j] = slice[j], slice[i]

			i++
		}
	}

	slice[i], slice[high] = slice[high], slice[i]

	return i
}

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

func loadInput(inputPath string) ([]int, []int) {
	lefts := make([]int, 0)
	rights := make([]int, 0)

	input, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), "   ")

		left, err := strconv.Atoi(words[0])
		if err != nil {
			panic(err)
		}

		lefts = append(lefts, left)

		right, err := strconv.Atoi(words[1])
		if err != nil {
			panic(err)
		}

		rights = append(rights, right)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lefts, rights
}
