package main

import (
	"math/rand"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := map[string]struct {
		Total  int
		ToFind int
	}{
		"should search on 10 numbers": {
			ToFind: rand.Intn(10-1) + 1,
			Total:  10,
		},
		"should search on 100 numbers": {
			ToFind: rand.Intn(100-1) + 1,
			Total:  100,
		},
		"should search on 1k numbers": {
			ToFind: rand.Intn(1000-1) + 1,
			Total:  1000,
		},
		"should search on 10k numbers": {
			ToFind: rand.Intn(10000-1) + 1,
			Total:  10000,
		},
		"should search on 100k numbers": {
			ToFind: rand.Intn(100000-1) + 1,
			Total:  100000,
		},
		"should search on 1M numbers": {
			ToFind: rand.Intn(1000000-1) + 1,
			Total:  1000000,
		},
		"should search on 1B numbers": {
			ToFind: rand.Intn(10000000-1) + 1,
			Total:  10000000,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var list []int

			for i := 0; i < test.Total; i++ {
				list = append(list, i)
			}

			var maxTry, totalLocal int = 0, test.Total

			for {
				if totalLocal == 1 {
					break
				}
				totalLocal = totalLocal / 2
				maxTry++
			}

			result := BinarySearch(test.ToFind, list)

			if result > maxTry {
				t.Errorf("expected max %d but got %d", maxTry, result)
			}
		})
	}
}
