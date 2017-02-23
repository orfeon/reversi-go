package reversi

import (
	"fmt"
	"testing"
)

func aTestComputerPlayer(t *testing.T) {

	board := NewBoard()

	player1 := NewComputerPlayer(STONE_BLACK)
	player2 := NewComputerPlayer(STONE_WHITE)

	var pos Pos

	for !board.CheckGameover() {
		pos = player1.Think(*board)
		board.Move(pos.Index, STONE_BLACK)
		fmt.Println(pos, board.CountStone(STONE_BLACK), board.CountStone(STONE_WHITE))
		pos = player2.Think(*board)
		board.Move(pos.Index, STONE_WHITE)
		fmt.Println(pos, board.CountStone(STONE_BLACK), board.CountStone(STONE_WHITE))
	}

	if pos.Index < 0 {
		t.Errorf("Failed!")
	}

}
