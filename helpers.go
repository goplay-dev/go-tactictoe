package main

import (
	"context"
	"fmt"
)

func PrintAvailableSteps(ctx context.Context, availableSteps []*step) {
	fmt.Print("available steps: ")
	for _, availStep := range availableSteps {
		fmt.Print(fmt.Sprintf("%v, ", availStep))
	}

	fmt.Println()
}

func PrintPlayerSteps(ctx context.Context, player *player, playerSteps map[player][]*step) {
	fmt.Print(fmt.Sprintf("%s steps: ", player.String()))
	for _, pStep := range playerSteps[*player] {
		fmt.Print(fmt.Sprintf("%v, ", pStep))
	}

	fmt.Println()
}

func PrintWinSteps(ctx context.Context, winSteps [][]*step) {
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
