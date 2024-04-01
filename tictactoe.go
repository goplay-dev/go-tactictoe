package tictactoe

import (
	"context"
)

func SetupAvailableSteps(ctx context.Context, currDimension int32) AvailableSteps {
	var aSteps AvailableSteps

	for cy := int32(0); cy < currDimension; cy++ {
		var aStep []*Step
		for cx := int32(0); cx < currDimension; cx++ {
			aStep = append(aStep, &Step{CX: cx, CY: cy})
		}

		aSteps = append(aSteps, aStep)
	}

	return aSteps
}

func SetupWinSteps(ctx context.Context, config *GameConfig) WinSteps {
	var wSteps WinSteps

	wSteps = append(wSteps, SetupHorWinSteps(ctx, config)...)
	wSteps = append(wSteps, SetupVerWinSteps(ctx, config)...)
	wSteps = append(wSteps, SetupLDiagWinSteps(ctx, config)...)
	wSteps = append(wSteps, SetupRDiagWinSteps(ctx, config)...)

	return wSteps
}

func SetupHorWinSteps(ctx context.Context, config *GameConfig) WinSteps {
	var hWinSteps WinSteps
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

func SetupVerWinSteps(ctx context.Context, config *GameConfig) WinSteps {
	var vWinSteps WinSteps
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

func SetupLDiagWinSteps(ctx context.Context, config *GameConfig) WinSteps {
	var dWinSteps WinSteps
	var ldWinSteps []*Step

	var currDimension = config.Dimension.Current
	var availableSteps = config.AvailableSteps

	for ld := int32(0); ld < currDimension; ld++ {
		ldWinSteps = append(ldWinSteps, availableSteps[ld][ld])
	}

	return append(dWinSteps, ldWinSteps)
}

func SetupRDiagWinSteps(ctx context.Context, config *GameConfig) WinSteps {
	var dWinSteps WinSteps
	var ldWinSteps []*Step
	var currDimension = config.Dimension.Current
	var availableSteps = config.AvailableSteps

	for ld := currDimension - 1; ld >= 0; ld-- {
		ldWinSteps = append(ldWinSteps, availableSteps[(currDimension-1)-ld][ld])
	}

	return append(dWinSteps, ldWinSteps)
}

func RemoveSelectedStep(ctx context.Context, req *RemoveSelectedStepReq) {
	availableSteps := req.AvailableSteps
	step := req.Step

	availableSteps[step.CY] = append(availableSteps[step.CY][:step.CX], availableSteps[step.CY][step.CX+1:]...)
}

func ValidatePlayerStep(ctx context.Context, req *ValidatePlayerStepReq) bool {
	pStep := req.PlayerStep
	availableSteps := req.AvailableSteps

	if ValidateStep(ctx, &ValidateStepReq{
		Step1: pStep.Step,
		Step2: availableSteps[pStep.Step.CY][pStep.Step.CX],
	}) {
		RemoveSelectedStep(ctx, &RemoveSelectedStepReq{
			Step:           req.PlayerStep.Step,
			AvailableSteps: req.AvailableSteps,
		})

		return true
	}

	return false
}

func ValidateStep(ctx context.Context, req *ValidateStepReq) bool {
	if req.Step1.CX == req.Step2.CX && req.Step1.CY == req.Step2.CY {
		return true
	}

	return false
}

func SavePlayerStep(ctx context.Context, req *PlayerStepReq) PlayerSteps {
	playerSteps := req.PlayerSteps
	pStep := req.PlayerStep

	playerSteps[*pStep.Player] = append(playerSteps[*pStep.Player], pStep.Step)
	return playerSteps
}

func CheckingIsStepObtained(ctx context.Context, req *PlayerStepReq) bool {
	playerSteps := req.PlayerSteps
	pStep := req.PlayerStep

	for _, pSteps := range playerSteps[*pStep.Player] {
		if ValidateStep(ctx, &ValidateStepReq{
			Step1: pSteps,
			Step2: pStep.Step,
		}) {
			return true
		}
	}

	return false
}

func GetActualPos(ctx context.Context, req *GetActualPosReq) ActualPositions {
	var positions ActualPositions

	dimension := req.Dimension
	playerSteps := req.PlayerSteps

	for cy := 0; cy < int(dimension.Current); cy++ {
		var position []string

		for cx := 0; cx < int(dimension.Current); cx++ {
			pos := "-"

			stepPoint := &Step{
				CX: int32(cx),
				CY: int32(cy),
			}

			playerX := X
			playerO := O

			playerStepReq := &PlayerStepReq{
				PlayerStep: &PlayerStep{
					Player: &playerX,
					Step:   stepPoint,
				},
				PlayerSteps: playerSteps,
			}

			isX := CheckingIsStepObtained(ctx, playerStepReq)
			if isX {
				pos = X.String()
			} else {
				playerStepReq.PlayerStep.Player = &playerO
				isO := CheckingIsStepObtained(ctx, playerStepReq)
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

func CheckWinStep(ctx context.Context, req *CheckWinStepReq) bool {
	var knownWinSteps []*Step

	winSteps := req.WinSteps
	playerSteps := req.PlayerSteps
	pStep := req.PlayerStep
	dimension := req.Dimension

	for _, winStep := range winSteps {
		knownWinSteps = []*Step{}
		for _, ws := range winStep {
			var knownPlayerStep *Step

			for _, ps := range playerSteps[*pStep.Player] {
				if ValidateStep(ctx, &ValidateStepReq{
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

		if len(knownWinSteps) == int(dimension.Current) {
			return true
		}
	}

	return false
}
