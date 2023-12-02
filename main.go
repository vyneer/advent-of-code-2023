package main

import (
	"aoc2023/day"
	"aoc2023/day1"
	"fmt"
	"log"
)

const inputFolder string = "./inputs"

func main() {
	days := []day.Solver{
		day1.New(fmt.Sprintf("%s/%s", inputFolder, "day1.txt")),
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
