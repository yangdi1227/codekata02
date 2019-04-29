package coordinate

import "strconv"

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) Add(step *Coordinate)  {
	if step == nil {
		return
	}
	c.SetX(c.X + step.X)
	c.SetY(c.Y + step.Y)

	return
}


func (c *Coordinate)Equal(Pos *Coordinate) bool {
	flag := false
	if c.X == Pos.X && c.Y == Pos.Y{
		flag = true
	}

	return flag
}

func (c *Coordinate)SetX(i interface{}){
	switch v := i.(type){
	case string:
		c.X, _ = strconv.Atoi(v)
	case int:
		c.X = v
	}
}

func (c *Coordinate)SetY(i interface{}){
	switch v := i.(type){
	case string:
		c.Y, _ = strconv.Atoi(v)
	case int:
		c.Y = v
	}
}