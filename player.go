package reversi

// Player TODO
type Player interface {
	Think(b Board) Pos
}

// ComputerPlayer TODO
type ComputerPlayer struct {
	stone int
}

// Evaluate Abstract Method
type Evaluate func(b *Board, stone int) int

// NewComputerPlayer Constructor of ComputerPlayer struct
func NewComputerPlayer(stone int) *ComputerPlayer {
	if stone != STONE_BLACK && stone != STONE_WHITE {
		panic("Stone must be STONE_BLACK or STONE_WHITE!")
	}
	p := new(ComputerPlayer)
	p.stone = stone
	return p
}

func (p *ComputerPlayer) Think(b Board, e Evaluate, depth int) Pos {

	movables := b.CalcMovable(p.stone)

	if len(movables) == 0 {
		return Pos{Index: INDEX_SKIP, Stone: p.stone}
	} else if len(movables) == 1 {
		return Pos{Index: movables[0], Stone: p.stone}
	}

	score, index := p.alphabeta(&b, e, p.stone, -10000, 10000, depth)
	return Pos{Index: index, Stone: p.stone, Score: score}
}

func (p *ComputerPlayer) alphabeta(b *Board, eval Evaluate, stone, alpha, beta, depth int) (int, int) {

	if depth <= 0 || b.CheckGameover() {
		//return p.evaluate(b, stone), INDEX_SKIP
		return eval(b, stone), INDEX_SKIP
	}

	movables := b.CalcMovable(stone)
	if len(movables) == 0 {
		b.Skip(stone)
		score, _ := p.alphabeta(b, eval, -stone, -beta, -alpha, depth-1)
		b.Undo()
		score = -score
		return score, INDEX_SKIP
	}

	bestindex := INDEX_SKIP
	for _, index := range movables {
		b.Move(index, stone)
		score, _ := p.alphabeta(b, eval, -stone, -beta, -alpha, depth-1)
		b.Undo()
		score = -score

		if score > alpha {
			alpha = score
			bestindex = index
		}
		if alpha >= beta {
			break
		}
	}
	return alpha, bestindex
}
