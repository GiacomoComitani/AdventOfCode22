package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type arg struct {
	you string
	me  string
}

func main() {
	s := leggiFile()
	fmt.Println(score(s))

}

func leggiFile() []arg {
	slice := []arg{}
	file, _ := os.Open("day2.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		slice = append(slice, arg{args[0], args[1]})
	}
	return slice
}

func score(slice []arg) int {
	var ris int
	for _, v := range slice {
		ris += result(v.you, v.me)
	}
	return ris
}

// first part
/*
func result(y, m string) int {
	if y == "A" {
		switch m {
		case "X":
			return 3 + 1
		case "Y":
			return 6 + 2
		case "Z":
			return 0 + 3
		}
	} else if y == "B" {
		switch m {
		case "X":
			return 0 + 1
		case "Y":
			return 3 + 2
		case "Z":
			return 6 + 3
		}
	} else if y == "C" {
		switch m {
		case "X":
			return 6 + 1
		case "Y":
			return 0 + 2
		case "Z":
			return 3 + 3
		}
	}
	return 0
}*/

//second part
func result(y, m string) int {
	if y == "A" {
		switch m {
		case "X":
			return 0 + 3
		case "Y":
			return 3 + 1
		case "Z":
			return 6 + 2
		}
	} else if y == "B" {
		switch m {
		case "X":
			return 0 + 1
		case "Y":
			return 3 + 2
		case "Z":
			return 6 + 3
		}
	} else if y == "C" {
		switch m {
		case "X":
			return 0 + 2
		case "Y":
			return 3 + 3
		case "Z":
			return 6 + 1
		}
	}
	return 0
}
