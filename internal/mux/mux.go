package mux

import (
	"errors"

	"github.com/scherepiuk/aoc-2024/internal/day1"
	"github.com/scherepiuk/aoc-2024/internal/day2"
)

type Solution interface {
	FirstPart() error
	SecondPart() error
}

var solutions = []Solution{
	new(day1.Solution),
	new(day2.Solution),
}

var ErrInvalidDayNumber = errors.New("invalid day number")

func GetSolution(dayNumber int) (Solution, error) {
	index := dayNumber - 1

	if index < 0 || index > len(solutions) {
		var zero Solution
		return zero, ErrInvalidDayNumber
	}

	return solutions[index], nil
}
