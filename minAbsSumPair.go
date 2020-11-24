package main

import (
	"fmt"
	"math"
)

func minAbsSumPair(data []int) {
	var sum int64
	size := len(data)
	if size < 2 {
		fmt.Println("Invalid Input")
	}
	minFirst := 0
	minSecond := 1
	minSum := int64(math.Abs(float64(data[0] + data[1])))
	for l := 0; l < size-1; l++ {
		for r := l + 1; r < size; r++ {
			sum = int64(math.Abs(float64(data[l] + data[r])))
			if sum < minSum {
				minSum = sum
				minFirst = l
				minSecond = r
			}
		}
	}
	fmt.Println("The two minimum elements are :(", data[minFirst], ",", data[minSecond], ")")
}
func main() {
	data := []int{-9, 8, 5, 8, -10, -7, 5, 4, 14, -7}
	minAbsSumPair(data)
}
