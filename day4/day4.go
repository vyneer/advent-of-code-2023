package day4

import (
	"aoc2023/day"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type solverFunc func([]string) (int, error)

type Day4 struct {
	*day.Opts
}

func New(opts ...day.OptFunc) day.Solver {
	d := Day4{}
	d.Opts = day.DefaultOpts(&d)

	for _, o := range opts {
		o(d.Opts)
	}

	return &d
}

func (d *Day4) Solve() ([]string, error) {
	answers := []string{}
	solvers := []solverFunc{
		part1,
		part2,
	}

	for _, v := range solvers {
		num, err := solve(d.Path, v)
		if err != nil {
			return nil, err
		}

		answers = append(answers, fmt.Sprintf("%d", num))
	}

	return answers, nil
}

func (d *Day4) GetDayString() string {
	return d.Name
}

func solve(path string, solverFunc solverFunc) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	num, err := solverFunc(lines)
	if err != nil {
		return 0, nil
	}

	return num, nil
}

type card struct {
	id       int
	winning  []int
	have     []int
	matching []int
}

func part1(lines []string) (int, error) {
	sum := 0

	for _, s := range lines {
		c, err := newCard(s)
		if err != nil {
			return 0, err
		}

		sum += c.calculatePoints()
	}

	return sum, nil
}

func part2(lines []string) (int, error) {
	cardMap := map[int]int{}

	for _, s := range lines {
		c, err := newCard(s)
		if err != nil {
			return 0, err
		}

		cardMap = modifyCardMap(cardMap, c)
	}

	return sumCards(cardMap), nil
}

func newCard(s string) (card, error) {
	var err error
	c := card{}

	splitCard := strings.SplitN(s, ":", 2)

	c.id, err = parseCardID(splitCard[0])
	if err != nil {
		return card{}, err
	}

	splitNumbers := strings.SplitN(splitCard[1], "|", 2)

	c.winning, err = parseNumbers(splitNumbers[0])
	if err != nil {
		return card{}, err
	}

	c.have, err = parseNumbers(splitNumbers[1])
	if err != nil {
		return card{}, err
	}

	for _, num := range c.have {
		if slices.Contains(c.winning, num) {
			c.matching = append(c.matching, num)
		}
	}

	return c, nil
}

func (c *card) calculatePoints() int {
	return int(math.Pow(2, float64(len(c.matching))-1))
}

func modifyCardMap(cardMap map[int]int, c card) map[int]int {
	cardMap[c.id]++

	if len(c.matching) > 0 {
		for i := 0; i < cardMap[c.id]; i++ {
			for i := 1; i < len(c.matching)+1; i++ {
				cardMap[c.id+i]++
			}
		}
	}

	return cardMap
}

func sumCards(cardMap map[int]int) int {
	buf := 0

	for _, v := range cardMap {
		buf += v
	}

	return buf
}

func parseCardID(s string) (int, error) {
	cardIDString := strings.Fields(s)[1]
	cardIDString = strings.TrimSpace(cardIDString)
	cardID, err := strconv.Atoi(cardIDString)
	if err != nil {
		return 0, err
	}

	return cardID, nil
}

func parseNumbers(s string) ([]int, error) {
	numInts := []int{}
	numStrings := strings.Fields(s)

	for _, v := range numStrings {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		numInts = append(numInts, i)
	}

	return numInts, nil
}
