package main

import "fmt"

func main() {

	//возведение в степень
	res := pow(4, 3)
	fmt.Println("res", res)
}

// pow возведение number в степень power
func pow(number, power int) int {
	n := number

	for i := 1; i < power; i++ {
		number *= n
	}

	return number
}
