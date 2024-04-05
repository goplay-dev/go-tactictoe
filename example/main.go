package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	tictactoe "github.com/michaelwp/go-tactictoe/v3"
)

func main() {
	ConsolePlay(context.Background())
}

func ConsolePlay(ctx context.Context) {
INITGAME:
	var currentPlayer tictactoe.Player
	var stepCoord string
	var stepCX int32
	var stepCY int32
	var dimension = &tictactoe.Dimension{
		Current: 25,
		Min:     3,
		Max:     25,
	}
	var game = &tictactoe.GameConfig{
		Dimension: dimension,
	}

	fmt.Print(fmt.Sprintf("Input Dimension (%d - %d): ", dimension.Min, dimension.Max))
	fmt.Scanf("%d", &dimension.Current)

	fmt.Println()
	fmt.Println("==================================================")
	fmt.Println()

	err := game.InitGame(ctx)
	if err != nil {
		fmt.Println(err)
		goto INITGAME
	}

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

	step := &tictactoe.Step{
		CX: stepCX,
		CY: stepCY,
	}

	valid, win := game.ValidateSteps(ctx, &tictactoe.PlayerStepReq{
		Player: &currentPlayer,
		Step:   step,
	})

	if !valid {
		fmt.Println("steps not available !")
		goto STEP
	} else if valid {
		if win {
			PrintActualPos(ctx, game.ActualPositions, game.Dimension.Current)
			fmt.Println(fmt.Sprintf("%s Win !!!", currentPlayer.String()))
			return
		}
	}

	if currentPlayer == tictactoe.X {
		currentPlayer = tictactoe.O
	} else {
		currentPlayer = tictactoe.X
	}

	goto STEP
}

func PrintActualPos(ctx context.Context, positions tictactoe.ActualPositions, currDimension int32) {
	for cy := int32(0); cy < currDimension; cy++ {
		for cx := int32(0); cx < currDimension; cx++ {
			var mark = positions[cy][cx]

			if mark == tictactoe.E.String() {
				mark = "-"
			}

			fmt.Print(fmt.Sprintf("%v ", mark))
		}
		fmt.Println()
	}
}

func PrintWinSteps(ctx context.Context, winSteps tictactoe.WinSteps) {
	fmt.Println("win steps: ")
	i := 0
	for winPos := range winSteps {
		i++
		fmt.Print(fmt.Sprintf("%d: ", i))
		for _, s := range winPos {
			fmt.Print(fmt.Sprintf("%v, ", s))
		}
		fmt.Println()
	}

	fmt.Println()
}
