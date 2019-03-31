package game

import (
	"bytes"
	"fmt"
	"strconv"
)

type View interface {
	fmt.Stringer
}

type PlayerView struct {
	board Board
}

func NewPlayerView(board Board) PlayerView {
	return PlayerView{
		board: board,
	}
}

func (v *PlayerView) String() string {
	var buf bytes.Buffer

	// padding for number row
	buf.WriteString("   ")
	// column numbers
	// TODO make padding work right with numbers > 9
	for i := 0; i < v.board.size; i++ {
		buf.WriteString(" " + strconv.Itoa(i+1) + " ")
	}
	buf.WriteString("\n")

	// number row divider
	buf.WriteString(" +-")
	for i := 0; i < v.board.size; i++ {
		buf.WriteString("---")
	}
	buf.WriteString("\n")

	for i, row := range v.board.Points {
		// alpha column
		buf.WriteString(string('A'+byte(i)) + "| ")
		for _, p := range row {
			buf.WriteString(" " + p.String() + " ")
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

type OpponentView struct {
	board Board
}

func NewOpponentView(board Board) OpponentView {
	return OpponentView{
		board: board,
	}
}

func (v *OpponentView) String() string {
	var buf bytes.Buffer

	// padding for number row
	buf.WriteString("   ")
	// column numbers
	// TODO make padding work right with numbers > 9
	for i := 0; i < v.board.size; i++ {
		buf.WriteString(" " + strconv.Itoa(i+1) + " ")
	}
	buf.WriteString("\n")

	// number row divider
	buf.WriteString(" +-")
	for i := 0; i < v.board.size; i++ {
		buf.WriteString("---")
	}
	buf.WriteString("\n")

	for i, row := range v.board.Points {
		// alpha column
		buf.WriteString(string('A'+byte(i)) + "| ")
		for _, p := range row {
			buf.WriteString(" " + string(v.pointToRune(*p)) + " ")
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func (v *OpponentView) pointToRune(p Point) rune {
	switch p.State {
	case StateHit:
		return 'X'
	case StateMiss:
		return 'O'
	}

	return '.'
}
