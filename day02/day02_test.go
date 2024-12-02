package day02

import (
	"sync"
	"testing"
)

type InOutPair struct {
	in  []int
	out bool
}

func TestPartOne(t *testing.T) {
	pairs := []InOutPair{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, pair := range pairs {
		actOut := isRepSafe(pair.in)

		if pair.out != actOut {
			t.Fatalf("isRepSafe(%v) = %t, want %t", pair.in, actOut, pair.out)
		}
	}
}

func TestPartTwo(t *testing.T) {
	pairs := []InOutPair{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, true},
		{[]int{8, 6, 4, 4, 1}, true},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, pair := range pairs {
		actOutChan := make(chan bool, 1)
		wg := new(sync.WaitGroup)

		wg.Add(1)

		go isRepSafeTol(pair.in, actOutChan, wg)

		wg.Wait()

		actOut := <-actOutChan

		if pair.out != actOut {
			t.Fatalf("isRepSafeTol(%v, resChan, wg) = %t, want %t", pair.in, actOut, pair.out)
		}
	}
}
