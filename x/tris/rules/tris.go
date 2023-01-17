package rules

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

const (
	BOARD_DIM = 9
	X         = "x"
	O         = "O"
	ROW_SEP   = "|"
)

type Player struct {
	Symbol string
}

type Piece struct {
	Player Player
}

var PieceStrings = map[Player]string{
	X_PLAYER:  "x",
	O_PLAYER:  "o",
	NO_PLAYER: "*",
}

var X_PLAYER = Player{X}
var O_PLAYER = Player{O}
var NO_PLAYER = Player{
	Symbol: "NO_PLAYER",
}

var Players = map[string]Player{
	X: X_PLAYER,
	O: O_PLAYER,
}

var Opponents = map[Player]Player{
	X_PLAYER: O_PLAYER,
	O_PLAYER: X_PLAYER,
}

var NO_PIECE = Piece{NO_PLAYER}

var StringPieces = map[string]Piece{
	"x": Piece{X_PLAYER},
	"o": Piece{O_PLAYER},
	"*": NO_PIECE,
}

type Pos struct {
	X int
}

var NO_POS = Pos{-1}
var Usable = map[Pos]bool{}

func init() {
	// Initialize usable spaces
	for y := 0; y < BOARD_DIM; y++ {
		Usable[Pos{X: y}] = true
	}
}

type Game struct {
	Pieces map[Pos]Piece
	Turn   Player
}

func New() *Game {
	pieces := make(map[Pos]Piece)
	game := &Game{pieces, X_PLAYER}
	return game
}

func (game *Game) TurnIs(player Player) bool {
	return game.Turn == player
}

func (game *Game) Winner() Player {
	//TODO
	if game.Pieces[Pos{0}] == game.Pieces[Pos{1}] && game.Pieces[Pos{2}] == game.Pieces[Pos{1}] {
		return game.Pieces[Pos{0}].Player
	} else {
		return NO_PLAYER
	}
}

func (game *Game) PieceAt(pos Pos) bool {
	_, ok := game.Pieces[pos]
	return ok
}

func (game *Game) ValidMove(pos Pos) bool {
	if game.PieceAt(pos) {
		return false
	}
	return true
}
func (game *Game) updateTurn() {
	opponent := Opponents[game.Turn]
	game.Turn = opponent

}
func (game *Game) Place(pos Pos, player Player) (err error) {
	if !game.ValidMove(pos) {
		return errors.New(fmt.Sprintf("This is not a valid place"))
	}

	if !game.TurnIs(player) {
		return errors.New(fmt.Sprintf("Not %v's turn", player))
	}

	game.Pieces[pos] = Piece{Player: player}
	game.updateTurn()

	return nil
}

func (game *Game) String() string {
	var buf bytes.Buffer
	for x := 0; x < BOARD_DIM; x++ {
		pos := Pos{x}
		if game.PieceAt(pos) {
			piece := game.Pieces[pos]
			val := PieceStrings[piece.Player]
			buf.WriteString(val)
		} else {
			buf.WriteString(PieceStrings[NO_PLAYER])
		}
		if x == 2 || x == 5 {
			buf.WriteString(ROW_SEP)
		}
	}
	return buf.String()
}
func ParsePiece(s string) (Piece, bool) {
	piece, ok := StringPieces[s]
	return piece, ok
}

func Parse(s string) (*Game, error) {
	if len(s) != 11 {
		return nil, errors.New(fmt.Sprintf("invalid board string: %v", s))
	}
	pieces := make(map[Pos]Piece)
	result := &Game{pieces, X_PLAYER}
	for y, row := range strings.Split(s, ROW_SEP) {
		for x, c := range strings.Split(row, "") {
			if x >= BOARD_DIM || y >= BOARD_DIM {
				return nil, errors.New(fmt.Sprintf("invalid board, piece out of bounds: %v, %v", x, y))
			}
			if piece, ok := ParsePiece(c); !ok {
				return nil, errors.New(fmt.Sprintf("invalid board, invalid piece at %v, %v", x, y))
			} else if piece != NO_PIECE {
				result.Pieces[Pos{x + 3*y}] = piece
			}
		}
	}
	return result, nil
}

func main() {
	game := New()
	fmt.Println(game.String())

	game, err := Parse("xxx|*oo|*o*")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(game.String())

}

func testPlace(game *Game, x int, p Player) {
	err := game.Place(Pos{x}, p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(game.String())
	fmt.Println("Winner: ", game.Winner())

}
