package main

import "fmt"

func canBeIncreasing(nums []int) bool {
	br := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			br++
		}
	}
	if br > 1 {
		return false
	}
	return true
}

func main() {
	nums := []int{1, 2, 10, 5, 7}
	fmt.Println(canBeIncreasing(nums))
}
