package coordinate

type MarsMap struct {
	MapPos Coordinate
	RipPostionList []RipPosition
}

func (m *MarsMap)IsOutOfBoundary(c *Coordinate)bool{
	flag := false
	if c.X<0 || c.Y<0 || c.X>m.MapPos.X || c.Y > m.MapPos.Y {
		flag = true
	}

	return flag
}

func (m *MarsMap)AddRipList(rip *RipPosition){
	if rip!= nil{
		m.RipPostionList = append(m.RipPostionList, *rip)
	}
}
