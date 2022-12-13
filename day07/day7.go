package main

import (
	"bufio"
	"os"
	"strings"
)

type Dir struct {
	nome  string
	dim   int
	prec  *Dir
	figli map[string]*Dir
}

func main() {

}

// read a singol string for the input and updates the directories
func readFile() {
	main := Dir{}
	prec := Dir{}
	curr := main
	file, _ := os.Open("day7.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		if args[1] == "cd" {
			if args[2] == "/" {
				main = Dir{"/", 0, nil, make(map[string]*Dir)}
			}
			if args[2] != ".." {
				prec = curr
				curr = Dir{args[2], 0, *&curr.prec, make(map[string]*Dir)}

			}
		}

	}
}
