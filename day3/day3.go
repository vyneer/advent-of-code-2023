package day3

import (
	"aoc2023/day"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type solverFunc func([]string) (int, error)

type Day3 struct {
	name string
	path string
}

func New(path string) day.Solver {
	return &Day3{
		name: "Day 3",
		path: path,
	}
}

func (d *Day3) Solve() ([]string, error) {
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

func (d *Day3) GetDayString() string {
	return d.name
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

func newSchematic(lines []string, runeFunc func(rune) bool) (Schematic, error) {
	sch := Schematic{
		lines:    lines,
		runeFunc: runeFunc,
	}

	var err error
	sch.numberMap, sch.symbolMap, err = sch.getInfoMaps()
	if err != nil {
		return Schematic{}, err
	}

	return sch, nil
}

func part1(lines []string) (int, error) {
	sch, err := newSchematic(lines, func(r rune) bool {
		return !(r >= '0' && r <= '9') && r != '.'
	})
	if err != nil {
		return 0, err
	}

	return sch.calculateTask(calculateNumberSum), nil
}

func part2(lines []string) (int, error) {
	sch, err := newSchematic(lines, func(r rune) bool {
		return r == '*'
	})
	if err != nil {
		return 0, err
	}

	return sch.calculateTask(calculateGearRatio), nil
}

type Number struct {
	N     int
	Added bool
}

type Schematic struct {
	numberMap map[int]map[[2]int]*Number
	symbolMap map[int]map[int]bool
	lines     []string
	runeFunc  func(r rune) bool
}

func (sch *Schematic) calculateTask(taskFunc func(map[int]bool, []map[[2]int]*Number) int) int {
	buf := 0

	for lineIdx := range sch.lines {
		lineSymbols, ok := sch.symbolMap[lineIdx]
		if !ok {
			continue
		}

		adjRows := []map[[2]int]*Number{sch.numberMap[lineIdx-1], sch.numberMap[lineIdx], sch.numberMap[lineIdx+1]}

		buf += taskFunc(lineSymbols, adjRows)
	}

	return buf
}

func (sch *Schematic) getInfoMaps() (map[int]map[[2]int]*Number, map[int]map[int]bool, error) {
	numberMap := map[int]map[[2]int]*Number{}
	symbolMap := map[int]map[int]bool{}

	for i, s := range sch.lines {
		lineNumbers, err := sch.getLineNumberPosition(s)
		if err != nil {
			return nil, nil, err
		}
		if len(lineNumbers) > 0 {
			numberMap[i] = lineNumbers
		}

		lineSymbols := sch.getLineSymbolPosition(s)
		if len(lineSymbols) > 0 {
			symbolMap[i] = lineSymbols
		}
	}

	return numberMap, symbolMap, nil
}

func (sch *Schematic) getLineSymbolPosition(s string) map[int]bool {
	symbolLineMap := map[int]bool{}

	for stringIdx, r := range s {
		if sch.runeFunc(r) {
			symbolLineMap[stringIdx] = true
		}
	}

	return symbolLineMap
}

func (sch *Schematic) getLineNumberPosition(s string) (map[[2]int]*Number, error) {
	numberLineMap := map[[2]int]*Number{}

	buf := []rune{}
	for stringIdx, r := range s {
		if r > 47 && r < 58 {
			buf = append(buf, r)
			if stringIdx != len(s)-1 {
				continue
			}
		}
		if len(buf) > 0 {
			num, err := strconv.Atoi(string(buf))
			if err != nil {
				return nil, err
			}
			pos := [2]int{stringIdx - len(buf), stringIdx - 1}
			if stringIdx == len(s)-1 && (r >= '0' && r <= '9') {
				pos = [2]int{stringIdx - len(buf) + 1, stringIdx}
			}
			numberLineMap[pos] = &Number{
				N: num,
			}
			buf = []rune{}
		}
	}

	return numberLineMap, nil
}

func calculateNumberSum(lineSymbols map[int]bool, adjRows []map[[2]int]*Number) int {
	buf := 0

	for symbolIdx := range lineSymbols {
		leftBound := symbolIdx - 1
		rightBound := symbolIdx + 1

		for _, adjRow := range adjRows {
			for numIdx, v := range adjRow {
				if (numIdx[1] >= leftBound && numIdx[0] <= rightBound) && !v.Added {
					v.Added = true
					buf += v.N
				}
			}
		}
	}

	return buf
}

func calculateGearRatio(lineSymbols map[int]bool, adjRows []map[[2]int]*Number) int {
	buf := 0

	for symbolIdx := range lineSymbols {
		leftBound := symbolIdx - 1
		rightBound := symbolIdx + 1

		adj := []*Number{}

		for _, adjRow := range adjRows {
			for numIdx, v := range adjRow {
				if (numIdx[1] >= leftBound && numIdx[0] <= rightBound) && !v.Added {
					v.Added = true
					adj = append(adj, v)
				}
			}
		}

		if len(adj) == 2 {
			b := 1
			for _, v := range adj {
				b *= v.N
			}
			buf += b
		}
	}

	return buf
}
