package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(readFile())
}

// read the input file and returns the numbers of assignment pairs in which one range fully contain the other
func readFile() int {
	var sum int
	file, _ := os.Open("day4.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), ",")
		elve_1 := strings.Split(args[0], "-")
		elve_2 := strings.Split(args[1], "-")
		if overlap(convert(elve_1, elve_2)) {
			fmt.Println(args)
			sum++
		}
	}
	return sum
}

// return true if one range fully contain the other, false if not
func isContained(n1, n2, n3, n4 int) bool {
	if (n1 >= n3 && n2 <= n4) || (n1 <= n3 && n2 >= n4) {
		return true
	}
	return false
}

// return true if the range overlaps
func overlap(n1, n2, n3, n4 int) bool {
	if (n2 < n3 && n1 < n4) || (n2 > n3 && n1 > n4) {
		return false
	}
	return true
}

func convert(s, s2 []string) (int, int, int, int) {
	n1, _ := strconv.Atoi(s[0])
	n2, _ := strconv.Atoi(s[1])
	n3, _ := strconv.Atoi(s2[0])
	n4, _ := strconv.Atoi(s2[1])
	return n1, n2, n3, n4
}
