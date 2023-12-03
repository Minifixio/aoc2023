package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(scanner *bufio.Scanner) int {
	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		r1 := 0
		r2 := 0
		for _, char := range line {
			if unicode.IsDigit(char) {
				if r1 == 0 {
					r1 = int(char - '0')
				}
				r2 = int(char - '0')
			}
		}
		res += (r1*10 + r2)
	}
	return res
}

func part2(scanner *bufio.Scanner) int {
	res := 0
	digits := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for scanner.Scan() {
		line := scanner.Text()
		r := 0
		i := 0
		r1 := 0
		r2 := 0
		for i < len(line) {
			if unicode.IsDigit(rune(line[i])) {
				r = int(line[i] - '0')
			} else {
				for ind, digit := range digits {
					j := i
					for j < len(line) && (j-i) < len(digit) && line[j] == digit[j-i] {
						j++
					}
					if j-i == len(digit) {
						r = ind + 1
						i = j - 2
						break
					}
				}
			}

			if r1 == 0 {
				r1 = r
			}
			r2 = r

			i++
		}
		res += (r1*10 + r2)
	}
	fmt.Println(res)
	return res
}

func main() {
	day := 1
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
