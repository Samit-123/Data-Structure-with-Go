package main

import "fmt"

func findMissingNumber(data []int) (int, bool) {
	var found int
	size := len(data)
	for i := 1; i <= size; i++ {
		found = 0
		for j := 0; j < size; j++ {
			if data[j] == i {
				found = 1
				break
			}
		}
		if found == 0 {
			return i, true
		}
	}
	fmt.Println("No Number Missing::")
	return 0, false
}
func main() {
	data := []int{1, 2, 3, 4}
	res, bool1 := findMissingNumber(data)
	fmt.Println(res)
	fmt.Println(bool1)
}
