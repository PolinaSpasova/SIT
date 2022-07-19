package main

import "fmt"

func canBeIncreasing(nums []int) bool {
	br := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			br++
		}
	}
	if br > 1 {
		return false
	} else {
		return true
	}
}

func main() {
	nums := []int{2, 3, 5, 10, 7, 9}
	fmt.Println(canBeIncreasing(nums))
}
