package boggle

type Game interface {
	// Return the size of the game board.
	Size() (height int, width int)

	// Return the visible letters on the game board.
	Visible() [][]byte

	// Mix up the visible letters.
	Shake()
}

// Contains all data for a single game.
type Board struct {
	Height, Width  int
	Dice           DieSet
	VisibleLetters [][]byte
}

// 2008 version of Boggle.
var Boggle = Board{
	Height: 4,
	Width:  4,
	Dice:   NewBoggleDice,
}

// 1976 version of Boggle.
var OldBoggle = Board{
	Height: 4,
	Width:  4,
	Dice:   OldBoggleDice,
}

// Shake the dice on the board.
// TODO(inhies): Actually do stuff...
func (b *Board) Shake() {}

// Return the size of the board.
func (b *Board) Size() (height int, width int) {
	return b.Height, b.Width
}

// Return the visible letters on the board.
func (b *Board) Visible() [][]byte {
	return [][]byte{}
}
