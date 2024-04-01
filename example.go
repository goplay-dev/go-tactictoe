package tictactoe

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

func ConsolePlay(ctx context.Context) {
	var currentPlayer Player
	var stepCoord string
	var stepCX int32
	var stepCY int32
	var playerSteps PlayerSteps
	var dimension = &Dimension{
		Current: 25,
		Min:     3,
		Max:     25,
	}

	fmt.Print(fmt.Sprintf("Input Dimension (%d - %d): ", dimension.Min, dimension.Max))
	fmt.Scanf("%d", &dimension.Current)

	fmt.Println()
	fmt.Println("==================================================")
	fmt.Println()

	availableSteps := SetupAvailableSteps(ctx, dimension.Current)
	winSteps := SetupWinSteps(ctx, &GameConfig{
		Dimension:      dimension,
		AvailableSteps: availableSteps,
	})

INPUTPLAYER:
	fmt.Print("Input Player (0 or 1): ")
	fmt.Scanf("%d", &currentPlayer)

	if currentPlayer > 1 || currentPlayer < 0 {
		fmt.Println("unknown player !")
		goto INPUTPLAYER
	}

STEP:
	PrintActualPos(ctx, GetActualPos(ctx, &GetActualPosReq{
		Dimension:   dimension,
		PlayerSteps: playerSteps,
	}), dimension.Current)

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

	step := &Step{
		CX: stepCX,
		CY: stepCY,
	}

	playerStep := &PlayerStep{
		Player: &currentPlayer,
		Step:   step,
	}

	if ValidatePlayerStep(ctx, &ValidatePlayerStepReq{
		PlayerStep:     playerStep,
		AvailableSteps: availableSteps,
	}) {
		playerSteps = SavePlayerStep(ctx, &PlayerStepReq{
			PlayerStep:  playerStep,
			PlayerSteps: playerSteps,
		})
		if CheckWinStep(ctx, &CheckWinStepReq{
			PlayerStep:  playerStep,
			WinSteps:    winSteps,
			PlayerSteps: playerSteps,
			Dimension:   dimension,
		}) {
			PrintActualPos(ctx, GetActualPos(ctx, &GetActualPosReq{
				Dimension:   dimension,
				PlayerSteps: playerSteps,
			}), dimension.Current)
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
