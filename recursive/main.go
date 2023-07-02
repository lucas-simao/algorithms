package main

import "fmt"

func main() {
	fmt.Println(sumDescending(5))
	fmt.Println(sumList([]int{1, 2, 3, 4, 5}))
	fmt.Println(countList([]int{1, 2, 3, 4, 5}))
}

func sumDescending(num int) int {
	if num == 0 {
		return num
	}

	return num + sumDescending(num-1)
}

func sumList(list []int) int {
	if len(list) == 0 {
		return 0
	}

	return list[0] + sumList(list[1:])
}

func countList(list []int) int {
	if len(list) == 0 {
		return 0
	}

	return 1 + countList(list[1:])
}
