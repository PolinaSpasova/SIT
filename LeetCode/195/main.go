package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		if line == 9 {
			fmt.Println(scanner.Text())
		}
		line++
	}
	if line < 9 {
		fmt.Println(" There is no tenth line")
	}
}
