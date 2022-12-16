package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(monkeys(readFile()))
}

type monkey struct {
	si                   []int
	op                   []string
	test, t, f, ispected int
}

func readFile() []monkey {
	s := []monkey{}
	file, _ := os.Open("day11.txt")
	scanner := bufio.NewScanner(file)
	var c, n int
	var m monkey
	for scanner.Scan() {
		args := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		c++
		if len(args) > 1 {
			if args[1] == "true:" {
				m.t, _ = strconv.Atoi(args[5])
			}
			if args[1] == "false:" {
				m.f, _ = strconv.Atoi(args[5])
			}
			if args[0] == "Test:" {
				m.test, _ = strconv.Atoi(args[3])
			}
			if args[0] == "Operation:" {
				m.op = append(m.op, args[4], args[5])
			}
			if args[0] == "Starting" {
				for i := 2; i < len(args); i++ {
					if args[i][len(args[i])-1] == ',' {
						args[i] = args[i][:len(args[i])-1]
					}
					num, _ := strconv.Atoi(args[i])
					m.si = append(m.si, num)
				}
			}
		}
		if c == 7 {
			n++
			c = 0
			s = append(s, m)
			m = monkey{}
		}
	}
	s = append(s, m)
	return s
}

func operation(m monkey) monkey {
	num, _ := strconv.Atoi(m.op[1])
	for i := 0; i < len(m.si); i++ {
		if m.op[0] == "*" && m.op[1] != "old" {
			m.si[i] = m.si[i] * num
		} else if m.op[0] == "+" {
			m.si[i] = m.si[i] + num
		} else if m.op[0] == "*" && m.op[1] == "old" {
			m.si[i] = (m.si[i] * m.si[i])
		}
		m.si[i] = m.si[i] % (11 * 17 * 5 * 13 * 19 * 2 * 3 * 7)
	}

	return m
}

func test(s []monkey, idx int) []monkey {
	var x int = len(s[idx].si)
	s[idx].ispected += x
	for i := 0; i < x; i++ {
		if s[idx].si[0]%s[idx].test == 0 {
			s[s[idx].t].si = append(s[s[idx].t].si, s[idx].si[0])
		} else {
			s[s[idx].f].si = append(s[s[idx].f].si, s[idx].si[0])
		}
		sl := []int{}
		sl = append(sl, s[idx].si[1:]...)
		s[idx].si = sl
	}
	return s
}

func monkeys(s []monkey) int {
	var max1, max2 int
	for i := 0; i < 10000; i++ {
		for v, m := range s {
			m = operation(m)
			s = test(s, v)
		}
	}
	for _, m := range s {
		fmt.Println(m.ispected)
		if m.ispected > max2 {
			max1 = max2
			max2 = m.ispected
		}
	}
	return max1 * max2
}
