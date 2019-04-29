package main

import (
	"Marsrover/car"
	"Marsrover/coordinate"
	"Marsrover/library"
	"bufio"
	"fmt"
	"os"
)

type FUNC func(string)(list []string, err error)

func main(){

	fmt.Println("Game start...")

	start()

}

func start(){

	r := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command-> ")

	//First command
	command1 := HandleCommandFactory(r, "first")
	marsMap := BuildAreasWithString(command1)

	for {
		//第二条命令处理
		command2 := HandleCommandFactory(r, "second")
		marsCar := CreateNewCarWithString(command2)

		//第三条命令处理
		command3 := HandleCommandFactory(r, "third")
		err := marsCar.ReceiveAndExecute(marsMap, command3)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}




}


func HandleCommandFactory(r *bufio.Reader, index string)[]string{
	for{
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		var f FUNC
		switch index {
		case "first":
			f = library.ValidateFristLine
		case "second":
			f = library.ValidateCommand1
		case "third":
			f= library.ValidateCommand2
		default:
			f = library.ValidateCommand2
		}

		list, err := f(line)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return list
	}
}

func BuildAreasWithString(XYList []string) *coordinate.MarsMap{
	var marsMap = new(coordinate.MarsMap)
	marsMap.MapPos.SetX(XYList[0])
	marsMap.MapPos.SetX(XYList[1])
	marsMap.RipPostionList = []coordinate.RipPosition{}

	return marsMap
}

func CreateNewCarWithString(commandList []string) *car.Car {
	fmt.Println("Create a new car...")
	c := ListToCoordinater(commandList)

	car := &car.Car{c,commandList[2]}

	return car
}

func ListToCoordinater(list []string) *coordinate.Coordinate{
	c := new(coordinate.Coordinate)
	c.SetX(list[0])
	c.SetY(list[1])

	return c
}

