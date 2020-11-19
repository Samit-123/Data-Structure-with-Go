package main

import "fmt"

func getMax(data []int) int {
	size := len(data)
	max := data[0]
	count := 1
	maxCount := 1
	for i := 0; i < size; i++ {
		count = 1
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				count++
			}
		}
		if count > maxCount {
			max = data[i]
			maxCount = count
		}
	}
	return max
}
func main() {
	data := []int{3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6}
	fmt.Println(getMax(data))
}
