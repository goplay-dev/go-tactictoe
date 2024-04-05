package tictactoe

import (
	"context"
	"testing"
)

func TestPrintActualPos(t *testing.T) {
	playerX := X
	playerO := O

	type args struct {
		ctx           context.Context
		positions     ActualPositions
		currDimension int32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nil pos",
			args: args{
				ctx: context.Background(),
				positions: ActualPositions{
					{nil, nil, nil},
					{nil, &playerX, nil},
					{nil, nil, &playerO},
				},
				currDimension: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintActualPos(tt.args.ctx, tt.args.positions, tt.args.currDimension)
		})
	}
}
