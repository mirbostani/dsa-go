package main

import "fmt"

func bubbleSortOptimizedAsc(arr []int) {
	n := len(arr) // array size
	isSwapped := true

	i := n - 1
	for isSwapped {
		isSwapped = false

		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // swapped
				isSwapped = true
			}
		}

		i--
	}
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	bubbleSortOptimizedAsc(arr)
	fmt.Println(arr)
}
