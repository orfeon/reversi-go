# reversi-go
This library provides reversi common functions such as search and board management.
You can develop computer reversi player simply.
You should only implement board evaluation logic.

## Install

```
go get -u github.com/orfeon/reversi-go
```

## Get Start

```go
package main

import (
	reversi "github.com/orfeon/reversi-go"
)

func main() {

  board := reversi.NewBoard()

  player1 := reversi.NewComputerPlayer(reversi.STONE_BLACK)
  player2 := reversi.NewComputerPlayer(reversi.STONE_WHITE)

  var pos Pos

  for !board.CheckGameover() {
    pos = player1.Think(*board)
    board.Move(pos.Index, STONE_BLACK)
    fmt.Println(pos, board.CountStone(STONE_BLACK), board.CountStone(STONE_WHITE))
    pos = player2.Think(*board)
    board.Move(pos.Index, STONE_WHITE)
    fmt.Println(pos, board.CountStone(STONE_BLACK), board.CountStone(STONE_WHITE))
  }
}
```
