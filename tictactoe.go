package tictactoe

import (
	"context"
)

func (g *GameConfig) InitGame(ctx context.Context) {
	g.setupWinSteps(ctx).setupActualPos(ctx)
}

func (g *GameConfig) setupWinSteps(ctx context.Context) *GameConfig {
	g.WinSteps = WinSteps{
		Hor:   make(StepsList, g.Dimension.Current),
		Ver:   make(StepsList, g.Dimension.Current),
		LDiag: make(StepsList, 1),
		RDiag: make(StepsList, 1),
	}

	g.setupHorWinSteps(ctx).
		setupVerWinSteps(ctx).
		setupLDiagWinSteps(ctx).
		setupRDiagWinSteps(ctx)

	return g
}

func (g *GameConfig) setupHorWinSteps(ctx context.Context) *GameConfig {
	var currDimension = g.Dimension.Current

	for cy := int32(0); cy < currDimension; cy++ {
		var winStep = make(Steps, currDimension)
		for cx := int32(0); cx < currDimension; cx++ {
			winStep[cx] = &Step{
				CX: cx,
				CY: cy,
			}
		}

		g.WinSteps[Hor][cy] = winStep
	}

	return g
}

func (g *GameConfig) setupVerWinSteps(ctx context.Context) *GameConfig {
	var currDimension = g.Dimension.Current

	for cx := int32(0); cx < currDimension; cx++ {
		var winStep = make(Steps, currDimension)
		for cy := int32(0); cy < currDimension; cy++ {
			winStep[cy] = &Step{
				CX: cx,
				CY: cy,
			}
		}

		g.WinSteps[Ver][cx] = winStep
	}

	return g
}

func (g *GameConfig) setupLDiagWinSteps(ctx context.Context) *GameConfig {
	var currDimension = g.Dimension.Current
	var winStep = make(Steps, currDimension)

	for ld := int32(0); ld < currDimension; ld++ {
		winStep[ld] = &Step{
			CX: ld,
			CY: ld,
		}
	}

	g.WinSteps[LDiag][0] = winStep
	return g
}

func (g *GameConfig) setupRDiagWinSteps(ctx context.Context) *GameConfig {
	var currDimension = g.Dimension.Current
	var winStep = make(Steps, currDimension)

	for ld := currDimension - 1; ld >= 0; ld-- {
		winStep[ld] = &Step{
			CX: (currDimension - 1) - ld,
			CY: ld,
		}
	}

	g.WinSteps[RDiag][0] = winStep
	return g
}

func (g *GameConfig) ValidateSteps(ctx context.Context, req *PlayerStepReq) (valid bool, win bool) {
	valid = false
	win = false

	if g.validateAvailableStep(ctx, req) {
		valid = true

		if g.saveActualPos(ctx, req).validateWinStep(ctx, req.Player) {
			win = true
		}
	}

	return
}

func (g *GameConfig) validateAvailableStep(ctx context.Context, req *PlayerStepReq) bool {
	return g.ActualPositions[req.Step.CY][req.Step.CX] == nil
}

func (g *GameConfig) setupActualPos(ctx context.Context) *GameConfig {
	var currDimension = g.Dimension.Current
	g.ActualPositions = make(ActualPositions, currDimension)

	for cy := int32(0); cy < currDimension; cy++ {
		var actualPosition = make(Players, currDimension)
		for cx := int32(0); cx < currDimension; cx++ {
			actualPosition[cx] = nil
		}

		g.ActualPositions[cy] = actualPosition
	}

	return g
}

func (g *GameConfig) saveActualPos(ctx context.Context, req *PlayerStepReq) *GameConfig {
	g.ActualPositions[req.Step.CY][req.Step.CX] = req.Player
	return g
}

func (g *GameConfig) validateWinStep(ctx context.Context, player *Player) (win bool) {
	for _, val := range g.WinSteps {
		for _, pos := range val {
			for i, step := range pos {
				if g.ActualPositions[step.CX][step.CY] != player {
					break
				}

				if int32(i) == g.Dimension.Current-1 {
					return true
				}
			}
		}
	}

	return false
}
