package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	(readFile())
}

func readFile() {
	file, _ := os.Open("day10.txt")
	scanner := bufio.NewScanner(file)
	var count int
	var x int = 1
	if count == x || count == x-1 || count == x+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		count++
		count = count % 40
		if count == x || count == x-1 || count == x+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if count == 39 {
			fmt.Println()
		}
		if args[0] == "addx" {
			count++
			num, _ := strconv.Atoi(args[1])
			x += num
			count = count % 40
			if count == x || count == x-1 || count == x+1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			if count == 39 {
				fmt.Println()
			}
		}
	}
}
