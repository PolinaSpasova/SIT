package main

import "fmt"

func isLongPressedName(name, typed string) bool {
	if len(typed) < len(name) || name[0] != typed[0] {
		return false
	}

	i, j := 1, 1
	for i < len(name) && j < len(typed) {
		if name[i] == typed[j] {
			i++
			j++
		} else if typed[j] == name[i-1] {
			j++
		} else {
			return false
		}
	}
	return j >= len(typed)

}

func main() {
	fmt.Println(isLongPressedName("alex", "aaleexxy"))
}
