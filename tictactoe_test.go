package tictactoe

import (
	"context"
	"os"
	"reflect"
	"testing"
)

var game = &GameConfig{
	Dimension:       nil,
	WinSteps:        nil,
	ActualPositions: nil,
}

func TestMain(m *testing.M) {
	game.Dimension = &Dimension{
		Current: 3,
		Min:     3,
		Max:     25,
	}

	game.WinSteps = WinSteps{
		Hor:   make(StepsList, game.Dimension.Current),
		Ver:   make(StepsList, game.Dimension.Current),
		LDiag: make(StepsList, 1),
		RDiag: make(StepsList, 1),
	}

	game.InitGame(context.Background())
	os.Exit(m.Run())
}

func TestSetupWinSteps(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want map[WinPos][][]Step
	}{
		{
			name: "3x3",
			args: args{
				ctx: context.Background(),
			},
			want: map[WinPos][][]Step{
				Hor: {
					{{0, 0}, {1, 0}, {2, 0}},
					{{0, 1}, {1, 1}, {2, 1}},
					{{0, 2}, {1, 2}, {2, 2}},
				},
				Ver: {
					{{0, 0}, {0, 1}, {0, 2}},
					{{1, 0}, {1, 1}, {1, 2}},
					{{2, 0}, {2, 1}, {2, 2}},
				},
				LDiag: {{{0, 0}, {1, 1}, {2, 2}}},
				RDiag: {{{2, 0}, {1, 1}, {0, 2}}},
			},
		},
	}
	for _, tt := range tests {
		game.setupWinSteps(tt.args.ctx)
		actual := map[WinPos][][]Step{
			Hor:   {},
			Ver:   {},
			LDiag: {},
			RDiag: {},
		}

		for key, val := range game.WinSteps {
			for _, pos := range val {
				var steps []Step
				for _, step := range pos {
					steps = append(steps, *step)
				}
				actual[key] = append(actual[key], steps)
			}
		}

		if !reflect.DeepEqual(actual, tt.want) {
			t.Errorf("the actual not match with expected value,\n"+
				" actual = %v,\n"+
				" want = %v", actual, tt.want)
		}
	}
}

func TestSetupHorWinSteps(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want [][]Step
	}{
		{
			name: "Hor 3x3",
			args: args{
				ctx: context.Background(),
			},
			want: [][]Step{
				{{0, 0}, {1, 0}, {2, 0}},
				{{0, 1}, {1, 1}, {2, 1}},
				{{0, 2}, {1, 2}, {2, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game.setupHorWinSteps(tt.args.ctx)
			var actual [][]Step

			for _, aSteps := range game.WinSteps[Hor] {
				var ws []Step
				for _, as := range aSteps {
					ws = append(ws, *as)
				}

				actual = append(actual, ws)
			}

			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("the actual not match with expected value,\n"+
					" actual = %v,\n"+
					" want = %v", actual, tt.want)
			}
		})
	}
}

func TestSetupVerWinSteps(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want [][]Step
	}{
		{
			name: "Hor 3x3",
			args: args{
				ctx: context.Background(),
			},
			want: [][]Step{
				{{0, 0}, {0, 1}, {0, 2}},
				{{1, 0}, {1, 1}, {1, 2}},
				{{2, 0}, {2, 1}, {2, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game.setupVerWinSteps(tt.args.ctx)
			var actual [][]Step

			for _, aSteps := range game.WinSteps[Ver] {
				var ws []Step
				for _, as := range aSteps {
					ws = append(ws, *as)
				}

				actual = append(actual, ws)
			}

			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("the actual not match with expected value,\n"+
					" actual = %v,\n"+
					" want = %v", actual, tt.want)
			}
		})
	}
}

func TestSetupLDiagWinSteps(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want [][]Step
	}{
		{
			name: "Left Diag 3x3",
			args: args{
				ctx: context.Background(),
			},
			want: [][]Step{
				{{0, 0}, {1, 1}, {2, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game.setupLDiagWinSteps(tt.args.ctx)
			var actual [][]Step

			for _, aSteps := range game.WinSteps[LDiag] {
				var ws []Step
				for _, as := range aSteps {
					ws = append(ws, *as)
				}

				actual = append(actual, ws)
			}

			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("the actual not match with expected value,\n"+
					" actual = %v,\n"+
					" want = %v", actual, tt.want)
			}
		})
	}
}

func TestSetupRDiagWinSteps(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want [][]Step
	}{
		{
			name: "Left Diag 3x3",
			args: args{
				ctx: context.Background(),
			},
			want: [][]Step{
				{{2, 0}, {1, 1}, {0, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game.setupRDiagWinSteps(tt.args.ctx)
			var actual [][]Step

			for _, aSteps := range game.WinSteps[RDiag] {
				var ws []Step
				for _, as := range aSteps {
					ws = append(ws, *as)
				}

				actual = append(actual, ws)
			}

			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("the actual not match with expected value,\n"+
					" actual = %v,\n"+
					" want = %v", actual, tt.want)
			}
		})
	}
}

func Test_gameConfig_ValidateAvailableStep(t *testing.T) {
	type args struct {
		ctx  context.Context
		step *Step
	}
	tests := []struct {
		name            string
		args            args
		ActualPositions ActualPositions
		want            bool
	}{
		{
			name: "return false",
			args: args{
				ctx: context.Background(),
				step: &Step{
					CX: 1,
					CY: 2,
				},
			},
			ActualPositions: ActualPositions{
				{E.String(), E.String(), E.String()},
				{E.String(), E.String(), E.String()},
				{E.String(), X.String(), E.String()},
			},
			want: false,
		},
		{
			name: "return true",
			args: args{
				ctx: context.Background(),
				step: &Step{
					CX: 2,
					CY: 0,
				},
			},
			ActualPositions: ActualPositions{
				{E.String(), E.String(), E.String()},
				{E.String(), E.String(), E.String()},
				{E.String(), E.String(), E.String()},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game.ActualPositions = tt.ActualPositions
			if got := game.validateAvailableStep(tt.args.ctx, tt.args.step); got != tt.want {
				t.Errorf("ValidatePlayerStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameConfig_SetupActualPos(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want ActualPositions
	}{
		{
			name: "all pos are nil",
			args: args{
				ctx: context.Background(),
			},
			want: ActualPositions{
				{E.String(), E.String(), E.String()},
				{E.String(), E.String(), E.String()},
				{E.String(), E.String(), E.String()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game.setupActualPos(tt.args.ctx)

			if !reflect.DeepEqual(game.ActualPositions, tt.want) {
				t.Errorf("the actual not match with expected value,\n"+
					" actual = %v,\n"+
					" want = %v", game.ActualPositions, tt.want)
			}
		})
	}
}

func TestGameConfig_SaveActualPos(t *testing.T) {
	xPlayer := X

	type args struct {
		ctx context.Context
		req *PlayerStepReq
	}
	tests := []struct {
		name string
		args args
		want ActualPositions
	}{
		{
			name: "mark X at 0,2",
			args: args{
				ctx: context.Background(),
				req: &PlayerStepReq{
					Player: &xPlayer,
					Step: &Step{
						CX: 0,
						CY: 2,
					},
				},
			},
			want: ActualPositions{
				{E.String(), E.String(), E.String()},
				{E.String(), E.String(), E.String()},
				{X.String(), E.String(), E.String()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game.saveActualPos(tt.args.ctx, tt.args.req)

			if !reflect.DeepEqual(game.ActualPositions, tt.want) {
				t.Errorf("the actual not match with expected value,\n"+
					" actual = %v,\n"+
					" want = %v", game.ActualPositions, tt.want)
			}
		})
	}
}

func TestGameConfig_ValidateWinStep(t *testing.T) {
	playerX := X
	playerO := O

	type args struct {
		ctx    context.Context
		player *Player
	}
	tests := []struct {
		name            string
		args            args
		actualPositions ActualPositions
		want            bool
	}{
		{
			name: "X Win",
			args: args{
				ctx:    context.Background(),
				player: &playerX,
			},
			actualPositions: ActualPositions{
				{E.String(), E.String(), X.String()},
				{E.String(), X.String(), E.String()},
				{X.String(), E.String(), E.String()},
			},
			want: true,
		},
		{
			name: "O Not Win yet",
			args: args{
				ctx:    context.Background(),
				player: &playerO,
			},
			actualPositions: ActualPositions{
				{O.String(), E.String(), E.String()},
				{E.String(), X.String(), E.String()},
				{X.String(), E.String(), O.String()},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game.ActualPositions = tt.actualPositions
			if got := game.validateWinStep(tt.args.ctx, tt.args.player); got != tt.want {
				t.Errorf("CheckingWinStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameConfig_validatePlayer(t *testing.T) {
	type args struct {
		ctx    context.Context
		player Player
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "validate player O",
			args: args{
				ctx:    context.Background(),
				player: O,
			},
			want: true,
		},
		{
			name: "validate unknown player",
			args: args{
				ctx:    context.Background(),
				player: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := game.validatePlayer(tt.args.ctx, tt.args.player); got != tt.want {
				t.Errorf("validatePlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameConfig_validateStepRange(t *testing.T) {
	type args struct {
		ctx  context.Context
		step *Step
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "false step",
			args: args{
				ctx: context.Background(),
				step: &Step{
					CX: -1,
					CY: 0,
				},
			},
			want: false,
		},
		{
			name: "true step",
			args: args{
				ctx: context.Background(),
				step: &Step{
					CX: 1,
					CY: 2,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := game.validateStepRange(tt.args.ctx, tt.args.step); got != tt.want {
				t.Errorf("validateStepRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
