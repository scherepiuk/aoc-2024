package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/scherepiuk/aoc-2024/internal/common"
)

type Solution struct{}

func (Solution) FirstPart() error {
	file, err := os.Open("internal/day2/input")
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}

	safe := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		levels, err := parseLine(line)
		if err != nil {
			return fmt.Errorf("failed to parse a line: %w", err)
		}

		if checkGradualProgression(levels) {
			safe += 1
		}
	}

	fmt.Println(safe)
	return nil
}

func (Solution) SecondPart() error {
	file, err := os.Open("internal/day2/input")
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}

	safe := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		levels, err := parseLine(line)
		if err != nil {
			return fmt.Errorf("failed to parse a line: %w", err)
		}

		for i := range len(levels) {
			filtered := copyWithExclusion(levels, i)
			if checkGradualProgression(filtered) {
				safe += 1
				break
			}
		}
	}

	fmt.Println(safe)
	return nil
}

func parseLine(line string) ([]int, error) {
	parts := strings.Split(line, " ")

	levels := make([]int, 0, len(parts))
	for _, part := range parts {
		level, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid level: %w", err)
		}

		levels = append(levels, level)
	}

	return levels, nil
}

func checkGradualProgression(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	var prevDirection common.Direction = common.DirectionUnknown

	for i := 0; i < len(levels)-1; i++ {
		var (
			prev = levels[i]
			curr = levels[i+1]
		)

		distance, direction := common.AbsDistance(prev, curr)

		if prevDirection == common.DirectionUnknown {
			prevDirection = direction
		}

		if direction == common.DirectionUnknown || direction != prevDirection {
			return false
		}

		gradualProgression := distance >= 1 && distance <= 3
		if !gradualProgression {
			return false
		}
	}

	return true
}

func copyWithExclusion(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}

	result := make([]int, 0, len(slice)-1)
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)
	return result
}
