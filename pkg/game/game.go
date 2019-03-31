package game

type Game struct {
	Board1 Board
	Board2 Board
	Ships  []Ship
}

func NewGame(size int) Game {
	return Game{
		Board1: NewBoard(size),
		Board2: NewBoard(size),
		Ships: []Ship{
			Ship{Size: 5, Name: "Carrier"},
			Ship{Size: 4, Name: "Battleship"},
			Ship{Size: 3, Name: "Cruiser"},
			Ship{Size: 3, Name: "Destroyer"},
			Ship{Size: 2, Name: "PT Boat"},
		},
	}
}
