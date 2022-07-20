package main

import "fmt"

func validMountainArray(arr []int) bool {
	up := true
	if len(arr) < 3 || arr[1] <= arr[0] {
		return false
	}

	for i := 2; i < len(arr); i++ {
		if arr[i-1] == arr[i] {
			return false
		} else if arr[i-1] > arr[i] && up {
			up = false
		} else if arr[i-1] < arr[i] && !up {
			return false
		}
	}
	return true

}

func main() {
	array := []int{0, 3, 2, 1, 5}
	fmt.Println(validMountainArray(array))
}
