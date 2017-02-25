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
	"fmt"
	reversi "github.com/orfeon/reversi-go"
)

func main() {

  board := reversi.NewBoard()

  player1 := reversi.NewComputerPlayer(reversi.STONE_BLACK)
  player2 := reversi.NewComputerPlayer(reversi.STONE_WHITE)

	// Define your custom evaluation functions which receives board, stone and return score
	eval := func(b *reversi.Board, stone int) int {
		mnum, enum := b.CountStone(stone), b.CountStone(-stone)
		return mnum - enum
	}

  var pos reversi.Pos

  for !board.CheckGameover() {
		pos = player1.Think(*board, eval, 6)
		board.Move(pos.Index, reversi.STONE_BLACK)
		fmt.Println(pos, board.CountStone(reversi.STONE_BLACK), board.CountStone(reversi.STONE_WHITE))
		pos = player2.Think(*board, eval, 6)
		board.Move(pos.Index, reversi.STONE_WHITE)
		fmt.Println(pos, board.CountStone(reversi.STONE_BLACK), board.CountStone(reversi.STONE_WHITE))
  }
}
```
