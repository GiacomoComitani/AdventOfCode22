package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(leggiFile())
}

func leggiFile() int {
	file, _ := os.Open("day1.txt")
	scanner := bufio.NewScanner(file)
	var max, max2, max3, sum int
	for scanner.Scan() {
		str, _ := strconv.Atoi(scanner.Text())
		sum += str
		if str == 0 {
			max, max2, max3 = maxs(sum, max, max2, max3)
			sum = 0
		}
	}
	max, max2, max3 = maxs(sum, max, max2, max3)
	return (max + max2 + max3)
}

func maxs(sum, max, max2, max3 int) (int, int, int) {
	if sum > max3 {
		max = max2
		max2 = max3
		max3 = sum
	} else if sum > max2 {
		max = max2
		max2 = sum
	} else if sum > max {
		max = sum
	}
	return max, max2, max3
}
