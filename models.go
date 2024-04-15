package go_tictactoe

type Dimension struct {
	Current int32 `json:"current"`
	Min     int32 `json:"min"`
	Max     int32 `json:"max"`
}

type Step struct {
	CX int32
	CY int32
}

type Player int32
type PlayerMark string
type Players []PlayerMark
type PlayersList []Players
type ActualPositions PlayersList

type Steps []*Step
type StepsList []Steps

type WinPos string
type WinSteps map[WinPos]StepsList

const (
	Hor   = WinPos("Hor")
	Ver   = WinPos("Ver")
	LDiag = WinPos("LDiag")
	RDiag = WinPos("RDiag")
)

const (
	X Player = iota
	O
	E
)

type PlayerStepReq struct {
	Player *Player `json:"player"`
	Step   *Step   `json:"step"`
}

type ValidateStepReq struct {
	Step1 *Step
	Step2 *Step
}

type GameConfig struct {
	*Dimension
	WinSteps
	ActualPositions
}

func (p Player) String() PlayerMark {
	switch p {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return ""
	}
}
