package main

import (
	"sort"
	"testing"
)

func TestSortList(t *testing.T) {
	var generateList = func(length int) (list []int) {
		for i := 0; i < length; i++ {
			list = append([]int{i}, list...)
		}

		return
	}

	tests := map[string]struct {
		ListToOrder  []int
		ExpectedList []int
	}{
		"sort list with 10 values": {
			ListToOrder:  []int{9, 5, 2, 1, 6, 4, 7, 3, 8},
			ExpectedList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		"sort list with 10k": {
			ListToOrder: generateList(1e4),
			ExpectedList: func() []int {
				list := generateList(1e4)
				sort.Ints(list)
				return list
			}(),
		},
		"sort list with 50k": {
			ListToOrder: generateList(5e4),
			ExpectedList: func() []int {
				list := generateList(5e5)
				sort.Ints(list)
				return list
			}(),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			list := SortList(test.ListToOrder)

			for i := 0; i < len(list); i++ {
				if list[i] != test.ExpectedList[i] {
					t.Errorf("error on %s, expected: %d but got: %d", name, test.ExpectedList[i], list[i])
				}
			}
		})
	}
}
