package main

import "fmt"

func thirdMax(nums []int) int {
	second := -1
	third := -1
	first := nums[0]
	for _, v := range nums {
		if v > first {
			first = v
		}
	}
	for _, v := range nums {
		if v > second && v < first {
			second = v
		}
	}
	for _, v := range nums {
		if v > third && v < second {
			third = v
		}
	}
	if third != -1 {
		return third
	} else {
		return first
	}
}

func main() {

	nums := []int{1, 3}
	fmt.Println(thirdMax(nums))
}
