# TICTACTOE v3.2.2

Simple golang package/ library for custom tic-tac-toe game

## Installation
```
go get github.com/goplay-dev/go-tactictoe/v3
```

## Available Api

```go
import tictactoe "github.com/goplay-dev/go-tactictoe/v3"

var game = &tictactoe.GameConfig{
    Dimension:       dimension,
}
```

- InitGame
  ```go
  game.InitGame(context.Background())
  ```

- ValidateSteps
  ```go
  currentPlayer := tictactoe.X
  step := &tictactoe.Step{CX: 0, CY: 1}

  valid, win := game.ValidateSteps(ctx, &tictactoe.PlayerStepReq{
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