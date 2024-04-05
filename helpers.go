package tictactoe

import (
	"context"
	"fmt"
)

func PrintWinSteps(ctx context.Context, winSteps WinSteps) {
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

func PrintActualPos(ctx context.Context, positions ActualPositions, currDimension int32) {
	for cy := int32(0); cy < currDimension; cy++ {
		for cx := int32(0); cx < currDimension; cx++ {
			fmt.Print(positions[cy][cx])
		}
		fmt.Println()
	}
}
