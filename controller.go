package main

import (
	"context"
	"fmt"
)

type ValidateStepReq struct {
	Step1 *step
	Step2 *step
}

func ValidatePlayerStep(ctx context.Context, pStep *playerStep, availableSteps []*step) bool {
	for index, availStep := range availableSteps {
		if ValidateStep(ctx, &ValidateStepReq{
			Step1: pStep.Step,
			Step2: availStep,
		}) {
			RemoveSelectedStep(ctx, index, availableSteps)
			return true
		}
	}

	return false
}

func RemoveSelectedStep(ctx context.Context, index int, availableSteps []*step) {
	availableSteps = append(availableSteps[:index], availableSteps[index+1:]...)
}

func SaveStep(ctx context.Context, pStep *playerStep, playerSteps map[player][]*step) {
	playerSteps[*pStep.Player] = append(playerSteps[*pStep.Player], pStep.Step)
}

func Draw(ctx context.Context, dimension int32) {
	for cy := 0; cy < int(dimension); cy++ {
		for cx := 0; cx < int(dimension); cx++ {
			mark := "-"

			stepPoint := &step{
				CX: int32(cx),
				CY: int32(cy),
			}

			playerX := X
			playerO := O

			isX := CheckingIsStepObtained(ctx, &playerStep{
				Player: &playerX,
				Step:   stepPoint,
			}, playerSteps)
			if isX {
				mark = X.String()
			} else {
				isO := CheckingIsStepObtained(ctx, &playerStep{
					Player: &playerO,
					Step:   stepPoint,
				}, playerSteps)
				if isO {
					mark = O.String()
				}
			}

			fmt.Print(fmt.Sprintf(" %s ", mark))
		}

		fmt.Println()
	}
}

func ValidateStep(ctx context.Context, req *ValidateStepReq) bool {
	if req.Step1.CX == req.Step2.CX && req.Step1.CY == req.Step2.CY {
		return true
	}

	return false
}

func CheckingIsStepObtained(ctx context.Context, pStep *playerStep, playerSteps map[player][]*step) bool {
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

func CheckWinStep(
	ctx context.Context,
	pStep *playerStep,
	winSteps [][]*step,
	dimension int32) bool {

	var knownWinSteps []*step

	for _, winStep := range winSteps {
		knownWinSteps = []*step{}
		for _, ws := range winStep {
			var knownPlayerStep *step

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

		if len(knownWinSteps) == int(dimension) {
			return true
		}
	}

	return false
}

func SetupAvailableSteps(ctx context.Context, dimension int32) []*step {
	var avSteps []*step

	for cy := int32(0); cy < dimension; cy++ {
		for cx := int32(0); cx < dimension; cx++ {
			avSteps = append(avSteps, &step{CX: cx, CY: cy})
		}
	}

	return avSteps
}

func SetupWinSteps(ctx context.Context, dimension int32) [][]*step {
	var wSteps [][]*step

	wSteps = append(SetupHorVerWinSteps(ctx, dimension), SetupDiagWinSteps(ctx, dimension)...)
	return wSteps
}

func SetupHorVerWinSteps(ctx context.Context, dimension int32) [][]*step {
	var wSteps [][]*step

	for cy := int32(0); cy < dimension; cy++ {
		var hSteps []*step
		var vSteps []*step

		for cx := int32(0); cx < dimension; cx++ {
			hSteps = append(hSteps, &step{CX: cx, CY: cy})
			vSteps = append(vSteps, &step{CX: cy, CY: cx})
		}

		wSteps = append(wSteps, hSteps, vSteps)
	}

	return wSteps
}

func SetupDiagWinSteps(ctx context.Context, dimension int32) [][]*step {
	var wSteps [][]*step
	var rdSteps []*step
	var ldSteps []*step

	for cy := dimension - 1; cy >= 0; cy-- {
		rdSteps = append(rdSteps, &step{CX: cy, CY: cy})
		ldSteps = append(ldSteps, &step{CX: (dimension - 1) - cy, CY: cy})
	}

	wSteps = append(wSteps, rdSteps, ldSteps)

	return wSteps
}
