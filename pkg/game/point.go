package game

const (
	StateEmpty = iota
	StateHit
	StateMiss
)

type Point struct {
	State int
	// A reference to the ship that is on this point
	Ship *Ship
}

func (p *Point) String() string {
	switch p.State {
	case StateHit:
		return "x"
	case StateMiss:
		return "o"
	}

	if p.Ship == nil {
		return "."
	} else {
		return string(p.Ship.Name[0])
	}
}
