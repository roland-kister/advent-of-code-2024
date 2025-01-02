// https://adventofcode.com/2024/day/24

package day24

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"
)

type Day24 struct {
	wires map[string]uint8
	gates map[string]logicGate
}

type logcOp uint8

const (
	andOp logcOp = iota
	orOp
	xorOp
)

const undefined = math.MaxUint8

type logicGate struct {
	wires   *map[string]uint8
	wireA   string
	wireB   string
	wireOut string
	op      logcOp
}

func (lg logicGate) combine() {
	var res uint8

	switch lg.op {
	case andOp:
		res = (*lg.wires)[lg.wireA] & (*lg.wires)[lg.wireB]
	case orOp:
		res = (*lg.wires)[lg.wireA] | (*lg.wires)[lg.wireB]
	case xorOp:
		res = (*lg.wires)[lg.wireA] ^ (*lg.wires)[lg.wireB]
	}

	if res != 0 {
		res = 1
	}

	(*lg.wires)[lg.wireOut] = res
}

func (d *Day24) LoadInput(input io.Reader) {
	d.wires = make(map[string]uint8)
	d.gates = make(map[string]logicGate)

	wireRe := regexp.MustCompile(`(\w\d+): (1|0)`)
	gateRe := regexp.MustCompile(`(\w+) (AND|OR|XOR) (\w+) -> (\w+)`)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}

		match := wireRe.FindStringSubmatch(text)

		wireName := match[1]
		wireVal := match[2][0] - '0'

		d.wires[wireName] = uint8(wireVal)
	}

	for scanner.Scan() {
		match := gateRe.FindStringSubmatch(scanner.Text())

		wireA := match[1]
		_, ok := d.wires[wireA]
		if !ok {
			d.wires[wireA] = undefined
		}

		wireB := match[3]
		_, ok = d.wires[wireB]
		if !ok {
			d.wires[wireB] = undefined
		}

		wireOut := match[4]
		_, ok = d.wires[wireOut]
		if !ok {
			d.wires[wireOut] = undefined
		}

		opStr := match[2]

		var op logcOp
		switch opStr {
		case "AND":
			op = andOp
		case "OR":
			op = orOp
		case "XOR":
			op = xorOp
		}

		d.gates[wireOut] = logicGate{
			wires:   &d.wires,
			wireA:   wireA,
			wireB:   wireB,
			wireOut: wireOut,
			op:      op,
		}
	}
}

func (d *Day24) PartOne() int {
	res := 0

	for wireOut := range d.gates {
		if wireOut[0] != 'z' {
			continue
		}

		outIndex, err := strconv.Atoi(wireOut[1:])
		if err != nil {
			panic(err)
		}

		d.execute(wireOut)

		res |= (int(d.wires[wireOut]) << outIndex)
	}

	return res
}

func (d *Day24) PartTwo() int {
	xVal := 0
	yVal := 0
	zReal := 0

	for wire, val := range d.wires {
		if wire[0] != 'x' && wire[0] != 'y' && wire[0] != 'z' {
			continue
		}

		zIndex, err := strconv.Atoi(wire[1:])
		if err != nil {
			panic(err)
		}

		shiftedVal := (int(val) << zIndex)

		if wire[0] == 'x' {
			xVal |= shiftedVal
		} else if wire[0] == 'y' {
			yVal |= shiftedVal
		} else {
			zReal |= shiftedVal
		}
	}

	zExpected := xVal + yVal

	fmt.Println(xVal)
	fmt.Println(yVal)
	fmt.Println(zReal)
	fmt.Println(zExpected)

	return 1
}

func (d *Day24) execute(wireOut string) {
	gate := d.gates[wireOut]

	wireAVal := d.wires[gate.wireA]
	if wireAVal == undefined {
		d.execute(gate.wireA)
	}

	wireBVal := d.wires[gate.wireB]
	if wireBVal == undefined {
		d.execute(gate.wireB)
	}

	gate.combine()
}
