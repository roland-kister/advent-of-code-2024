// https://adventofcode.com/2024/day/23

package day23

import (
	"bufio"
	"io"
	"slices"
	"strings"
)

type Day23 struct {
	comps map[int]*computer
}

const (
	shift      = 5
	firstMask  = 0b1111100000
	secondMask = 0b11111
)

func (d *Day23) LoadInput(input io.Reader) {
	d.comps = make(map[int]*computer)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		compNames := strings.Split(scanner.Text(), "-")

		compOneId := (int(compNames[0][0]-'a') << shift) | int(compNames[0][1]-'a')
		compTwoId := (int(compNames[1][0]-'a') << shift) | int(compNames[1][1]-'a')

		compOne, ok := d.comps[compOneId]
		if !ok {
			compOne = &computer{
				id:          compOneId,
				connections: make([]*computer, 0),
			}
			d.comps[compOneId] = compOne
		}

		compTwo, ok := d.comps[compTwoId]
		if !ok {
			compTwo = &computer{
				id:          compTwoId,
				connections: make([]*computer, 0),
			}
			d.comps[compTwoId] = compTwo
		}

		compOne.connections = append(compOne.connections, compTwo)
		compTwo.connections = append(compTwo.connections, compOne)
	}
}

func (d *Day23) PartOne() int {
	networkIdSet := make(map[int]bool)
	for _, comp := range d.comps {
		networkIds := comp.getNetworks(2)

	AppendLoop:
		for _, networkId := range networkIds {
			networkIdCopy := networkId >> shift

			for networkIdCopy != 0 {
				if (networkIdCopy&secondMask)+'a' == 't' {
					networkIdSet[networkId] = true
					continue AppendLoop
				}

				networkIdCopy >>= shift
				networkIdCopy >>= shift
			}

		}
	}

	return len(networkIdSet)
}

func (d *Day23) PartTwo() int {
	return 0
}

type computer struct {
	id          int
	connections []*computer
}

func (c *computer) getNetworks(connCount int) []int {
	networks := make([][]int, 0)

	for _, conn := range c.connections {
		subNetworks := conn.getSubNetworks(connCount-1, c.id)

		if len(subNetworks) == 0 {
			continue
		}

		for _, subNetwork := range subNetworks {
			networks = append(networks, append(subNetwork, c.id))
		}
	}

	networkIdSet := make(map[int]bool)
	for _, network := range networks {
		slices.Sort(network)

		networkId := 0
		for _, compId := range network {
			networkId <<= shift
			networkId <<= shift

			networkId |= compId
		}

		networkIdSet[networkId] = true
	}

	networkIds := make([]int, 0)
	for networkId := range networkIdSet {
		networkIds = append(networkIds, networkId)
	}

	return networkIds
}

func (c *computer) getSubNetworks(connCount, startCompId int) [][]int {
	if connCount == 0 {
		for _, conn := range c.connections {
			if conn.id == startCompId {
				return [][]int{{c.id}}
			}
		}

		return [][]int{}
	} else if c.id == startCompId {
		return [][]int{}
	}

	networks := make([][]int, 0)
	for _, conn := range c.connections {
		subNetworks := conn.getSubNetworks(connCount-1, startCompId)

		if len(subNetworks) == 0 {
			continue
		}

		for _, subNetwork := range subNetworks {
			networks = append(networks, append(subNetwork, c.id))
		}
	}

	return networks
}
