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

func readFile() int {
	file, _ := os.Open("day10.txt")
	scanner := bufio.NewScanner(file)
	var count, sum int
	var x int = 1
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		count++
		if count == 20 || count == 60 || count == 100 || count == 140 || count == 180 || count == 220 {
			sum += count * x
		}
		if args[0] == "addx" {
			count++
			if count == 20 || count == 60 || count == 100 || count == 140 || count == 180 || count == 220 {
				sum += count * x
			}
			num, _ := strconv.Atoi(args[1])
			x += num
		}
	}
	return sum
}
