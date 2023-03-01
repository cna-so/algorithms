package main

import "fmt"

func main() {
	list := []int{1, 3, 5, 7, 9}
	res := binarySearch(list, 3)
	fmt.Println(res)
}

func binarySearch(arr []int, item int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0
}
