package reversi

import (
	"fmt"
	"math"
)

const (
	STONE_BLACK = -1
	STONE_WHITE = 1
	STONE_BLANK = 0

	MAX_STONE_NUM = 64
	INDEX_SKIP    = -1
)

var (
	directions    = [8]int{-9, -8, -7, -1, 1, 7, 8, 9} // leftup, up, rightup, left, right, leftdown, down, rightdown
	directionsxth = [8]int{7, -1, 0, 7, 0, 7, -1, 0}
	line1         = [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	line2         = [8]int{56, 57, 58, 59, 60, 61, 62, 63}
	line3         = [8]int{0, 8, 16, 24, 32, 40, 48, 56}
	line4         = [8]int{7, 15, 23, 31, 39, 47, 55, 63}
)

// Pos Stone position struct
type Pos struct {
	Index       int
	Stone       int
	Score       int
	Acquirables []int
}

func (p *Pos) String() string {
	return fmt.Sprintf("%d", p.Index)
}

// History Stone position history struct
type History struct {
	poslist   [120]Pos
	lastindex int
}

// NewHistory Constructor of History
func NewHistory() *History {
	h := new(History)
	return h
}

// Push Push history struct to stack
func (h *History) Push(pos Pos) {
	h.poslist[h.lastindex] = pos
	h.lastindex++
}

// Pop Pos history struct from stack
func (h *History) Pop() Pos {
	if h.lastindex == 0 {
		return Pos{Index: -1, Stone: 0}
	}
	h.lastindex--
	pos := h.poslist[h.lastindex]
	return pos
}

func (h *History) Last() Pos {
	if h.lastindex == 0 {
		return Pos{Index: -1, Stone: 0}
	}
	pos := h.poslist[h.lastindex]
	return pos
}

func (h *History) Size() int {
	return h.lastindex
}

func (h *History) Clear() {
	h.lastindex = 0
}

func (h *History) CheckLastSkipNum() int {
	skipnum := 0
	if h.lastindex > 0 && h.poslist[h.lastindex-1].Index == INDEX_SKIP {
		skipnum += 1
	}
	if h.lastindex > 1 && h.poslist[h.lastindex-2].Index == INDEX_SKIP {
		skipnum += 1
	}
	return skipnum
}

type Board struct {
	Stones  [MAX_STONE_NUM]int
	history History
	Liverty [MAX_STONE_NUM]int
}

func NewBoard() *Board {
	b := new(Board)
	b.Stones[27], b.Stones[28] = STONE_WHITE, STONE_BLACK
	b.Stones[35], b.Stones[36] = STONE_BLACK, STONE_WHITE
	b.Liverty = b.Stones
	return b
}

func (b *Board) Move(index, stone int) int {

	if index == INDEX_SKIP {
		pos := Pos{Index: index, Stone: stone}
		b.history.Push(pos)
		return 0
	}
	acquirables := b.calcAcquirables(index, stone)
	if len(acquirables) == 0 {
		return 0
	}
	for _, acquirable := range acquirables {
		b.Stones[acquirable] = stone
	}
	b.Stones[index] = stone
	for _, dir := range directions {
		if index+dir < 0 || index+dir > 63 {
			continue
		}
		b.Liverty[index+dir]--
	}
	pos := Pos{Index: index, Stone: stone, Acquirables: acquirables}
	b.history.Push(pos)
	return len(acquirables) + 1
}

func (b *Board) Undo() {

	if b.history.Size() == 0 {
		return
	}
	pos := b.history.Pop()
	if pos.Index == INDEX_SKIP {
		return
	}
	stone := b.Stones[pos.Index]
	for _, index := range pos.Acquirables {
		b.Stones[index] = -stone
	}
	b.Stones[pos.Index] = STONE_BLANK
	for _, dir := range directions {
		if pos.Index+dir < 0 || pos.Index+dir > 63 {
			continue
		}
		b.Liverty[pos.Index+dir]++
	}

}

func (b *Board) Skip(stone int) int {
	pos := Pos{Index: INDEX_SKIP, Stone: stone}
	b.history.Push(pos)
	return 0
}

func (b *Board) Turn() int {
	pos := b.history.Last()
	return -pos.Stone
}

func (b *Board) CalcMovable(stone int) []int {

	movables := make([]int, 0, 60)
	for index := 0; index < MAX_STONE_NUM; index++ {
		acquirables := b.calcAcquirables(index, stone)
		if len(acquirables) > 0 {
			movables = append(movables, index)
		}
	}
	return movables
}

func (b *Board) CountStone(stoneToCount int) int {

	stoneCount := 0
	for _, stone := range b.Stones {
		if stone == stoneToCount {
			stoneCount++
		}
	}
	return stoneCount
}

func (b *Board) CheckGameover() bool {

	if b.CountStone(STONE_BLANK) == 0 {
		return true
	}
	if b.history.CheckLastSkipNum() == 2 {
		return true
	}
	return false
}

func (b *Board) Clear() {
	for index := 0; index < MAX_STONE_NUM; index++ {
		b.Stones[index] = 0
	}
	b.Stones[27], b.Stones[28] = STONE_WHITE, STONE_BLACK
	b.Stones[35], b.Stones[36] = STONE_BLACK, STONE_WHITE
	b.history.Clear()
}

func (b *Board) calcAcquirables(index, stone int) []int {

	acquirables := make([]int, 0, 18)

	if index < 0 || b.Stones[index] != 0 {
		return acquirables
	}

	for _, dir := range directions {

		dirAcquirables := make([]int, 0, 6)

		preX := index % 8
		for i := index + dir; i < MAX_STONE_NUM && i >= 0; i += dir {
			curX := i % 8
			if math.Abs(float64(curX)-float64(preX)) > 1 {
				break
			}
			preX = curX

			if b.Stones[i] == -stone {
				dirAcquirables = append(dirAcquirables, i)
			} else {
				if b.Stones[i] == stone && len(dirAcquirables) > 0 {
					acquirables = append(acquirables, dirAcquirables...)
				}
				break
			}
		}
	}

	return acquirables
}

func (b *Board) CountMobility(stone int) int {
	return len(b.CalcMovable(stone))
}

func (b *Board) CountLiverty(stone int) int {
	score := 0
	for index := 0; index < MAX_STONE_NUM; index++ {
		if b.Liverty[index] == stone {
			score += b.Liverty[index]
		}
	}
	return score
}

func (b *Board) CountStable(stone int) int {
	score := 0
	score += b.countStable(line1, stone)
	score += b.countStable(line2, stone)
	score += b.countStable(line3, stone)
	score += b.countStable(line4, stone)
	return score
}

func (b *Board) CountWindow(stone int) int {
	score := 0
	score += b.countWindow(line1, stone)
	score += b.countWindow(line2, stone)
	score += b.countWindow(line3, stone)
	score += b.countWindow(line4, stone)
	return score
}

func (b *Board) CountCornerStone(stone int) int {
	score := 0
	for i := range [4]int{0, 7, 56, 63} {
		if b.Stones[i] == stone {
			score++
		}
	}
	return score
}

func (b *Board) countStable(line [8]int, stone int) int {
	score := 0
	for i := 0; i < 8; i++ {
		if b.Stones[line[i]] != stone {
			break
		}
		score++
	}
	if score < 8 {
		for i := 7; i >= 0; i-- {
			if b.Stones[line[i]] != stone {
				break
			}
			score++
		}
	}
	return score
}

func (b *Board) countWindow(line [8]int, stone int) int {
	if b.Stones[line[0]] == stone || b.Stones[line[7]] == stone {
		return 0
	}
	for i := 2; i < 6; i++ {
		if b.Stones[line[i]] != stone {
			return 0
		}
	}
	if (b.Stones[line[1]] == stone && b.Stones[line[7]] == stone) || (b.Stones[line[1]] != stone && b.Stones[line[7]] != stone) {
		return 0
	}
	return 1
}
