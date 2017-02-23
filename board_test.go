package reversi

import (
	_ "fmt"
	"testing"
)

//  Board(8 x 8) index.
//
//   0, 1, 2, 3, 4, 5, 6, 7,
//   8, 9,10,11,12,13,14,15,
//  16,17,18,19,20,21,22,23,
//  24,25,26,27,28,29,30,31,
//  32,33,34,35,36,37,38,39,
//  40,41,42,43,44,45,46,47,
//  48,49,50,51,52,53,54,55,
//  56,57,58,59,60,61,62,63

func TestMove(t *testing.T) {

	board := NewBoard()
	// Test init board. All OK pattern
	for _, index := range []int{19, 26, 37, 44} {
		board.Clear()
		acount := board.Move(index, STONE_BLACK)
		if acount == 0 {
			t.Errorf("%s must be set zero but %s", index, acount)
		}
	}

	for _, index := range []int{20, 29, 34, 43} {
		board.Clear()
		acount := board.Move(index, STONE_WHITE)
		if acount == 0 {
			t.Errorf("%s must be set zero but %s", index, acount)
		}
	}

	// Test init board. NG pattern.
	for _, index := range []int{19, 26, 37, 44, 0, 7, 56, 63} {
		board.Clear()
		acount := board.Move(index, STONE_WHITE)
		if acount != 0 {
			t.Errorf("%s must be set zero but %s", index, acount)
		}
	}

	for _, index := range []int{20, 29, 34, 43, 0, 7, 56, 63} {
		board.Clear()
		acount := board.Move(index, STONE_BLACK)
		if acount != 0 {
			t.Errorf("%s must be set zero but %s", index, acount)
		}
	}

}
