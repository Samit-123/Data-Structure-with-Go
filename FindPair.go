package main

import (
	"fmt"
	"sort"
)

func FindPair(data []int, value int) bool {
	size := len(data)
	ret := false
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if (data[i] + data[j]) == value {
				fmt.Print(data[i], " ")
				fmt.Print(data[j])
				ret = true
			}
		}
	}
	return ret
}
func FindPair1(data []int, value int) bool {
	size := len(data)
	l := 0
	r := size - 1
	sort.Ints(data)
	ret := false
	for l < r {
		curr := data[l] + data[r]
		if curr == value {
			fmt.Println("The pair is :", "(", data[l], ",", data[r], ")")
			ret = true
		}
		if curr < value {
			l++
		} else {
			r--
		}
	}
	return ret
}
func FindPair3(data []int, value int) bool {
	s := make(Set)
	size := len(data)
	ret := false
	for i := 0; i < size; i++ {
		if s.Find(value - data[i]) {
			fmt.Println(i, "The pair is:", data[i], ",", (value - data[i]))
			ret = true
		} else {
			s.Add(data[i])
		}
	}
	return ret
}

func main() {
	data := []int{1, 3, 5, 8}
	sum := 8
	fmt.Println(FindPair3(data, sum))
}
