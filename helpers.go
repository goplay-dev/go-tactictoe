package tictactoe

import (
	"context"
	"fmt"
)

func PrintAvailableSteps(ctx context.Context, availableSteps []*Step) {
	fmt.Print("available steps: ")
	for _, availStep := range availableSteps {
		fmt.Print(fmt.Sprintf("%v, ", availStep))
	}

	fmt.Println()
}

func PrintPlayerSteps(ctx context.Context, player *player, playerSteps map[player][]*Step) {
	fmt.Print(fmt.Sprintf("%s steps: ", player.String()))
	for _, pStep := range playerSteps[*player] {
		fmt.Print(fmt.Sprintf("%v, ", pStep))
	}

	fmt.Println()
}

func PrintWinSteps(ctx context.Context, winSteps [][]*Step) {
	fmt.Println("win steps: ")
	for i, ws := range winSteps {
		fmt.Print(fmt.Sprintf("%d: ", i))
		for _, s := range ws {
			fmt.Print(fmt.Sprintf("%v, ", s))
		}
		fmt.Println()
	}

	fmt.Println()
}

func PrintActualPos(ctx context.Context, positions [][]string, currDimension int32) {
	for cy := int32(0); cy < currDimension; cy++ {
		for cx := int32(0); cx < currDimension; cx++ {
			fmt.Print(positions[cy][cx])
		}
		fmt.Println()
	}
}
