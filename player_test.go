package reversi

import (
	"fmt"
	"testing"
)

func TestComputerPlayer(t *testing.T) {

	board := NewBoard()

	player1 := NewComputerPlayer(STONE_BLACK)
	player2 := NewComputerPlayer(STONE_WHITE)

	a := func(b *Board, stone int) int {
		restTurn := b.CountStone(STONE_BLANK)

		mnum, enum := b.CountStone(stone), b.CountStone(-stone)
		if restTurn < 5 {
			return mnum - enum
		}
		mmovables := b.CalcMovable(stone)
		emovables := b.CalcMovable(-stone)

		return (mnum - enum) + 6*(len(mmovables)-len(emovables))
	}

	var pos Pos

	for !board.CheckGameover() {
		pos = player1.Think(*board, a, 6)
		board.Move(pos.Index, STONE_BLACK)
		fmt.Println(pos, board.CountStone(STONE_BLACK), board.CountStone(STONE_WHITE))
		pos = player2.Think(*board, a, 6)
		board.Move(pos.Index, STONE_WHITE)
		fmt.Println(pos, board.CountStone(STONE_BLACK), board.CountStone(STONE_WHITE))
	}

	if pos.Index < 0 {
		t.Errorf("Failed!")
	}

}
