package tictactoe

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"
)

var game InitGameConfig

func TestMain(m *testing.M) {
	ctx := context.Background()
	var err error

	var dimension = &Dimension{
		Current: 25,
		Min:     3,
		Max:     25,
	}

	game, err = InitGame(ctx, &GameConfig{
		Dimension: dimension,
	})
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestSetupAvailableSteps(t *testing.T) {
	type args struct {
		ctx           context.Context
		currDimension int32
	}
	tests := []struct {
		name string
		args args
		want []Step
	}{
		{
			name: "3x3",
			args: args{
				ctx:           context.Background(),
				currDimension: 3,
			},
			want: []Step{
				{0, 0}, {1, 0}, {2, 0},
				{0, 1}, {1, 1}, {2, 1},
				{0, 2}, {1, 2}, {2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetupAvailableSteps(tt.args.ctx, tt.args.currDimension)
			var actual []Step

			for _, as := range got {
				actual = append(actual, *as)
			}

			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("the actual not match with expected value,\n"+
					" actual = %v,\n"+
					" want = %v", actual, tt.want)
			}
		})
	}
}

func TestSetupWinSteps(t *testing.T) {
	type args struct {
		ctx    context.Context
		config *GameConfig
	}
	tests := []struct {
		name string
		args args
		want [][]Step
	}{
		{
			name: "3x3",
			args: args{
				ctx: context.Background(),
				config: &GameConfig{
					Dimension: &Dimension{
						Current: 3,
						Min:     3,
						Max:     25,
					},
					AvailableSteps: []*Step{
						{0, 0}, {1, 0}, {2, 0},
						{0, 1}, {1, 1}, {2, 1},
						{0, 2}, {1, 2}, {2, 2},
					},
				},
			},
			want: [][]Step{
				{{0, 0}, {1, 0}, {2, 0}},
				{{0, 1}, {1, 1}, {2, 1}},
				{{0, 2}, {1, 2}, {2, 2}},

				{{0, 0}, {0, 1}, {0, 2}},
				{{1, 0}, {1, 1}, {1, 2}},
				{{2, 0}, {2, 1}, {2, 2}},

				{{0, 0}, {1, 1}, {2, 2}},
				{{2, 0}, {1, 1}, {0, 2}},
			},
		},
	}
	for _, tt := range tests {
		got := SetupWinSteps(tt.args.ctx, tt.args.config)
		var actual [][]Step

		for _, aSteps := range got {
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
	}
}

func TestSetupLDiagWinSteps(t *testing.T) {
	type args struct {
		ctx    context.Context
		config *GameConfig
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
				config: &GameConfig{
					Dimension: &Dimension{
						Current: 3,
						Min:     3,
						Max:     25,
					},
					AvailableSteps: []*Step{
						{0, 0}, {1, 0}, {2, 0},
						{0, 1}, {1, 1}, {2, 1},
						{0, 2}, {1, 2}, {2, 2},
					},
					WinSteps: nil,
				},
			},
			want: [][]Step{
				{{0, 0}, {1, 1}, {2, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetupLDiagWinSteps(tt.args.ctx, tt.args.config)
			var actual [][]Step

			for _, aSteps := range got {
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
		ctx    context.Context
		config *GameConfig
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
				config: &GameConfig{
					Dimension: &Dimension{
						Current: 3,
						Min:     3,
						Max:     25,
					},
					AvailableSteps: []*Step{
						{0, 0}, {1, 0}, {2, 0},
						{0, 1}, {1, 1}, {2, 1},
						{0, 2}, {1, 2}, {2, 2},
					},
					WinSteps: nil,
				},
			},
			want: [][]Step{
				{{2, 0}, {1, 1}, {0, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetupRDiagWinSteps(tt.args.ctx, tt.args.config)
			var actual [][]Step

			for _, aSteps := range got {
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

func TestSetupHorWinSteps(t *testing.T) {
	type args struct {
		ctx    context.Context
		config *GameConfig
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
				config: &GameConfig{
					Dimension: &Dimension{
						Current: 3,
						Min:     3,
						Max:     25,
					},
					AvailableSteps: []*Step{
						{0, 0}, {1, 0}, {2, 0},
						{0, 1}, {1, 1}, {2, 1},
						{0, 2}, {1, 2}, {2, 2},
					},
					WinSteps: nil,
				},
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
			got := SetupHorWinSteps(tt.args.ctx, tt.args.config)
			var actual [][]Step

			for _, aSteps := range got {
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
		ctx    context.Context
		config *GameConfig
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
				config: &GameConfig{
					Dimension: &Dimension{
						Current: 3,
						Min:     3,
						Max:     25,
					},
					AvailableSteps: []*Step{
						{0, 0}, {1, 0}, {2, 0},
						{0, 1}, {1, 1}, {2, 1},
						{0, 2}, {1, 2}, {2, 2},
					},
					WinSteps: nil,
				},
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
			got := SetupVerWinSteps(tt.args.ctx, tt.args.config)
			var actual [][]Step

			for _, aSteps := range got {
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
