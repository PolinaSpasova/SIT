package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	possible := 0
	if n == 0 {
		return true
	}

	if len(flowerbed) == 1 {
		return flowerbed[0] == 0
	}

	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		possible++
		flowerbed[0] = 1
	}

	for i := 1; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			if i == (len(flowerbed)-1) && flowerbed[i-1] == 0 {
				possible++
				flowerbed[i] = 1
			} else if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				possible++
				flowerbed[i] = 1
			}
		}
	}

	return possible >= n
}

func main() {
	flowerbed := []int{1, 0, 0, 1, 0}
	fmt.Println(canPlaceFlowers(flowerbed, 1))
}
