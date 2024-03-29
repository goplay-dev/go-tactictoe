package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type InitStep struct {
	Dimension      int32
	AvailableSteps []*step
	WinSteps       [][]*step
}

func main() {
	Play(context.TODO())
}

func Play(ctx context.Context) {
INIT:
	initSteps, err := Init(ctx)
	if err != nil {
		fmt.Println(err)
		goto INIT
	}

	fmt.Println("initSteps:", initSteps.Dimension)

START:
	fmt.Println()
	fmt.Println("==================================================")
	fmt.Println()

	var currentPlayer player
	var stepCoord string
	var stepCX int32
	var stepCY int32

	Draw(ctx, initSteps.Dimension)

INPUTPLAYER:
	fmt.Print("Input Player (0 or 1): ")
	fmt.Scanf("%d", &currentPlayer)

	if currentPlayer > 1 || currentPlayer < 0 {
		fmt.Println("unknown player !")
		goto INPUTPLAYER
	}

STEP:
	fmt.Print(fmt.Sprintf("Input CX,CY step (0 - %d) (ex: 1,2): ", initSteps.Dimension-1))
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

	if stepCX > (initSteps.Dimension-1) || stepCX < 0 {
		fmt.Println(fmt.Sprintf("CX must between 0 - %d", initSteps.Dimension-1))
		goto STEP
	}

	if stepCY > (initSteps.Dimension-1) || stepCY < 0 {
		fmt.Println(fmt.Sprintf("CY must between 0 - %d", initSteps.Dimension-1))
		goto STEP
	}

	stepsReq := &playerStep{
		Player: &currentPlayer,
		Step: &step{
			CX: stepCX,
			CY: stepCY,
		},
	}

	if ValidatePlayerStep(ctx, stepsReq, initSteps.AvailableSteps) {
		SaveStep(ctx, stepsReq, playerSteps)
		if CheckWinStep(ctx, stepsReq, initSteps.WinSteps, initSteps.Dimension) {
			Draw(ctx, initSteps.Dimension)
			fmt.Println(fmt.Sprintf("%s Win !!!", currentPlayer.String()))
			return
		}
	} else {
		fmt.Println("steps not available !")
		goto STEP
	}

	goto START
}

func Init(ctx context.Context) (*InitStep, error) {
	var dimension int32 = 3

	fmt.Print("Input Dimension (3 - 10): ")
	fmt.Scanf("%d", &dimension)

	if dimension < 3 || dimension > 10 {
		return nil, fmt.Errorf("dimension must between 3 - 10 ")
	}

	return &InitStep{
		AvailableSteps: SetupAvailableSteps(ctx, dimension),
		WinSteps:       SetupWinSteps(ctx, dimension),
		Dimension:      dimension,
	}, nil
}
