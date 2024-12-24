// https://adventofcode.com/2024/day/19

package day19

import (
	"bufio"
	"io"
	"math"
	"regexp"
	"strings"
)

type Day19 struct {
	patterns      []string
	patternsSet   map[string]bool
	maxPatternLen int
	minPatternLen int
	patternsCache map[string]int
	designs       []string
}

func (d *Day19) LoadInput(input io.Reader) {
	d.designs = []string{}
	d.patternsSet = map[string]bool{}
	d.patternsCache = map[string]int{}
	d.maxPatternLen = 0
	d.minPatternLen = math.MaxInt

	scanner := bufio.NewScanner(input)

	scanner.Scan()
	d.patterns = strings.Split(scanner.Text(), ", ")
	scanner.Scan()

	for scanner.Scan() {
		d.designs = append(d.designs, scanner.Text())
	}

	for _, pattern := range d.patterns {
		d.patternsSet[pattern] = true
		if len(pattern) > d.maxPatternLen {
			d.maxPatternLen = len(pattern)
		}

		if len(pattern) < d.minPatternLen {
			d.minPatternLen = len(pattern)
		}
	}
}

func (d *Day19) PartOne() int {
	re := regexp.MustCompile("^(?:" + strings.Join(d.patterns, "|") + ")+$")

	possibleDesigns := []string{}

	for _, design := range d.designs {
		if re.MatchString(design) {
			possibleDesigns = append(possibleDesigns, design)
		}
	}

	d.designs = possibleDesigns

	return len(d.designs)
}

func (d *Day19) PartTwo() int {
	total := 0

	for _, design := range d.designs {
		total += d.countCombinations(design)
	}

	return total
}

func (d *Day19) countCombinations(design string) int {
	if len(design) == 0 {
		return 1
	}

	cached, ok := d.patternsCache[design]
	if ok {
		return cached
	}

	total := 0

	maxLen := len(design)
	if maxLen > d.maxPatternLen {
		maxLen = d.maxPatternLen
	}

	for i := maxLen; i >= d.minPatternLen; i-- {
		_, ok = d.patternsSet[design[:i]]
		if !ok {
			continue
		}

		total += d.countCombinations(design[i:])
	}

	d.patternsCache[design] = total
	return total
}
