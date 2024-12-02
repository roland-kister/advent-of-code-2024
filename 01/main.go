// https://adventofcode.com/2024/day/1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lefts := make([]int, 0)
	rights := make([]int, 0)

	loadInput("./input.txt", &lefts, &rights)

	quicksort(lefts, 0, len(lefts)-1)
	quicksort(rights, 0, len(rights)-1)

	diff := 0

	for i := 0; i < len(lefts); i++ {
		diff += abs(lefts[i] - rights[i])
	}

	fmt.Printf("Total distance (Part One): %d\n", diff)

	occurMapLen := lefts[len(lefts)-1]
	if lefts[len(lefts)-1] < rights[len(rights)-1] {
		occurMapLen = rights[len(rights)-1]
	}

	occurMapLen++

	occurMap := make([]int, occurMapLen)

	for _, right := range rights {
		occurMap[right]++
	}

	simScore := 0

	for _, left := range lefts {
		simScore += occurMap[left] * left
	}

	fmt.Printf("Similarity score (Part Two): %d\n", simScore)
}

func loadInput(inputPath string, lefts *[]int, rights *[]int) {
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

		*lefts = append(*lefts, left)

		right, err := strconv.Atoi(words[1])
		if err != nil {
			panic(err)
		}

		*rights = append(*rights, right)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
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
