package main

var playerSteps = map[player][]*step{}

type player int32

const (
	X player = iota
	O
)

type step struct {
	CX int32
	CY int32
}

type playerStep struct {
	Player *player `json:"user"`
	Step   *step   `json:"step"`
}

func (p player) String() string {
	switch p {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return "-"
	}
}
