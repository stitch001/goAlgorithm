package search

import "fmt"

func binSearch(arr []int, target int) int {
	if len(arr) <= 0 {
		return -1
	}
	start := 0
	end := len(arr) - 1
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1

}

func Run() {
	arr := []int{1, 2, 3, 6, 7, 9, 10, 12, 15, 19, 24, 28}
	fmt.Println(binSearch(arr, 6))
	fmt.Println(binSearch(arr, 1))
	fmt.Println(binSearch(arr, 28))
	fmt.Println(binSearch(arr, -1))
	fmt.Println(binSearch(arr, 99))
	fmt.Println(binSearch(arr, 24))
}
