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
	registerA   int
	registerB   int
	registerC   int
	program     []byte
	instPointer int
	output      int
}

func (d *Day17) LoadInput(input io.Reader) {
	d.comp = &computer{}

	scanner := bufio.NewScanner(input)

	registerRe := regexp.MustCompile(`Register (?:A|B|C): (\d+)`)
	var err error

	scanner.Scan()
	aMatch := registerRe.FindStringSubmatch(scanner.Text())
	d.comp.registerA, err = strconv.Atoi(aMatch[1])
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	bMatch := registerRe.FindStringSubmatch(scanner.Text())
	d.comp.registerB, err = strconv.Atoi(bMatch[1])
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	cMatch := registerRe.FindStringSubmatch(scanner.Text())
	d.comp.registerC, err = strconv.Atoi(cMatch[1])
	if err != nil {
		panic(err)
	}

	scanner.Scan()

	programRe := regexp.MustCompile(`Program: ((?:\d,?)+)`)

	scanner.Scan()
	programMatch := programRe.FindStringSubmatch(scanner.Text())

	d.comp.program = make([]byte, 0)
	for _, numStr := range strings.Split(programMatch[1], ",") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		d.comp.program = append(d.comp.program, byte(num))
	}
}

func (d *Day17) PartOne() int {
	comp := d.comp.copy()
	comp.exec()

	octalStr := strconv.FormatInt(int64(comp.output), 8)
	res, _ := strconv.Atoi(octalStr)

	return res
}

func (d *Day17) PartTwo() int {
	input := 0
	for _, num := range d.comp.program {
		input *= 8
		input += int(num)
	}

	for a := 0; a < math.MaxInt; a++ {
		comp := d.comp.copy()
		comp.registerA = a
		comp.exec()

		if comp.output == input {
			return a
		}

		if a > 10_000 {
			return -1
		}
	}

	return 0
}

func (c *computer) copy() *computer {
	newC := &computer{
		registerA:   c.registerA,
		registerB:   c.registerB,
		registerC:   c.registerC,
		program:     make([]byte, len(c.program)),
		instPointer: 0,
		output:      0,
	}

	copy(newC.program, c.program)

	return newC
}

func (c *computer) exec() {
	for c.instPointer < len(c.program) {

		jump := c.program[c.instPointer] != 3

		switch c.program[c.instPointer] {
		case 0:
			c.adv()
		case 1:
			c.bxl()
		case 2:
			c.bst()
		case 3:
			c.jnz()
		case 4:
			c.bxc()
		case 5:
			c.out()
		case 6:
			c.bdv()
		case 7:
			c.cdv()
		}

		if jump {
			c.instPointer += 2
		}
	}
}

func (c *computer) adv() {
	num := 1 << c.comboOp()

	c.registerA /= num
}

func (c *computer) bxl() {
	num := c.literalOp()

	c.registerB ^= num
}

func (c *computer) bst() {
	num := c.comboOp()

	c.registerB = num % 8
}

func (c *computer) jnz() {
	if c.registerA == 0 {
		c.instPointer += 2
		return
	}

	num := c.literalOp()

	c.instPointer = num
}

func (c *computer) bxc() {
	c.registerB ^= c.registerC
}

func (c *computer) out() {
	num := c.comboOp()

	c.output *= 8
	c.output += num % 8
}

func (c *computer) bdv() {
	num := 1 << c.comboOp()

	c.registerB = c.registerA / num
}

func (c *computer) cdv() {
	num := 1 << c.comboOp()

	c.registerC = c.registerA / num
}

func (c *computer) comboOp() int {
	switch c.program[c.instPointer+1] {
	case 0, 1, 2, 3:
		return int(c.program[c.instPointer+1])
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

func (c *computer) literalOp() int {
	return int(c.program[c.instPointer+1])
}
