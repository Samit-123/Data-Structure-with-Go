package main

import (
	"fmt"
	"sort"
)

func third_Approach(data1 []int, data2 []int, value int) bool {
	size1 := len(data1)
	size2 := len(data2)
	sort.Ints(data1)
	sort.Ints(data2)
	ret := false
	low := 0
	high := size2 - 1
	for low < size1 && high >= 0 {
		curr := data1[low] + data2[high]
		if curr == value {
			fmt.Println("The Pair is : (", data1[low], ",", data2[high], ")")
			ret = true
		}
		if curr < value {
			low++
		} else {
			high--
		}
	}
	return ret
}
func main() {
	data1 := []int{10, 20, 3, 6, 8, 11}
	data2 := []int{12, 13, 17, 19, 5}
	value := 15
	fmt.Println(third_Approach(data1, data2, value))
}
