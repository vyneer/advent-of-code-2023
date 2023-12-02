package day1

import (
	"aoc2023/day"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

type solverFunc func(string) []string

type Day1 struct {
	name string
	path string
}

func New(path string) day.Solver {
	return &Day1{
		name: "Day 1",
		path: path,
	}
}

func (d *Day1) Solve() ([]string, error) {
	answers := []string{}
	solvers := []solverFunc{
		part1,
		part2,
	}

	for _, v := range solvers {
		num, err := solve(d.path, v)
		if err != nil {
			return nil, err
		}

		answers = append(answers, fmt.Sprintf("%d", num))
	}

	return answers, nil
}

func (d *Day1) GetDayString() string {
	return d.name
}

func solve(path string, solverFunc solverFunc) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		numSlice := solverFunc(scanner.Text())

		num, err := strconv.Atoi(fmt.Sprintf("%s%s", numSlice[0], numSlice[len(numSlice)-1]))
		if err != nil {
			return 0, err
		}

		sum += num
	}

	return sum, nil
}

func part1(s string) []string {
	buf := []string{}

	for _, r := range s {
		if n, ok := numberRuneToString(r); ok {
			buf = append(buf, n)
		}
	}

	return buf
}

func part2(s string) []string {
	buf := []string{}

	idxMap := getNumberWordIndices(s)

	for i, r := range s {
		if n, ok := numberWordToString(i, s, idxMap); ok {
			buf = append(buf, n)
		}
		if n, ok := numberRuneToString(r); ok {
			buf = append(buf, n)
		}
	}

	return buf
}

func numberRuneToString(r rune) (string, bool) {
	if r > 47 && r < 58 {
		return string(r), true
	}

	return "", false
}

func numberWordToString(i int, s string, idxMap map[int]string) (string, bool) {
	if v, ok := idxMap[i]; ok {
		idx := strings.Index(s, v)
		part := s[idx : idx+len(v)]

		return numbers[part], true
	}

	return "", false
}

func getNumberWordIndices(s string) map[int]string {
	idxMap := map[int]string{}

	for k := range numbers {
		idx := 0
		working := true
		for working {
			nIdx := strings.Index(s[idx:], k)
			if nIdx < 0 {
				working = false
				continue
			}
			idxMap[idx+nIdx] = k
			idx = idx + nIdx + 1
		}
	}

	return idxMap
}
