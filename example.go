package tictactoe

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

func (p Player) String() string {
	switch p {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return ""
	}
}

func ConsolePlay(ctx context.Context) {
	var currentPlayer Player
	var stepCoord string
	var stepCX int32
	var stepCY int32
	var dimension = &Dimension{
		Current: 25,
		Min:     3,
		Max:     25,
	}
	var game = &GameConfig{
		Dimension:       dimension,
		WinSteps:        nil,
		ActualPositions: nil,
	}

	fmt.Print(fmt.Sprintf("Input Dimension (%d - %d): ", dimension.Min, dimension.Max))
	fmt.Scanf("%d", &dimension.Current)

	fmt.Println()
	fmt.Println("==================================================")
	fmt.Println()

	game.InitGame(ctx)

INPUTPLAYER:
	fmt.Print("Input Player (0 or 1): ")
	fmt.Scanf("%d", &currentPlayer)

	if currentPlayer > 1 || currentPlayer < 0 {
		fmt.Println("unknown player !")
		goto INPUTPLAYER
	}

STEP:
	PrintActualPos(ctx, game.ActualPositions, game.Dimension.Current)
	fmt.Print(fmt.Sprintf("(player %s) Input CX,CY step (0 - %d) (ex: 1,2): ",
		currentPlayer.String(), game.Dimension.Current-1))
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

	if stepCX > (game.Dimension.Current-1) || stepCX < 0 {
		fmt.Println(fmt.Sprintf("CX must between 0 - %d", game.Dimension.Current-1))
		goto STEP
	}

	if stepCY > (game.Dimension.Current-1) || stepCY < 0 {
		fmt.Println(fmt.Sprintf("CY must between 0 - %d", game.Dimension.Current-1))
		goto STEP
	}

	step := &Step{
		CX: stepCX,
		CY: stepCY,
	}

	valid, win := game.ValidateSteps(ctx, &PlayerStepReq{
		Player: &currentPlayer,
		Step:   step,
	})

	if !valid {
		fmt.Println("steps not available !")
		goto STEP
	} else if valid {
		PrintActualPos(ctx, game.ActualPositions, game.Dimension.Current)
		if win {
			fmt.Println(fmt.Sprintf("%s Win !!!", currentPlayer.String()))
			return
		}
	}

	if currentPlayer == X {
		currentPlayer = O
	} else {
		currentPlayer = X
	}

	goto STEP
}
