package main

import "fmt"

// ПУЗЫРЬКОВАЯ СОРТИРОВКА. O(n**2)
func bubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func main() {
	unsorted := []int{3, 7, 9, 44, 11, 7, 999, 2}
	sorted := bubbleSort(unsorted)
	fmt.Println(sorted)
}
