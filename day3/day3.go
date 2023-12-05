package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func part1(scanner *bufio.Scanner) int {
	res := 0
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		panic(err)
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		j := 0
		s := 0
		e := 0
		for j < len(line) {
			if unicode.IsDigit(rune(line[j])) {
				s = j
				for j < len(line) && unicode.IsDigit(rune(line[j])) {
					j++
				}
				e = j
				symbol := false

				for k := s - 1; k <= e; k++ {
					if i-1 > 0 {
						if k >= 0 && k < len(line) && lines[i-1][k] != '.' && !unicode.IsDigit(rune(lines[i-1][k])) {
							symbol = true
							break
						}
					}
					if i+1 < len(lines) {
						if k >= 0 && k < len(line) && lines[i+1][k] != '.' && !unicode.IsDigit(rune(lines[i+1][k])) {
							symbol = true
							break
						}
					}
				}
				if s-1 > 0 {
					if lines[i][s-1] != '.' && !unicode.IsDigit(rune(lines[i][s-1])) {
						symbol = true
					}
				}
				if e < len(line) {
					if rune(lines[i][e]) != '.' && !unicode.IsDigit(rune(lines[i][e])) {
						symbol = true
					}
				}

				if symbol {
					n, err := strconv.Atoi(line[s:e])
					fmt.Println(n)
					if err != nil {
						panic(err)
					}
					res += n
				}

			} else {
				j++
			}

		}
	}
	return res
}

func part2(scanner *bufio.Scanner) int {
	res := 0
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		panic(err)
	}

	m := make(map[string][]int)

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		j := 0
		s := 0
		e := 0
		for j < len(line) {
			if unicode.IsDigit(rune(line[j])) {
				s = j
				for j < len(line) && unicode.IsDigit(rune(line[j])) {
					j++
				}
				e = j

				n, err := strconv.Atoi(line[s:e])
				if err != nil {
					panic(err)
				}

				for k := s - 1; k <= e; k++ {
					if i-1 > 0 {
						if k >= 0 && k < len(line) && lines[i-1][k] == '*' && !unicode.IsDigit(rune(lines[i-1][k])) {
							key := fmt.Sprintf("%d,%d", i-1, k)
							m[key] = append(m[key], n)
							break
						}
					}
					if i+1 < len(lines) {
						if k >= 0 && k < len(line) && lines[i+1][k] == '*' && !unicode.IsDigit(rune(lines[i+1][k])) {
							key := fmt.Sprintf("%d,%d", i+1, k)
							m[key] = append(m[key], n)
							break
						}
					}
				}
				if s-1 > 0 {
					if lines[i][s-1] != '.' && !unicode.IsDigit(rune(lines[i][s-1])) {
						key := fmt.Sprintf("%d,%d", i, s-1)
						m[key] = append(m[key], n)
					}
				}
				if e < len(line) {
					if rune(lines[i][e]) != '.' && !unicode.IsDigit(rune(lines[i][e])) {
						key := fmt.Sprintf("%d,%d", i, e)
						m[key] = append(m[key], n)
					}
				}
			} else {
				j++
			}
		}
	}

	for _, v := range m {
		if len(v) == 2 {
			res += v[0] * v[1]
		}
	}

	return res
}

func main() {
	day := 3
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
		fmt.Println(part2(scanner))
	}
}
