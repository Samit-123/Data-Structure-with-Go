package main

import "fmt"

func getMax3(data []int, dataRange int) int {
	max := data[0]
	maxCount := 1
	size := len(data)
	count := make([]int, dataRange)
	for i := 0; i < size; i++ {
		count[data[i]]++
		if count[data[i]] > maxCount {
			maxCount = count[data[i]]
			max = data[i]
		}
	}
	return max
}
func main() {
	data := []int{2, 3, 4, 5, 1, 2, 6, 7, 8, 2}
	dataRange := 9
	fmt.Println(getMax3(data, dataRange))
}
