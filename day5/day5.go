package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	mappa := map[string][]string{"1": {"B", "W", "N"}, "2": {"L", "Z", "S", "P", "T", "D", "M", "B"}, "3": {"Q", "H", "Z", "W", "R"}, "4": {"W", "D", "V", "J", "Z", "R"}, "5": {"S", "H", "M", "B"}, "6": {"L", "G", "N", "J", "H", "V", "P", "V"}, "7": {"J", "Q", "Z", "F", "H", "D", "L", "S"}, "8": {"W", "S", "F", "J", "G", "Q", "B"}, "9": {"z", "W", "M", "S", "C", "D", "J"}}

	fmt.Println(readFile(mappa))

}

func readFile(mappa map[string][]string) string {
	file, _ := os.Open("day5instructions.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		move(mappa, scanner.Text())
	}
	return finalWord(mappa)
}

func move(mappa map[string][]string, s string) {
	args := strings.Split(s, " ")
	num, _ := strconv.Atoi(args[1])
	//mappa[args[5]] = append(mappa[args[5]], reverse(mappa[args[3]][len(mappa[args[3]])-num:])...)
	mappa[args[5]] = append(mappa[args[5]], mappa[args[3]][len(mappa[args[3]])-num:]...)
	mappa[args[3]] = mappa[args[3]][:len(mappa[args[3]])-num]
}
func finalWord(mappa map[string][]string) string {
	var s string
	var count int = 1
	for count <= len(mappa) {
		idx := strconv.Itoa(count)
		s += mappa[idx][len(mappa[idx])-1]
		count++
	}
	return s
}

func reverse(s []string) []string {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return s
}
