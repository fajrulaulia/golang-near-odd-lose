package axiatadigitaltest

import (
	"sort"
)

type Algo struct {
}

func FindOddLost(N []int) int {
	var c Algo
	sort.Ints(sort.IntSlice(N))
	N = c.arrayDistinct(N)
	return c.getOddNearestPositiveNumber(N)
}

func (c *Algo) arrayDistinct(arr []int) []int {
	lists := make(map[int]bool)
	newList := []int{}

	for _, entry := range arr {
		if !lists[entry] {
			lists[entry] = true
			newList = append(newList, entry)
		}
	}
	return newList
}

func (c *Algo) getOddNearestPositiveNumber(arr []int) int {
	var start int
	isFirstPositiveFetch := true
	for _, entry := range arr {
		if entry > 0 {
			if isFirstPositiveFetch {
				start = entry
				isFirstPositiveFetch = false
			} else {
				start += 2
			}

			if start != entry {
				return start
			}
		}
	}

	if arr[len(arr)-1] > 0 {
		return arr[len(arr)-1] + 2
	}

	return 1
}
