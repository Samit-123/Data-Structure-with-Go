package main

import (
	"fmt"
	"sort"
)

func getMax2(data []int) int {
	size := len(data)
	max := data[0]
	maxCount := 1
	curr := data[0]
	currCount := 1
	sort.Ints(data)
	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			currCount++
		} else {
			currCount = 1
			curr = data[i]
		}
		if currCount > maxCount {
			maxCount = currCount
			max = curr
		}
	}
	return max
}
func main() {
	data := []int{1, 2, 3, 3, 2, 1, 5, 6, 7, 8, 9, 1, 2, 4}
	fmt.Println(getMax2(data))
}
