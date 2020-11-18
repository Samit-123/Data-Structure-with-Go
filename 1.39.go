package main

import "fmt"

func BinarySearchRecursive(data []int, low int, high int, value int) int {
	mid := low + (high-low)/2 
	if data[mid] == value {
		return mid
	} else if data[mid] < value {
		return BinarySearchRecursive(data, mid+1, high, value)
	} else {
		return BinarySearchRecursive(data, low, mid-1, value)
	}
}
func main() {
	data := []int{1, 2, 3, 4, 5, 6}
	n := len(data)
	fmt.Println(BinarySearchRecursive(data, 0, n, 4))
}
