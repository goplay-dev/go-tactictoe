# TICTACTOE v3.0.0

Simple golang package/ library for custom tic-tac-toe game

## Available Api
- InitGame
    ```go
        var game = &GameConfig{
            Dimension:       dimension,
            WinSteps:        nil,
            ActualPositions: nil,
        }
        
        game.InitGame(context.Background())
    ```

- ValidateSteps
    ```go
        var game = &GameConfig{
            Dimension:       dimension,
            WinSteps:        nil,
            ActualPositions: nil,
        }
        
        currentPlayer := X
        step := &Step{CX: 0, CY: 1}
    
        valid, win := game.ValidateSteps(ctx, &PlayerStepReq{
            Player: &currentPlayer,
            Step:   step,
        })
  
        if !valid {
            fmt.Println("steps not available !")
            goto STEP
        } else if valid {
            if win {
                fmt.Println(fmt.Sprintf("%s Win !!!", currentPlayer.String()))
                return
            }
        }
    ```
  
## Test
```go 
  func ConsolePlay(ctx context.Context)
```