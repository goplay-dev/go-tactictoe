package tictactoe

import (
	"context"
	"fmt"
)

type InitGameConfig interface {
	ValidatePlayerStep(ctx context.Context, pStep *playerStep) bool
	RemoveSelectedStep(ctx context.Context, step *Step)
	SaveStep(ctx context.Context, pStep *playerStep, playerSteps map[player][]*Step)
	GetActualPos(ctx context.Context) [][]string
	ValidateStep(ctx context.Context, req *ValidateStepReq) bool

	CheckingIsStepObtained(ctx context.Context, pStep *playerStep, playerSteps map[player][]*Step) bool
	CheckWinStep(ctx context.Context, pStep *playerStep) bool
}

type gameConfig struct {
	*GameConfig
}

func InitGame(ctx context.Context, config *GameConfig) (InitGameConfig, error) {
	if config.Dimension.Current < config.Dimension.Min ||
		config.Dimension.Current > config.Dimension.Max {

		return nil, fmt.Errorf(fmt.Sprintf("dimension must between %d - %d",
			config.Dimension.Min, config.Dimension.Max))
	}

	config.AvailableSteps = SetupAvailableSteps(ctx, config.Dimension.Current)
	config.WinSteps = SetupWinSteps(ctx, config)

	return &gameConfig{config}, nil
}

func SetupAvailableSteps(ctx context.Context, currDimension int32) [][]*Step {
	var aSteps [][]*Step

	for cy := int32(0); cy < currDimension; cy++ {
		var aStep []*Step
		for cx := int32(0); cx < currDimension; cx++ {
			aStep = append(aStep, &Step{CX: cx, CY: cy})
		}

		aSteps = append(aSteps, aStep)
	}

	return aSteps
}

func SetupWinSteps(ctx context.Context, config *GameConfig) [][]*Step {
	var wSteps [][]*Step

	wSteps = append(wSteps, SetupHorWinSteps(ctx, config)...)
	wSteps = append(wSteps, SetupVerWinSteps(ctx, config)...)
	wSteps = append(wSteps, SetupLDiagWinSteps(ctx, config)...)
	wSteps = append(wSteps, SetupRDiagWinSteps(ctx, config)...)

	return wSteps
}

func SetupHorWinSteps(ctx context.Context, config *GameConfig) [][]*Step {
	var hWinSteps [][]*Step
	var currDimension = config.Dimension.Current
	var availableSteps = config.AvailableSteps

	for cy := int32(0); cy < currDimension; cy++ {
		var hWinStep []*Step
		for cx := int32(0); cx < currDimension; cx++ {
			hWinStep = append(hWinStep, availableSteps[cy][cx])
		}
		hWinSteps = append(hWinSteps, hWinStep)
	}

	return hWinSteps
}

func SetupVerWinSteps(ctx context.Context, config *GameConfig) [][]*Step {
	var vWinSteps [][]*Step
	var currDimension = config.Dimension.Current
	var availableSteps = config.AvailableSteps

	for cx := int32(0); cx < currDimension; cx++ {
		var vSteps []*Step
		for cy := int32(0); cy < currDimension; cy++ {
			vSteps = append(vSteps, availableSteps[cy][cx])
		}

		vWinSteps = append(vWinSteps, vSteps)
	}

	return vWinSteps
}

func SetupLDiagWinSteps(ctx context.Context, config *GameConfig) [][]*Step {
	var dWinSteps [][]*Step
	var ldWinSteps []*Step

	var currDimension = config.Dimension.Current
	var availableSteps = config.AvailableSteps

	for ld := int32(0); ld < currDimension; ld++ {
		ldWinSteps = append(ldWinSteps, availableSteps[ld][ld])
	}

	return append(dWinSteps, ldWinSteps)
}

func SetupRDiagWinSteps(ctx context.Context, config *GameConfig) [][]*Step {
	var dWinSteps [][]*Step
	var ldWinSteps []*Step
	var currDimension = config.Dimension.Current
	var availableSteps = config.AvailableSteps

	for ld := currDimension - 1; ld >= 0; ld-- {
		ldWinSteps = append(ldWinSteps, availableSteps[(currDimension-1)-ld][ld])
	}

	return append(dWinSteps, ldWinSteps)
}

func (g *gameConfig) ValidatePlayerStep(ctx context.Context, pStep *playerStep) bool {
	if g.ValidateStep(ctx, &ValidateStepReq{
		Step1: pStep.Step,
		Step2: g.AvailableSteps[pStep.Step.CY][pStep.Step.CX],
	}) {
		g.RemoveSelectedStep(ctx, pStep.Step)
		return true
	}

	return false
}

func (g *gameConfig) RemoveSelectedStep(ctx context.Context, step *Step) {
	g.AvailableSteps[step.CY] = append(g.AvailableSteps[step.CY][:step.CX], g.AvailableSteps[step.CY][step.CX+1:]...)
}

func (g *gameConfig) SaveStep(ctx context.Context, pStep *playerStep, playerSteps map[player][]*Step) {
	playerSteps[*pStep.Player] = append(playerSteps[*pStep.Player], pStep.Step)
}

func (g *gameConfig) GetActualPos(ctx context.Context) [][]string {
	var positions [][]string

	for cy := 0; cy < int(g.Dimension.Current); cy++ {
		var position []string

		for cx := 0; cx < int(g.Dimension.Current); cx++ {
			pos := "-"

			stepPoint := &Step{
				CX: int32(cx),
				CY: int32(cy),
			}

			playerX := X
			playerO := O

			isX := g.CheckingIsStepObtained(ctx, &playerStep{
				Player: &playerX,
				Step:   stepPoint,
			}, playerSteps)
			if isX {
				pos = X.String()
			} else {
				isO := g.CheckingIsStepObtained(ctx, &playerStep{
					Player: &playerO,
					Step:   stepPoint,
				}, playerSteps)
				if isO {
					pos = O.String()
				}
			}

			position = append(position, pos)
		}

		positions = append(positions, position)
	}

	return positions
}

func (g *gameConfig) ValidateStep(ctx context.Context, req *ValidateStepReq) bool {
	if req.Step1.CX == req.Step2.CX && req.Step1.CY == req.Step2.CY {
		return true
	}

	return false
}

func (g *gameConfig) CheckingIsStepObtained(
	ctx context.Context, pStep *playerStep, playerSteps map[player][]*Step) bool {
	for _, pSteps := range playerSteps[*pStep.Player] {
		if g.ValidateStep(ctx, &ValidateStepReq{
			Step1: pSteps,
			Step2: pStep.Step,
		}) {
			return true
		}
	}

	return false
}

func (g *gameConfig) CheckWinStep(ctx context.Context, pStep *playerStep) bool {
	var knownWinSteps []*Step

	for _, winStep := range g.WinSteps {
		knownWinSteps = []*Step{}
		for _, ws := range winStep {
			var knownPlayerStep *Step

			for _, ps := range playerSteps[*pStep.Player] {
				if g.ValidateStep(ctx, &ValidateStepReq{
					Step1: ws,
					Step2: ps,
				}) {
					knownPlayerStep = ps
					knownWinSteps = append(knownWinSteps, knownPlayerStep)
					break
				}
			}

			if knownPlayerStep == nil {
				break
			}
		}

		if len(knownWinSteps) == int(g.Dimension.Current) {
			return true
		}
	}

	return false
}
