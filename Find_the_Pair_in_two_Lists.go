package main

import (
	"fmt"
	"sort"
)

func first_Approach(data1 []int, data2 []int, value int) bool {
	size1 := len(data1)
	size2 := len(data2)
	ret := false
	for i := 0; i < size1; i++ {
		x := (value - data1[i])
		for j := 0; j < size2; j++ {
			if x == data2[j] {
				fmt.Println("The pair is: (", data1[i], x, ")")
				ret = true
			}
		}
	}
	return ret
}
func BinarySearch(data []int, value int) bool {
	size := len(data)
	low := 0
	high := size - 1
	for low <= high {
		mid := low + (high-low)/2
		if data[mid] == value {
			return true
		} else if data[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
func second_Approach(data1 []int, data2 []int, value int) bool {
	size1 := len(data1)
	ret := false
	sort.Ints(data2)
	for i := 0; i < size1; i++ {
		x := value - data1[i]
		res2 := BinarySearch(data2, x)
		if res2 == true {
			fmt.Println("The pair is: (", data1[i], ",", x, ")")
			ret = true
		}
	}
	return ret
}
func third_Approach(data1 []int, data2 []int, value int) bool {
	size1 := len(data1)
	size2 := len(data2)
	ret := false
	sort.Ints(data1)
	sort.Ints(data2)
	low := 0
	high := size2 - 1
	ret := false
	for low < size1 && high >= 0 {
		curr := data1[low] + data2[high]
		if curr == value {
			fmt.Println("The pairs: (", data1[low], ",", data2[high], ")")
			ret = true
			low++
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
	fmt.Println(first_Approach(data1, data2, value))
	fmt.Println("========================")
	fmt.Println(second_Approach(data1, data2, value))
	fmt.Println("========================")
	fmt.Println(third_Approach(data1, data2, value))
}
