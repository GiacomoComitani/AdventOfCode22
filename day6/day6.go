package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(processed(readFile()))
}

func readFile() string {
	var s string
	file, _ := os.Open("day6.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = scanner.Text()
	}
	return s
}

func processed(s string) int {
	for i := 14; i < len(s); i++ {
		str := s[i-14 : i]
		if allDifferent(str) {
			return i
		}
	}
	return 0
}

func allDifferent(s string) bool {
	mappa := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		_, ok := mappa[s[i]]
		if ok == true {
			return false
		} else {
			mappa[s[i]]++
		}
	}
	return true
}
