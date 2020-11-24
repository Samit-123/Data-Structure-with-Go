package main

import "fmt"

func getMajority(data []int) (int, bool) {
	size := len(data)
	max := 0
	count := 0
	maxCount := 0
	for i := 0; i < size; i++ {
		count = 1
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			max = data[i]
		}
	}
	if maxCount > size/2 {
		return max, true
	}
	fmt.Println("Majority Element does not Exist::")
	return 0, false
}
func main() {
	data := []int{1, 1, 2, 2}
	result, bool1 := getMajority(data)
	fmt.Println(result)
	fmt.Println(bool1)
}
