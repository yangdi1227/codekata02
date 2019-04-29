package car

import (
	"Marsrover/coordinate"
	"Marsrover/exception"
	"Marsrover/library"
	"fmt"
)

type Car struct {
	Coor      *coordinate.Coordinate
	Direction string
}

func (c *Car)GetDirection() string{
	return c.Direction
}


func (c *Car) Turn(oriented string) {
	newOriented, err := library.TurnTo(c.Direction, oriented)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Direction = newOriented

	fmt.Println("Turn from " + oriented + "to : " + c.Direction)
}


func (c *Car) Move(marsmap *coordinate.MarsMap) (err error){
	err = nil
	if isSkipRip := c.isRipCommand(marsmap);isSkipRip{
		err = exception.SkipCmdByRipInfoError
		return err
	}
	c.executeOneStep()
	fmt.Println("moved to", c.Coor.X, c.Coor.Y)

	if isRip := marsmap.IsOutOfBoundary(c.Coor);isRip{
		err = exception.RIPPositionError
		return err
	}

	return err
}

func (c *Car)ReceiveAndExecute(marsMap *coordinate.MarsMap, cmdList []string)(err error) {
	err = nil
	fmt.Println("Start to execute :", cmdList)
	for _, cmd := range cmdList {
		if cmd == "L" || cmd == "R"{
			c.Turn(cmd)
		}
		if cmd == "M"{
			oriX, oriY, oriDirection := c.Coor.X,c.Coor.Y, c.Direction
			preRipPosition := &coordinate.RipPosition{c.Coor, c.Direction}
			err = c.Move(marsMap)

			if err == exception.SkipCmdByRipInfoError {
				fmt.Println(err)
				continue
			}else if err == exception.RIPPositionError {
				//add new ripposition
				marsMap.AddRipList(preRipPosition)

				result := fmt.Sprintf("%s %s %s RIP", oriX, oriY, oriDirection)
				fmt.Println(result)
				return err
			}
		}
	}
	result := fmt.Sprintf("%s %s %s", c.Coor.X, c.Coor.Y, c.Direction)
	fmt.Println(result)
	return err

}

func (c *Car)executeOneStep(){
	currentDirection := c.Direction
	oneStep := library.Vertor[library.DirectionIndex[currentDirection]]
	c.Coor.Add(&oneStep)
}

func (c *Car)isRipCommand(marsMap *coordinate.MarsMap)bool{
	flag := false
	carInfo := c.getCarInfo()
	for _,rip := range marsMap.RipPostionList{
		if rip.Equals(carInfo) {
			flag = true
			break
		}
	}
	return flag
}

func (c *Car) getCarInfo()(rip *coordinate.RipPosition){
	return &coordinate.RipPosition{c.Coor, c.Direction}


}