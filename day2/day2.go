package day2

import (
	"aoc2023/day"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	redMax   int = 12
	greenMax int = 13
	blueMax  int = 14
)

type solverFunc func(string) (int, error)

type Day2 struct {
	*day.Opts
}

func New(opts ...day.OptFunc) day.Solver {
	d := Day2{}
	d.Opts = day.DefaultOpts(&d)

	for _, o := range opts {
		o(d.Opts)
	}

	return &d
}

func (d *Day2) Solve() ([]string, error) {
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

func (d *Day2) GetDayString() string {
	return d.Name
}

type Pull struct {
	Red   int
	Green int
	Blue  int
}

func (game *Pull) getCubePower() int {
	return game.Red * game.Green * game.Blue
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
		num, err := solverFunc(scanner.Text())
		if err != nil {
			return 0, nil
		}

		sum += num
	}

	return sum, nil
}

func part1(s string) (int, error) {
	splitGame := strings.SplitN(s, ":", 2)

	id, err := parseGameID(splitGame[0])
	if err != nil {
		return 0, err
	}

	results, err := parseGameResults(splitGame[1])
	if err != nil {
		return 0, err
	}

	for _, pull := range results {
		if pull.Red > redMax || pull.Green > greenMax || pull.Blue > blueMax {
			return 0, nil
		}
	}

	return id, nil
}

func part2(s string) (int, error) {
	var game Pull

	splitGame := strings.SplitN(s, ":", 2)

	results, err := parseGameResults(splitGame[1])
	if err != nil {
		return 0, err
	}

	for _, pull := range results {
		if pull.Red > game.Red {
			game.Red = pull.Red
		}
		if pull.Green > game.Green {
			game.Green = pull.Green
		}
		if pull.Blue > game.Blue {
			game.Blue = pull.Blue
		}
	}

	power := game.getCubePower()

	return power, nil
}

func parseGameID(s string) (int, error) {
	gameIDString := strings.Fields(s)[1]
	gameIDString = strings.TrimSpace(gameIDString)
	gameID, err := strconv.Atoi(gameIDString)
	if err != nil {
		return 0, err
	}

	return gameID, nil
}

func parseGameResults(s string) ([]Pull, error) {
	pullSlice := []Pull{}

	pulls := strings.Split(s, ";")
	for _, pull := range pulls {
		p, err := parsePull(pull)
		if err != nil {
			return nil, err
		}

		pullSlice = append(pullSlice, p)
	}

	return pullSlice, nil
}

func parsePull(s string) (Pull, error) {
	var pull Pull

	colors := strings.Split(s, ",")
	for _, color := range colors {
		name, count, err := parseColor(color)
		if err != nil {
			return pull, err
		}

		switch name {
		case "red":
			pull.Red += count
		case "green":
			pull.Green += count
		case "blue":
			pull.Blue += count
		}
	}

	return pull, nil
}

func parseColor(s string) (string, int, error) {
	color := strings.TrimSpace(s)
	res := strings.Fields(color)

	num, err := strconv.Atoi(res[0])
	if err != nil {
		return "", 0, err
	}

	return res[1], num, nil
}
