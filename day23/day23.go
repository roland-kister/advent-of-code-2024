// https://adventofcode.com/2024/day/23

package day23

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"sort"
	"strings"
)

type Day23 struct {
	comps map[int]*computer
}

const (
	shift = 5
	mask  = 0b11111
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
				connections: make(map[int]*computer, 0),
			}
			d.comps[compOneId] = compOne
		}

		compTwo, ok := d.comps[compTwoId]
		if !ok {
			compTwo = &computer{
				id:          compTwoId,
				connections: make(map[int]*computer, 0),
			}
			d.comps[compTwoId] = compTwo
		}

		compOne.connections[compTwoId] = compTwo
		compTwo.connections[compOneId] = compOne
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
				if (networkIdCopy&mask)+'a' == 't' {
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
	longestOverlap := make(map[int]*computer)

	for _, comp := range d.comps {
		for _, conn := range comp.connections {
			overlap := comp.getConnOverlap(conn)

			if len(overlap) <= len(longestOverlap) {
				continue
			}

			if checkOverlap(overlap) {
				longestOverlap = overlap
			}
		}
	}

	idsStr := make([]string, 0)
	for id := range longestOverlap {
		idStr := fmt.Sprintf("%c%c", ((id>>shift)&mask)+'a', (id&mask)+'a')

		idsStr = append(idsStr, idStr)
	}

	sort.Strings(idsStr)

	res := strings.Join(idsStr, ",")

	fmt.Printf("\tpart two = %s\n", res)

	return 0
}

type computer struct {
	id          int
	connections map[int]*computer
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

func (c *computer) getConnOverlap(another *computer) map[int]*computer {
	overlap := make(map[int]*computer)
	overlap[c.id] = c
	overlap[another.id] = another

MainLoop:
	for _, aConn := range c.connections {
		for _, bConn := range another.connections {
			if aConn.id == bConn.id {
				overlap[aConn.id] = aConn
				continue MainLoop
			}
		}
	}

	return overlap
}

func checkOverlap(overlap map[int]*computer) bool {
	for aId, aComp := range overlap {
		for bId := range overlap {
			if aId == bId {
				continue
			}

			_, ok := aComp.connections[bId]
			if !ok {
				return false
			}
		}
	}

	return true
}
