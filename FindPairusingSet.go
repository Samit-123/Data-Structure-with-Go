package main

import "fmt"

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
	data := []int{3, 5, 7, 9, 11}
	sum := 14
	fmt.Println(FindPair3(data, sum))
}
