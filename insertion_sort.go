package main

import "fmt"

// сортировака втсавками O(n**2)
func insertionSort(arr []int) {
	n := len(arr)

}

func main() {
	arr := []int{3, 7, 9, 44, 11, 7, 999, 2}
	insertionSort(arr)

	fmt.Println(arr)
}
