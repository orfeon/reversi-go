package reversi

const (
	DEPTH_TH = 6
)

type Player interface {
	Think(b Board) Pos
}

type ComputerPlayer struct {
	stone int
}

func NewComputerPlayer(stone int) *ComputerPlayer {
	if stone != STONE_BLACK && stone != STONE_WHITE {
		panic("Stone must be STONE_BLACK or STONE_WHITE!")
	}
	p := new(ComputerPlayer)
	p.stone = stone
	return p
}

func (p *ComputerPlayer) Think(b Board) Pos {

	movables := b.CalcMovable(p.stone)

	if len(movables) == 0 {
		return Pos{Index: INDEX_SKIP, Stone: p.stone}
	} else if len(movables) == 1 {
		return Pos{Index: movables[0], Stone: p.stone}
	}

	score, index := p.alphabeta(&b, p.stone, -10000, 10000, DEPTH_TH)
	return Pos{Index: index, Stone: p.stone, Score: score}
}

func (p *ComputerPlayer) alphabeta(b *Board, stone, alpha, beta, depth int) (int, int) {

	if depth <= 0 || b.CheckGameover() {
		return p.evaluate(b, stone), INDEX_SKIP
	}

	movables := b.CalcMovable(stone)
	if len(movables) == 0 {
		b.Skip(stone)
		score, _ := p.alphabeta(b, -stone, -beta, -alpha, depth-1)
		b.Undo()
		score = -score
		return score, INDEX_SKIP
	}

	bestindex := INDEX_SKIP
	for _, index := range movables {
		b.Move(index, stone)
		score, _ := p.alphabeta(b, -stone, -beta, -alpha, depth-1)
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

func (p *ComputerPlayer) evaluate(b *Board, stone int) int {

	rest_turn := b.CountStone(STONE_BLANK)

	mnum, enum := b.CountStone(stone), b.CountStone(-stone)
	if rest_turn < DEPTH_TH {
		return mnum - enum
	}
	mmovables := b.CalcMovable(p.stone)
	emovables := b.CalcMovable(-p.stone)

	return (mnum - enum) + 6*(len(mmovables)-len(emovables))
}
