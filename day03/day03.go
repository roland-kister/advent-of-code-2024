// https://adventofcode.com/2024/day/3

package day03

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

type Day03 struct {
	insts []string
}

func (d *Day03) LoadInput(input io.Reader) {
	d.insts = []string{}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		d.insts = append(d.insts, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func (d *Day03) PartOne() int {
	sum := 0

	for _, inst := range d.insts {
		instSum := 0

		re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

		for _, match := range re.FindAllStringSubmatch(inst, -1) {
			numA, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}

			numB, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}

			instSum += numA * numB
		}

		sum += instSum
	}

	return sum
}

func (d *Day03) PartTwo() int {
	instructs := make([]instruct, 0)

	for _, inst := range d.insts {
		instructs = append(instructs, parseInst(inst)...)
	}

	enabled := true
	sum := 0

	for _, instruct := range instructs {
		switch instruct.name {
		case do:
			enabled = true
		case dont:
			enabled = false
		case mul:
			if enabled {
				sum += instruct.args[0] * instruct.args[1]
			}
		}
	}

	return sum
}

func parseInst(inst string) []instruct {
	instructs := make([]instruct, 0)

	for i := 0; i < len(inst); {
		instruct, skip := parseDo(inst, i)
		if skip != 0 {
			instructs = append(instructs, instruct)
			i += skip
			continue
		}

		instruct, skip = parseDont(inst, i)
		if skip != 0 {
			instructs = append(instructs, instruct)
			i += skip
			continue
		}

		instruct, skip = parseMul(inst, i)
		if skip != 0 {
			instructs = append(instructs, instruct)
			i += skip
			continue
		}

		i++
	}

	return instructs
}

func parseDo(inst string, index int) (instruct, int) {
	if index+4 > len(inst) {
		return instruct{}, 0
	}

	if inst[index:index+4] == "do()" {
		return instruct{
			name: do,
		}, 4
	}

	return instruct{}, 0
}

func parseDont(inst string, index int) (instruct, int) {
	if index+7 > len(inst) {
		return instruct{}, 0
	}

	if inst[index:index+7] == "don't()" {
		return instruct{
			name: dont,
		}, 7
	}

	return instruct{}, 0
}

func parseMul(inst string, index int) (instruct, int) {
	if index+8 > len(inst) {
		return instruct{}, 0
	}

	if inst[index:index+4] != "mul(" {
		return instruct{}, 0
	}

	commaIndex, parIndex := -1, -1
	num := true

	for i := index + 4; i < len(inst); i++ {
		if inst[i] >= '0' && inst[i] <= '9' {
			num = false
			continue
		} else if num {
			break
		}

		if commaIndex < 0 && inst[i] == ',' {
			commaIndex = i
			num = true
			continue
		}

		if parIndex < 0 && inst[i] == ')' {
			parIndex = i
		}

		break
	}

	if commaIndex < 0 || parIndex < 0 {
		return instruct{}, 0
	}

	aStr := inst[index+4 : commaIndex]
	bStr := inst[commaIndex+1 : parIndex]

	aNum, err := strconv.Atoi(aStr)
	if err != nil {
		panic(err)
	}

	bNum, err := strconv.Atoi(bStr)
	if err != nil {
		panic(err)
	}

	return instruct{
		name: mul,
		args: []int{aNum, bNum},
	}, parIndex - index
}

type instructName int

const (
	do instructName = iota
	dont
	mul
)

type instruct struct {
	name instructName
	args []int
}
