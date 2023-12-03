package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) int {
	res := 0
	l := 1
	for scanner.Scan() {
		line := scanner.Text()
		separatorFunc := func(c rune) bool {
			return c == ';' || c == ',' || c == ':'
		}
		parts := strings.FieldsFunc(line, separatorFunc)[1:]
		valid := true

		for _, part := range parts {
			part = strings.Trim(part, " ")
			n, err := strconv.Atoi(strings.Split(part, " ")[0])
			if err != nil {
				panic(err)
			}
			if strings.Contains(part, "red") {
				if n > 12 {
					valid = false
					break
				}
			}
			if strings.Contains(part, "green") {
				if n > 13 {
					valid = false
					break
				}
			}
			if strings.Contains(part, "blue") {
				if n > 14 {
					valid = false
					break
				}
			}
		}

		if valid {
			res += l
		}

		l++
	}
	return res
}

func part2(scanner *bufio.Scanner) int {
	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		separatorFunc := func(c rune) bool {
			return c == ';' || c == ',' || c == ':'
		}
		parts := strings.FieldsFunc(line, separatorFunc)[1:]
		minRed := 0
		minGreen := 0
		minBlue := 0

		for _, part := range parts {
			part = strings.Trim(part, " ")
			n, err := strconv.Atoi(strings.Split(part, " ")[0])
			if err != nil {
				panic(err)
			}
			if strings.Contains(part, "red") {
				if n > minRed {
					minRed = n
				}
			}
			if strings.Contains(part, "green") {
				if n > minGreen {
					minGreen = n
				}
			}
			if strings.Contains(part, "blue") {
				if n > minBlue {
					minBlue = n
				}
			}
		}

		res += minRed * minGreen * minBlue
	}

	return res
}

func main() {
	day := 2
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
		fmt.Println(part1(scanner))
	} else {
		fmt.Println(part2(scanner))
	}
}
