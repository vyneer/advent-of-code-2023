package main_test

import (
	"aoc2023/day"
	"aoc2023/day1"
	"aoc2023/day2"
	"aoc2023/day3"
	"aoc2023/day4"
	"slices"
	"testing"
)

func TestDay1(t *testing.T) {
	t.Run("Check Part 1", func(t *testing.T) {
		want := "142"

		d := day1.New(day.WithTest("day1_pt1.txt"))
		res, err := d.Solve()
		if err != nil {
			t.Fatal(err)
		}

		if res[0] != want {
			t.Fatalf("invalid answer, expected %s, got %s\n", want, res[0])
		}
	})

	t.Run("Check Part 2", func(t *testing.T) {
		want := "281"

		d := day1.New(day.WithTest("day1_pt2.txt"))
		res, err := d.Solve()
		if err != nil {
			t.Fatal(err)
		}

		if res[1] != want {
			t.Fatalf("invalid answer, expected %s, got %s\n", want, res[1])
		}
	})
}

func TestDay2(t *testing.T) {
	want := []string{"8", "2286"}

	d := day2.New(day.WithTest())
	res, err := d.Solve()
	if err != nil {
		t.Fatal(err)
	}

	if slices.Compare(want, res) != 0 {
		t.Fatalf("invalid answer, expected %+v, got %+v\n", want, res)
	}
}

func TestDay3(t *testing.T) {
	want := []string{"4361", "467835"}

	d := day3.New(day.WithTest())
	res, err := d.Solve()
	if err != nil {
		t.Fatal(err)
	}

	if slices.Compare(want, res) != 0 {
		t.Fatalf("invalid answer, expected %+v, got %+v\n", want, res)
	}
}

func TestDay4(t *testing.T) {
	want := []string{"13", "30"}

	d := day4.New(day.WithTest())
	res, err := d.Solve()
	if err != nil {
		t.Fatal(err)
	}

	if slices.Compare(want, res) != 0 {
		t.Fatalf("invalid answer, expected %+v, got %+v\n", want, res)
	}
}
