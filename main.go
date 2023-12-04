package main

import (
	"aoc2023/day"
	"aoc2023/day1"
	"aoc2023/day2"
	"aoc2023/day3"
	"aoc2023/day4"
	"fmt"
	"log"
)

func main() {
	days := []day.Solver{
		day1.New(),
		day2.New(),
		day3.New(),
		day4.New(),
	}

	for _, v := range days {
		a, err := v.Solve()
		if err != nil {
			log.Fatalln(err)
		}

		for i, res := range a {
			fmt.Printf("%s (Part %d): %s\n", v.GetDayString(), i+1, res)
		}
	}
}
