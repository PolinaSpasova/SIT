package main

import "fmt"

func buddyStrings(s, goal string) bool {
	ind := make([]int, 0, 20)
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		for i := 1; i < len(s); i++ {
			if s[i-1] != s[i] {
				return false
			}
		}
		return true
	}
	for i := 0; i < len(s); i++ {
		if s[i] == goal[i] {
			continue
		}
		ind = append(ind, i)
	}

	if len(ind) == 2 && s[ind[0]] == goal[ind[1]] && goal[ind[0]] == s[ind[1]] {
		return true
	} else {
		return false
	}

}

func main() {
	fmt.Println(buddyStrings("cbad", "abcd"))
}
