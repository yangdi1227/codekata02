package coordinate

type RipPosition struct {
	Pos        *Coordinate
	LastAction string
}

func (rip *RipPosition)Equals(ripPosition *RipPosition)bool {
	flag := false
	if rip.Pos.X == ripPosition.Pos.X && rip.Pos.Y == ripPosition.Pos.Y {
		flag = true
	}

	return flag
}
