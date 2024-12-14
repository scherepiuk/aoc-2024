package common

type Direction int

const (
	DirectionUnknown = iota
	DirectionAscending
	DirectionDescending
)

func AbsDistance(a, b int) (int, Direction) {
	if a < b {
		return b - a, DirectionAscending
	} else if a > b {
		return a - b, DirectionDescending
	}
	return 0, DirectionUnknown
}
