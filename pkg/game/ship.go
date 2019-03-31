package game

type Ship struct {
	Size        int
	Name        string
	Orientation int
	Hits        int
}

func (s *Ship) IsSunk() bool {
	return s.Hits == s.Size
}
