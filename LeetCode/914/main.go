package main

import "fmt"

func hasGroupsSizeX(deck []int) bool {
	data := make([]int, 10000)
	temp := 0
	for _, v := range deck {
		data[v]++
	}
	for _, v := range data {
		if v != 0 {
			temp = v
			break
		}
	}
	for _, v := range data {
		if v != 0 && v != temp {
			return false
		}else if v!=0 && v<2{
			return false
		}
	}
	return true

}

func main() {
	deck:=[]int{1,2,3,4,4,3,2,1}
	fmt.Println(hasGroupsSizeX(deck))
}