package tictactoe

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

func ConsolePlay(ctx context.Context) {
INIT:
	var dimension = &Dimension{
		Current: 25,
		Min:     3,
		Max:     25,
	}

	fmt.Print(fmt.Sprintf("Input Dimension (%d - %d): ", dimension.Min, dimension.Max))
	fmt.Scanf("%d", &dimension.Current)

	game, err := InitGame(ctx, &GameConfig{
		Dimension: dimension,
	})
	if err != nil {
		fmt.Println(err)
		goto INIT
	}

	fmt.Println()
	fmt.Println("==================================================")
	fmt.Println()

	var currentPlayer player
	var stepCoord string
	var stepCX int32
	var stepCY int32
INPUTPLAYER:
	fmt.Print("Input Player (0 or 1): ")
	fmt.Scanf("%d", &currentPlayer)

	if currentPlayer > 1 || currentPlayer < 0 {
		fmt.Println("unknown player !")
		goto INPUTPLAYER
	}

STEP:
	PrintActualPos(ctx, game.GetActualPos(ctx), dimension.Current)

	fmt.Print(fmt.Sprintf("(player %s) Input CX,CY step (0 - %d) (ex: 1,2): ",
		currentPlayer.String(), dimension.Current-1))
	fmt.Scanf("%s", &stepCoord)

	splitCoord := strings.Split(stepCoord, ",")
	if len(splitCoord) < 2 {
		fmt.Println("wrong input !")
		goto STEP
	}

	stepX, _ := strconv.Atoi(splitCoord[0])
	stepY, _ := strconv.Atoi(splitCoord[1])

	stepCX = int32(stepX)
	stepCY = int32(stepY)

	if stepCX > (dimension.Current-1) || stepCX < 0 {
		fmt.Println(fmt.Sprintf("CX must between 0 - %d", dimension.Current-1))
		goto STEP
	}

	if stepCY > (dimension.Current-1) || stepCY < 0 {
		fmt.Println(fmt.Sprintf("CY must between 0 - %d", dimension.Current-1))
		goto STEP
	}

	stepsReq := &playerStep{
		Player: &currentPlayer,
		Step: &Step{
			CX: stepCX,
			CY: stepCY,
		},
	}

	if game.ValidatePlayerStep(ctx, stepsReq) {
		game.SaveStep(ctx, stepsReq, playerSteps)
		if game.CheckWinStep(ctx, stepsReq) {
			PrintActualPos(ctx, game.GetActualPos(ctx), dimension.Current)
			fmt.Println(fmt.Sprintf("%s Win !!!", currentPlayer.String()))
			return
		}
	} else {
		fmt.Println("steps not available !")
		goto STEP
	}

	if currentPlayer == X {
		currentPlayer = O
	} else {
		currentPlayer = X
	}

	goto STEP
}
