package main

import (
	"fmt"
	"strings"
)

func countValidWords(sentence string) int {
	words := strings.Fields(sentence)
	count := len(words)

	for _, v := range words {
		hyphen := 0
		flag := 1
		punct := 0
		for _, w := range v {
			if w >= '0' && w <= '9' {
				flag = 0
			} else if w == '-' {
				hyphen++
			} else if w == '!' || w == ',' || w == '.' {
				punct++
			}
		}

		if len(v) > 1 && (v[0] == '-' || v[len(v)-1] == '-' || v[0] == '!' || v[0] == '.' || v[0] == ',') {
			flag = 0
		}

		if flag == 0 || hyphen > 1 || punct > 1 {
			count--
		}

	}
	return count
}

func main() {
	fmt.Println(countValidWords("cat -and a d0g a-b. "))
}
