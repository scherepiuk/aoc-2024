package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/scherepiuk/aoc-2024/internal/mux"
)

func main() {
	err := internalRun()
	if err != nil {
		log.Fatal(err)
	}
}

func internalRun() error {
	if len(os.Args) != 2 {
		return errors.New("USAGE: aoc <day-number>")
	}

	dayNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return fmt.Errorf("failed to parse day number: %w", err)
	}

	solution, err := mux.GetSolution(dayNumber)
	if err != nil {
		return fmt.Errorf("failed to find solution: %w", err)
	}

	err = solution.FirstPart()
	if err != nil {
		return fmt.Errorf("solution (part 1) failed: %w", err)
	}

	err = solution.SecondPart()
	if err != nil {
		return fmt.Errorf("solution (part 2) failed: %w", err)
	}

	return nil
}
