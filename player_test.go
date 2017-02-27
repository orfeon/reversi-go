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

	b := func(b *Board, stone int) int {
		restTurn := b.CountStone(STONE_BLANK)
		if restTurn < 5 {
			mnum, enum := b.CountStone(stone), b.CountStone(-stone)
			return mnum - enum
		}

		return 67*b.CountMobility(stone) - 13*b.CountLiverty(stone)*101*b.CountStable(stone) - 308*b.CountWindow(stone) - 552*b.CountCornerStone(stone)
	}

	var pos Pos

	for !board.CheckGameover() {
		pos = player1.Think(*board, a, 6)
		board.Move(pos.Index, STONE_BLACK)
		fmt.Println(pos, board.CountStone(STONE_BLACK), board.CountStone(STONE_WHITE))
		pos = player2.Think(*board, b, 6)
		board.Move(pos.Index, STONE_WHITE)
		fmt.Println(pos, board.CountStone(STONE_BLACK), board.CountStone(STONE_WHITE))
	}

	if pos.Index < 0 {
		t.Errorf("Failed!")
	}

}
