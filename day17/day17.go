// https://adventofcode.com/2024/day/17

package day17

import (
	"bufio"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Day17 struct {
	comp *computer
}

type computer struct {
	registerA   uint64
	registerB   uint64
	registerC   uint64
	program     []uint64
	instPointer uint64
	output      []uint64
}

func (d *Day17) LoadInput(input io.Reader) {
	d.comp = &computer{}

	scanner := bufio.NewScanner(input)

	registerRe := regexp.MustCompile(`Register (?:A|B|C): (\d+)`)
	var err error

	scanner.Scan()
	aMatch := registerRe.FindStringSubmatch(scanner.Text())
	d.comp.registerA, err = strconv.ParseUint(aMatch[1], 10, 64)
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	bMatch := registerRe.FindStringSubmatch(scanner.Text())
	d.comp.registerB, err = strconv.ParseUint(bMatch[1], 10, 64)
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	cMatch := registerRe.FindStringSubmatch(scanner.Text())
	d.comp.registerC, err = strconv.ParseUint(cMatch[1], 10, 64)
	if err != nil {
		panic(err)
	}

	scanner.Scan()

	programRe := regexp.MustCompile(`Program: ((?:\d,?)+)`)

	scanner.Scan()
	programMatch := programRe.FindStringSubmatch(scanner.Text())

	d.comp.program = make([]uint64, 0)
	for _, numStr := range strings.Split(programMatch[1], ",") {
		num, err := strconv.ParseUint(numStr, 10, 64)
		if err != nil {
			panic(err)
		}

		d.comp.program = append(d.comp.program, num)
	}
}

func (d *Day17) PartOne() int {
	comp := d.comp.copy()
	comp.exec(false)

	res := 0
	for _, num := range comp.output {
		res *= 10
		res += int(num)
	}

	return res
}

func (d *Day17) PartTwo() int {
	for a := uint64(164_540_892_147_389); a < math.MaxUint64; a++ {
		comp := d.comp.copy()
		comp.registerA = a

		if comp.exec(true) {
			return int(a)
		}
	}

	return 0
}

func (c *computer) copy() *computer {
	newC := &computer{
		registerA:   c.registerA,
		registerB:   c.registerB,
		registerC:   c.registerC,
		program:     make([]uint64, len(c.program)),
		instPointer: 0,
		output:      make([]uint64, 0),
	}

	copy(newC.program, c.program)

	return newC
}

func (c *computer) exec(partTwo bool) bool {
	progLen := uint64(len(c.program))

ExecLoop:
	for c.instPointer < progLen {
		switch c.program[c.instPointer] {
		case 0:
			c.registerA >>= c.comboOp()
		case 1:
			c.registerB ^= c.literalOp()
		case 2:
			c.registerB = c.comboOp() & 7
		case 3:
			if c.registerA != 0 {
				c.instPointer = c.literalOp()
				continue ExecLoop
			}
		case 4:
			c.registerB ^= c.registerC
		case 5:
			c.output = append(c.output, c.comboOp()&7)
			if partTwo && (len(c.output) > len(c.program) || c.program[len(c.output)-1] != c.output[len(c.output)-1]) {
				return false
			}

		case 6:
			c.registerB = c.registerA >> c.comboOp()
		case 7:
			c.registerC = c.registerA >> c.comboOp()
		}

		c.instPointer += 2
	}

	return true
}

func (c *computer) comboOp() uint64 {
	switch c.program[c.instPointer+1] {
	case 0, 1, 2, 3:
		return uint64(c.program[c.instPointer+1])
	case 4:
		return c.registerA
	case 5:
		return c.registerB
	case 6:
		return c.registerC
	default:
		panic("invalid program")
	}
}

func (c *computer) literalOp() uint64 {
	return uint64(c.program[c.instPointer+1])
}
