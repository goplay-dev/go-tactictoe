package tictactoe

type Dimension struct {
	Current int32 `json:"current"`
	Min     int32 `json:"min"`
	Max     int32 `json:"max"`
}

var playerSteps = map[player][]*Step{}

type player int32

const (
	X player = iota
	O
)

type Step struct {
	CX int32
	CY int32
}

type playerStep struct {
	Player *player `json:"user"`
	Step   *Step   `json:"step"`
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

type ValidateStepReq struct {
	Step1 *Step
	Step2 *Step
}

type GameConfig struct {
	Dimension      *Dimension
	AvailableSteps []*Step
	WinSteps       [][]*Step
}
