package game

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	OrientationHorizontal = 0
	OrientationVertical   = 1
)

type Board struct {
	Points [][]*Point
	// since the board is square, we keep around the size to make calcs easier
	size int
}

func NewBoard(size int) Board {
	points := make([][]*Point, size)
	for y := 0; y < size; y++ {
		points[y] = make([]*Point, size)
		for x := 0; x < size; x++ {
			points[y][x] = &Point{}
		}
	}

	return Board{
		Points: points,
		size:   size,
	}
}

func (b *Board) setShip(ship *Ship, x int, y int) {
	b.Points[y][x].Ship = ship
}

func (b *Board) pointAt(x int, y int) *Point {
	return b.Points[y][x]
}

// place a ship on the board
// returns error if cannot be placed there
// x and y are the top-most or left-most point of the ship
func (b *Board) PlaceShip(ship *Ship, x int, y int) error {
	// check x and y are within range
	if x < 0 || x > b.size-1 {
		return fmt.Errorf("x: %d is beyond the edge of the board.", x)
	}

	if y < 0 || y > b.size-1 {
		return fmt.Errorf("x: %d is beyond the edge of the board.", x)
	}

	if ship.Orientation == OrientationHorizontal {
		end := x + ship.Size
		// will the ship fit?
		if end > b.size {
			return fmt.Errorf("Cannot put %s (%d long) at (%d, %d) because it would overflow the board", ship.Name, ship.Size, x, y)
		}

		// is it occupied?
		for i := x; i < end; i++ {
			if b.pointAt(i, y).Ship != nil {
				return fmt.Errorf("One of the points already has a ship there!")
			}
		}

		// place it
		for i := x; i < end; i++ {
			b.setShip(ship, i, y)
		}
	} else {
		// Vertical
		// will the ship fit?
		end := y + ship.Size
		if end > b.size {
			return fmt.Errorf("Cannot put %s (%d long) at (%d, %d) because it would overflow the board", ship.Name, ship.Size, x, y)
		}

		// is it occupied?
		for i := y; i < end; i++ {
			if b.pointAt(x, i).Ship != nil {
				return fmt.Errorf("One of the points already has a ship there!")
			}
		}

		// place it
		for i := y; i < end; i++ {
			b.setShip(ship, x, i)
		}
	}

	return nil
}

func (b *Board) Generate(ships []Ship) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// place ships randomly
	for i := range ships {
		// loop until there's no more error
		for {
			x := r.Intn(b.size)
			y := r.Intn(b.size)
			ships[i].Orientation = r.Intn(OrientationVertical + 1)

			if err := b.PlaceShip(&ships[i], x, y); err == nil {
				break
			}
		}
	}
}

func (b *Board) Fire(x int, y int) (bool, error) {
	p := b.pointAt(x, y)
	if p.State != StateEmpty {
		return false, fmt.Errorf("Already fired at (%d, %d)", x, y)
	}

	if p.Ship == nil {
		p.State = StateMiss
		return false, nil
	}

	p.State = StateHit
	p.Ship.Hits += 1
	if p.Ship.IsSunk() {
		// TODO do something when sunk
	}
	return true, nil
}
