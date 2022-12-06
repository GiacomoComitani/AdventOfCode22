package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(readFile())
}

// read the input file and return the final result
func readFile() int {
	var sum int
	file, _ := os.Open("day3.txt")
	scanner := bufio.NewScanner(file)
	var s []string
	var count int
	for scanner.Scan() {
		count++
		s = append(s, scanner.Text())
		if count == 3 {
			sum += findCommon(s)
			s = []string{}
			count = 0
		}
	}
	return sum
}

// find the common value and call findValue
func findCommon(s []string) int {
	var count int
	mappa := make(map[byte]int)
	for _, str := range s {
		for i := 0; i < len(str); i++ {
			if mappa[str[i]] == count {
				mappa[str[i]]++
			}
			if mappa[str[i]] == 3 {
				return findValue(str[i])
			}
		}
		count++
	}
	return 0
}

// return the priority value of the common letter
func findValue(letter byte) int {
	if letter < 'a' {
		return int(letter-'A') + 27
	} else {
		return int(letter-'a') + 1
	}
}
