package tictactoe

import (
	"context"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
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
		want [][]Step
	}{
		{
			name: "3x3",
			args: args{
				ctx:           context.Background(),
				currDimension: 3,
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
			got := SetupAvailableSteps(tt.args.ctx, tt.args.currDimension)
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
					AvailableSteps: [][]*Step{
						{{0, 0}, {1, 0}, {2, 0}},
						{{0, 1}, {1, 1}, {2, 1}},
						{{0, 2}, {1, 2}, {2, 2}},
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
					AvailableSteps: [][]*Step{
						{{0, 0}, {1, 0}, {2, 0}},
						{{0, 1}, {1, 1}, {2, 1}},
						{{0, 2}, {1, 2}, {2, 2}},
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
					AvailableSteps: [][]*Step{
						{{0, 0}, {1, 0}, {2, 0}},
						{{0, 1}, {1, 1}, {2, 1}},
						{{0, 2}, {1, 2}, {2, 2}},
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
					AvailableSteps: [][]*Step{
						{{0, 0}, {1, 0}, {2, 0}},
						{{0, 1}, {1, 1}, {2, 1}},
						{{0, 2}, {1, 2}, {2, 2}},
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
					AvailableSteps: [][]*Step{
						{{0, 0}, {1, 0}, {2, 0}},
						{{0, 1}, {1, 1}, {2, 1}},
						{{0, 2}, {1, 2}, {2, 2}},
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

func Test_gameConfig_ValidatePlayerStep(t *testing.T) {
	playerMark := X

	type fields struct {
		GameConfig *GameConfig
	}
	type args struct {
		ctx   context.Context
		pStep *playerStep
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "return false",
			fields: fields{GameConfig: &GameConfig{
				Dimension: &Dimension{
					Current: 3,
					Min:     3,
					Max:     25,
				},
				AvailableSteps: [][]*Step{
					{{0, 0}, {1, 0}, {2, 0}},
					{{0, 1}, {2, 1}},
					{{0, 2}, {1, 2}, {2, 2}},
				},
				WinSteps: nil,
			}},
			args: args{
				ctx: context.Background(),
				pStep: &playerStep{
					Player: &playerMark,
					Step: &Step{
						CX: 1,
						CY: 1,
					},
				},
			},
			want: false,
		},
		{
			name: "return true",
			fields: fields{GameConfig: &GameConfig{
				Dimension: &Dimension{
					Current: 3,
					Min:     3,
					Max:     25,
				},
				AvailableSteps: [][]*Step{
					{{0, 0}, {1, 0}, {2, 0}},
					{{0, 1}, {2, 1}},
					{{0, 2}, {1, 2}, {2, 2}},
				},
				WinSteps: nil,
			}},
			args: args{
				ctx: context.Background(),
				pStep: &playerStep{
					Player: &playerMark,
					Step: &Step{
						CX: 2,
						CY: 0,
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gameConfig{
				GameConfig: tt.fields.GameConfig,
			}
			if got := g.ValidatePlayerStep(tt.args.ctx, tt.args.pStep); got != tt.want {
				t.Errorf("ValidatePlayerStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gameConfig_RemoveSelectedStep(t *testing.T) {
	type fields struct {
		GameConfig *GameConfig
	}
	type args struct {
		ctx  context.Context
		step *Step
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   [][]Step
	}{
		{
			name: "remove step {2,1}",
			fields: fields{GameConfig: &GameConfig{
				Dimension: &Dimension{
					Current: 3,
					Min:     3,
					Max:     25,
				},
				AvailableSteps: [][]*Step{
					{{0, 0}, {1, 0}, {2, 0}},
					{{0, 1}, {1, 1}, {2, 1}},
					{{0, 2}, {1, 2}, {2, 2}},
				},
				WinSteps: nil,
			}},
			args: args{
				ctx: context.Background(),
				step: &Step{
					CX: 2,
					CY: 1,
				},
			},
			want: [][]Step{
				{{0, 0}, {1, 0}, {2, 0}},
				{{0, 1}, {1, 1}},
				{{0, 2}, {1, 2}, {2, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := InitGame(context.Background(), tt.fields.GameConfig)
			if err != nil {
				t.Fatal(err)
			}

			g.RemoveSelectedStep(tt.args.ctx, tt.args.step)

			var actual [][]Step

			for _, aSteps := range tt.fields.GameConfig.AvailableSteps {
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
