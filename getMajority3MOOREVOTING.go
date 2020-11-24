package main

import "fmt"

func getMajority3MOOREVOTING(data []int) (int, bool) {
	size := len(data)
	majindex := 0
	count := 1
	for i := 1; i < size; i++ {
		if data[majindex] == data[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			majindex = i
			count = 1
		}
	}
	candidate := data[majindex]
	count := 0
	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		return data[candidate], true
	}
	fmt.Println("majority element doest not exist")
	return 0, false
}
func main() {
	data := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	result, bool1 := getMajority3MOOREVOTING(data)
	fmt.Println(result)
	fmt.Println(bool1)
}
