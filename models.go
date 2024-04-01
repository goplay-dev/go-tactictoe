package tictactoe

type Dimension struct {
	Current int32 `json:"current"`
	Min     int32 `json:"min"`
	Max     int32 `json:"max"`
}

type ValidatePlayerStepReq struct {
	PlayerStep     *PlayerStep
	AvailableSteps AvailableSteps
}

type RemoveSelectedStepReq struct {
	Step           *Step
	AvailableSteps AvailableSteps
}

type PlayerStepReq struct {
	PlayerStep  *PlayerStep
	PlayerSteps PlayerSteps
}

type GetActualPosReq struct {
	Dimension   *Dimension
	PlayerSteps PlayerSteps
}

type CheckWinStepReq struct {
	PlayerStep  *PlayerStep
	WinSteps    WinSteps
	PlayerSteps PlayerSteps
	Dimension   *Dimension
}

type PlayerSteps map[Player][]*Step
type AvailableSteps [][]*Step
type WinSteps [][]*Step
type Player int32
type ActualPositions [][]string

const (
	X Player = iota
	O
)

type Step struct {
	CX int32
	CY int32
}

type PlayerStep struct {
	Player *Player `json:"user"`
	Step   *Step   `json:"step"`
}

func (p Player) String() string {
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
	AvailableSteps AvailableSteps
	WinSteps       WinSteps
}
