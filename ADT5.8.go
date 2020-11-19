package main

import "fmt"

func printRepeating4(data []int, intrange int) {
	size := len(data)
	count := make([]int, intrange)
	for i := 0; i < size; i++ {
		if count[data[i]] == 1 {
			fmt.Println(" ", data[i])
		} else {
			count[data[i]]++
		}
	}
	fmt.Println()
}

func main() {
	data := []int{1, 2, 3, 4, 5, 3, 2, 1}
	intrange := 6
	printRepeating4(data, intrange)
}
