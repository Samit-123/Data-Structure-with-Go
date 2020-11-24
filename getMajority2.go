package main

import (
	"fmt"
	"sort"
)

func getMajority2(data []int) (int, bool) {
	size := len(data)
	majIndex := size / 2
	sort.Ints(data)
	candidate := data[majIndex]
	count := 0
	for i := 0; i < size; i++ {
		if candidate == data[i] {
			count++
		}
	}
	if count > size/2 {
		return data[majIndex], true
	}
	fmt.Println("Major element does not exist::")
	return 0, false
}
func main() {
	data := []int{1, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3}
	result, bool1 := getMajority2(data)
	fmt.Println(result)
	fmt.Println(bool1)
}
