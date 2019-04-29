package library

import (
	"Marsrover/coordinate"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var Vertor = [4]coordinate.Coordinate{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

var DirectionIndex = map[string]int{
	"N": 0,
	"E": 1,
	"S": 2,
	"W": 3,
}

func TurnTo(current, oriented string) (new_oriented string, err error) {
	var index int
	value, ok := DirectionIndex[current]
	if ok {
		if oriented == "L" {
			index = (value - 1 + 4) % 4
		} else if oriented == "R" {
			index = (value + 1) % 4
		}

		for key, value := range DirectionIndex {
			if value == index {
				return key, nil
			}
		}
	}

	return "", errors.New("Wrong direction!")


}

func Validater(i int, line string) (splitItems []string, err error) {

	var splitLine []string
	// first line
	if i == 1 {
		splitLine, err = ValidateFristLine(line)
	}
	if i%2 == 0 {
		splitLine, err = ValidateCommand1(line)
	}
	if i>1 && i%2 == 1 {
		splitLine, err = ValidateCommand2(line)
	}

	return splitLine, err
}

func ValidateFristLine(line string)(list []string, err error){

	splitLine, err := validateNilAndLengthBySpace(line, 2)

	for _, ch := range splitLine {
		if !(ch >= "0" && ch <="9") {
			err = errors.New(fmt.Sprint("Must be numeric:%s. Wrong format! Shoule be like :5 5. Enter command again-> ", line) )
			break
		}
	}

	return splitLine, err

}

func ValidateCommand1(line string)(list []string, err error){

	splitLine, err := validateNilAndLengthBySpace(line, 3)

	if err!=nil {
		return nil, err
	}
	if !(isDIgitAndMatchLen(splitLine[0], 1) && isDIgitAndMatchLen(splitLine[1], 1) && isCharacterAndMatchLen(splitLine[0], 1)){
		return nil ,errors.New("Wrong format! Shoule be like :1 1 N. Enter command again-> ")
	}

	return splitLine, nil

}

func ValidateCommand2(line string)(list []string, err error){
	err = nil
	trimLine := strings.TrimSpace(line)
	splitLine := strings.Fields(trimLine)
	LRMList := []string{"L", "R", "M"}
	for _,item := range splitLine {
		if isIn, _ := in_array(item, LRMList);isIn == false{
			return nil, errors.New("Wrong command!should be L M R!")
		}
	}

	return splitLine, nil
}

func validateNilAndLengthBySpace(line string, length int) (list []string, err error) {
	err = nil
	trimLine := strings.TrimSpace(line)
	splitLine := strings.Fields(trimLine)
	if trimLine == "" || len(splitLine) == length {
		err = errors.New("line is nil or wrong length!")
	}

	return splitLine, err
}

func isDIgitAndMatchLen(ch string, i int)bool{
	flag := false
	if _, err:=validateNilAndLengthBySpace(ch, i);err==nil {
		if ch >= "0" && ch <="9" {
			flag = true
		}
	}


	return flag
}

func isCharacterAndMatchLen(ch string, i int)bool{
	flag := false
	if _, err:=validateNilAndLengthBySpace(ch, i);err==nil {
		if (ch >="a" && ch<="z" || ch>="A" && ch<="Z") {
			flag = true
		}
	}

	return flag
}

func in_array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}