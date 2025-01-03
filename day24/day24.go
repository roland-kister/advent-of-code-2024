// https://adventofcode.com/2024/day/24

package day24

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"regexp"
	"sort"
	"strings"
)

type Day24 struct {
	wires map[string]uint8
	gates map[string]logicGate
}

type logicOp uint8

const (
	andOp logicOp = iota
	orOp
	xorOp
)

const undefined = math.MaxUint8

type logicGate struct {
	wires   *map[string]uint8
	wireA   string
	wireB   string
	wireOut string
	op      logicOp
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

		var op logicOp
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
	outputWires := d.getOutputWires()
	for _, wire := range outputWires {
		d.execute(wire)
	}

	return d.getNumber('z')
}

func (d *Day24) PartTwo() int {
	invalidGates := d.getInvalidWires()
	sort.Strings(invalidGates)

	res := strings.Join(invalidGates, ",")
	fmt.Printf("\tpart two = %s\n", res)

	return 0
}

func (d *Day24) getOutputWires() map[int]string {
	gates := make(map[int]string)

	for wire := range d.wires {
		if wire[0] != 'z' {
			continue
		}

		bitIndex := int(wire[1]-'0')*10 + int(wire[2]-'0')

		gates[bitIndex] = wire
	}

	return gates
}

func (d *Day24) execute(wireOut string) {
	gate := d.gates[wireOut]

	if gate.wireA[0] != 'x' && gate.wireA[0] != 'y' {
		d.execute(gate.wireA)
	}

	if gate.wireB[0] != 'x' && gate.wireB[0] != 'y' {
		d.execute(gate.wireB)
	}

	gate.combine()
}

func (d *Day24) getNumber(char byte) int {
	num := 0

	for wire, val := range d.wires {
		if wire[0] != char {
			continue
		}

		bitIndex := int(wire[1]-'0')*10 + int(wire[2]-'0')

		num |= int(val) << bitIndex
	}

	return num
}

func (d *Day24) getInvalidWires() []string {
	lastZ := len(d.getOutputWires()) - 1

	invalidWires := make([]string, 0)

	for _, gate := range d.gates {
		invalid, isValid := d.isValidOutput(gate, lastZ)
		if !isValid {
			invalidWires = append(invalidWires, invalid)
			continue
		}

		invalid, isValid = d.isValidAnd(gate)
		if !isValid {
			invalidWires = append(invalidWires, invalid)
			continue
		}

		invalid, isValid = d.isValidXor(gate)
		if !isValid {
			invalidWires = append(invalidWires, invalid)
			continue
		}

		invalid, isValid = d.isValidOr(gate)
		if !isValid {
			invalidWires = append(invalidWires, invalid)
			continue
		}
	}

	return invalidWires
}

func (d *Day24) isValidOutput(gate logicGate, lastZ int) (string, bool) {
	if gate.wireOut[0] != 'z' {
		return "", true
	}

	zIndex := int(gate.wireOut[1]-'0')*10 + int(gate.wireOut[2]-'0')
	if zIndex != lastZ && gate.op != xorOp {
		return gate.wireOut, false
	}

	return "", true
}

func (d *Day24) isValidAnd(gate logicGate) (string, bool) {
	if gate.op != andOp {
		return "", true
	}

	prevGate, ok := d.gates[gate.wireA]
	if ok && prevGate.op == andOp {
		xIndexA := int(prevGate.wireA[1]-'0')*10 + int(prevGate.wireA[2]-'0')
		xIndexB := int(prevGate.wireB[1]-'0')*10 + int(prevGate.wireB[2]-'0')

		if xIndexA != 0 && xIndexB != 0 {
			return gate.wireA, false
		}
	}

	prevGate, ok = d.gates[gate.wireB]
	if ok && prevGate.op == andOp {
		xIndexA := int(prevGate.wireA[1]-'0')*10 + int(prevGate.wireA[2]-'0')
		xIndexB := int(prevGate.wireB[1]-'0')*10 + int(prevGate.wireB[2]-'0')

		if xIndexA != 0 && xIndexB != 0 {
			return gate.wireB, false
		}
	}

	return "", true
}

func (d *Day24) isValidOr(gate logicGate) (string, bool) {
	if gate.op != orOp {
		return "", true
	}

	prevGate, ok := d.gates[gate.wireA]
	if ok && prevGate.op == xorOp {
		_, valid := d.isValidXor(prevGate)
		if valid {
			return gate.wireA, false
		}
	}

	prevGate, ok = d.gates[gate.wireB]
	if ok && prevGate.op == xorOp {
		_, valid := d.isValidXor(prevGate)
		if valid {
			return gate.wireB, false
		}
	}

	return "", true
}

func (d *Day24) isValidXor(gate logicGate) (string, bool) {
	if gate.op != xorOp {
		return "", true
	}

	if gate.wireA[0] != 'x' && gate.wireA[0] != 'y' && gate.wireB[0] != 'x' && gate.wireB[0] != 'y' && gate.wireOut[0] != 'z' {
		return gate.wireOut, false
	}

	return "", true
}
