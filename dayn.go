package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func part1(scanner *bufio.Scanner) int {
	res := 0
	return res
}

func part2(scanner *bufio.Scanner) int {
	res := 0
	return res
}

func main() {
	day := -1
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
