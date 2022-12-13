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
			if isVisible(m, m[i][j], i, j) {
				sum++
			}
		}
	}
	return sum
}

// First Part
func isVisible(m [99][99]int, alb int, i, j int) bool {
	if i == 0 || i == 98 || j == 0 || j == 98 {
		return true
	}
	for k := i - 1; k >= 0; k-- {
		if m[k][j] >= alb {
			break
		}
		if k == 0 {
			return true
		}
	}
	for k := i + 1; k < 99; k++ {
		if m[k][j] >= alb {
			break
		}
		if k == 98 {
			return true
		}
	}
	for k := j - 1; k >= 0; k-- {
		if m[i][k] >= alb {
			break
		}
		if k == 0 {
			return true
		}
	}
	for k := j + 1; k < 99; k++ {
		if m[i][k] >= alb {
			break
		}
		if k == 98 {
			return true
		}
	}
	return false
}
