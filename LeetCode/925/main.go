package main

import "fmt"

func isLongPressedName(name, typed string) bool {
	if len(typed) < len(name) || name[0] != typed[0] {
		return false
	}

	i, j := 1, 1
	for j < len(typed) {
		if name[i] == typed[j] && i < len(name)-1 {
			i++
			j++
		} else if typed[j] == name[i-1] || (name[i] == typed[j] && i == len(name)-1) {
			j++
		} else {
			return false
		}
	}
	return true

}

func main() {
	fmt.Println(isLongPressedName("alex", "aaleexxx"))
}
