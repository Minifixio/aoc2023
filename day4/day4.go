package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func part1(scanner *bufio.Scanner) int {
	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		separatorFunc := func(c rune) bool {
			return c == ':' || c == '|'
		}
		parts := strings.FieldsFunc(line, separatorFunc)
		split := [][]string{strings.Fields(strings.Trim(parts[1], " ")), strings.Fields(strings.Trim(parts[2], " "))}
		c := 0
		for _, a := range split[0] {
			for _, b := range split[1] {
				if a == b {
					if c == 0 {
						c = 1
					} else {
						c = c * 2
					}
				}
			}
		}
		res += c
	}
	return res
}

func part2(scanner *bufio.Scanner) int {
	res := 0
	m := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		separatorFunc := func(c rune) bool {
			return c == ':' || c == '|'
		}
		parts := strings.FieldsFunc(line, separatorFunc)
		split := [][]string{strings.Fields(strings.Trim(parts[1], " ")), strings.Fields(strings.Trim(parts[2], " "))}
		c := 0
		for _, a := range split[0] {
			for _, b := range split[1] {
				if a == b {
					c += 1
				}
			}
		}
		m = append(m, []int{c, 1})
	}

	for ind, v := range m {
		c := v[0]
		t := v[1]
		for i := 1; i <= c; i++ {
			if ind+i < len(m) {
				m[ind+i][1] += t
			}
		}
		res += t
	}
	return res
}

func main() {
	day := 4
	var part int
	var test bool
	flag.IntVar(&part, "part", 1, "Part 1 or 2")
	flag.BoolVar(&test, "test", false, "Test or not")
	flag.Parse()
	var f *os.File
	var err error
	if test {
		f, err = os.Open(fmt.Sprintf("data/day%d_test.txt", day))
	} else {
		f, err = os.Open(fmt.Sprintf("data/day%d.txt", day))
	}
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	if part == 1 {
		s := part1(scanner)
		fmt.Println(s)
	} else {
		s := part2(scanner)
		fmt.Println(s)
	}
}
