package main

import "fmt"

func main() {
	list := []int{9, 5, 2, 1, 6, 4, 7, 3, 8}

	fmt.Println(SortList(list))
}

// i know exist the others types of sort list, example bubbleSort
// but this is for me exercise my mind
func SortList(list []int) []int {
	var copyList, result = list, []int{}

	for len(result) != len(list) {
		var min = copyList[0]
		var index int

		for i, v := range copyList {
			if min > v {
				min = v
				index = i
			}
		}

		result = append(result, min)
		copyList = append(copyList[:index], copyList[index+1:]...)
	}

	return result
}
