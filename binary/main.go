package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var total = 1000

	randomNumber := rand.Intn(total-1) + 1

	var list = make([]int, total)

	for i := 0; i <= total; i++ {
		list = append(list, i)
	}

	countTry := BinarySearch(randomNumber, list)

	fmt.Printf("number to find is %d on total of %d, is used %d to find", randomNumber, total, countTry)

}

func BinarySearch(toFind int, list []int) int {
	var count int

	var min, max int = 0, len(list) - 1

	for {
		middle := (min + max) / 2

		kick := list[middle]

		if kick == toFind {
			return count
		}
		if kick > toFind {
			max = middle - 1
		} else {
			min = middle + 1
		}

		count++
	}
}
