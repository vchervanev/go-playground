package main

import (
	"fmt"
	"math"
)

var nums = []int{12, 345, 2, 6, 7896}

func findNumbers(nums []int) (result int) {
	for _, value := range nums {
		if int(math.Log10(float64(value)))%2 != 0 {
			fmt.Println(value, math.Log10(float64(value)), int(math.Log10(float64(value)))%2)
			result++
		}
	}
	return
}
func main() {
	fmt.Println(findNumbers(nums))
}
