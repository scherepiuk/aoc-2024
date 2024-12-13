package day1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Solution struct{}

func (Solution) FirstPart() error {
	input, err := readInput()
	if err != nil {
		return fmt.Errorf("failed to get input: %w", err)
	}

	slices.Sort(input.left)
	slices.Sort(input.right)

	var distance int
	for i, left := range input.left {
		distance += absDistance(left, input.right[i])
	}

	fmt.Println(distance)
	return nil
}

func (Solution) SecondPart() error {
	input, err := readInput()
	if err != nil {
		return fmt.Errorf("failed to get input: %w", err)
	}

	rightCounts := make(map[int]int)
	for _, right := range input.right {
		rightCounts[right]++
	}

	var similarity int
	for _, left := range input.left {
		similarity += left * rightCounts[left]
	}

	fmt.Println(similarity)
	return nil
}

type input struct {
	right []int
	left  []int
}

func readInput() (*input, error) {
	file, err := os.Open("internal/day1/input")
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}

	input := input{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		left, right, err := parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse a line: %w", err)
		}

		input.left = append(input.left, left)
		input.right = append(input.right, right)
	}

	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("failed to read input file: %w", err)
	}

	return &input, nil
}

func parseLine(line string) (int, int, error) {
	parts := strings.Split(line, "   ")

	if len(parts) != 2 {
		return 0, 0, errors.New("invalid string")
	}

	left, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse left number: %w", err)
	}

	right, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse right number: %w", err)
	}

	return left, right, nil
}

func absDistance(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
