package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	m := readFile()
	fmt.Println(visible(m))
}

func readFile() [99][99]int {
	matrx := [99][99]int{}
	file, _ := os.Open("day8.txt")
	scanner := bufio.NewScanner(file)
	var i int
	for scanner.Scan() {
		for j := 0; j < len(scanner.Text()); j++ {
			num, _ := strconv.Atoi(string(scanner.Text()[j]))
			matrx[i][j] = num
		}
		i++
	}
	return matrx
}

func visible(m [99][99]int) int {
	var sum int
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			c := isVisible(m, m[i][j], i, j)
			if c > sum {
				sum = c
			}
		}
	}
	return sum
}

func isVisible(m [99][99]int, alb int, i, j int) int {
	var c1, c2, c3, c4 int
	for k := i - 1; k >= 0; k-- {
		c1++
		if m[k][j] >= alb {
			break
		}
	}
	for k := i + 1; k < 99; k++ {
		c2++
		if m[k][j] >= alb {
			break
		}
	}
	for k := j - 1; k >= 0; k-- {
		c3++
		if m[i][k] >= alb {
			break
		}
	}
	for k := j + 1; k < 99; k++ {
		c4++
		if m[i][k] >= alb {
			break
		}
	}
	return c1 * c2 * c3 * c4
}
