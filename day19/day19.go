// https://adventofcode.com/2024/day/19

package day19

import (
	"bufio"
	"io"
	"regexp"
	"slices"
	"strings"
)

type Day19 struct {
	patterns []string
	designs  []string
}

func (d *Day19) LoadInput(input io.Reader) {
	d.designs = []string{}

	scanner := bufio.NewScanner(input)

	scanner.Scan()
	d.patterns = strings.Split(scanner.Text(), ", ")
	scanner.Scan()

	for scanner.Scan() {
		d.designs = append(d.designs, scanner.Text())
	}

	slices.SortFunc(d.patterns, func(a, b string) int {
		return len(a) - len(b)
	})
	slices.Reverse(d.patterns)
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
	return 0
}
