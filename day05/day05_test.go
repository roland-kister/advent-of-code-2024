package day05

import (
	"reflect"
	"strings"
	"testing"
)

const example = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestLoadInput(t *testing.T) {
	d := Day05{}
	d.LoadInput(strings.NewReader(example))

	page97 := &page{
		num:       97,
		preceding: []int{},
	}

	page75 := &page{
		num:       75,
		preceding: []int{97},
	}

	page47 := &page{
		num:       47,
		preceding: []int{75, 97},
	}

	page61 := &page{
		num:       61,
		preceding: []int{47, 75, 97},
	}

	page53 := &page{
		num:       53,
		preceding: []int{47, 61, 75, 97},
	}

	page29 := &page{
		num:       29,
		preceding: []int{47, 53, 61, 75, 97},
	}

	page13 := &page{
		num:       13,
		preceding: []int{29, 47, 53, 61, 75, 97},
	}

	pgMap := map[int]*page{
		13: page13,
		29: page29,
		47: page47,
		53: page53,
		61: page61,
		75: page75,
		97: page97,
	}

	ups := []update{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}

	if !reflect.DeepEqual(d.pgMap, pgMap) {
		t.Fatalf("LoadInput() -> pgMap %v, want %v", d.pgMap, pgMap)
	}

	if !reflect.DeepEqual(d.ups, ups) {
		t.Fatalf("LoadInput() -> ups %v, want %v", d.ups, ups)
	}
}

func TestPartOne(t *testing.T) {
	d := Day05{}
	d.LoadInput(strings.NewReader(example))

	out := 143
	actOut := d.PartOne()

	if out != actOut {
		t.Fatalf("PartOne() = %d, want %d", actOut, out)
	}
}

func TestPartTwo(t *testing.T) {
	d := Day05{}
	d.LoadInput(strings.NewReader(example))

	out := 123
	actOut := d.PartTwo()

	if out != actOut {
		t.Fatalf("PartTwo() = %d, want %d", actOut, out)
	}
}
