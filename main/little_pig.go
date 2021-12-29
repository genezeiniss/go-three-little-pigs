package main

import (
	"fmt"
	"go-three-little-pigs/main/utils"
	"reflect"
	"time"
)

type LittlePig interface {
	GetMaterial() string
}

func BuildHouse(pig LittlePig, channel chan ChannelResult) {

	time.Sleep(5 * time.Second)
	printOutPigPerformance(pig)
	channel <- ChannelResult{Pig: pig, House: NewHouse(pig.GetMaterial())}
}

func printOutPigPerformance(pig LittlePig) {

	var suffix string
	material := pig.GetMaterial()

	switch material {
	case STRAW:
		suffix = ", because it was the easiest thing to do."
	case STICKS:
		suffix = ". This was a little bit stronger than a straw House."
	case BRICKS:
		suffix = "."
	}

	fmt.Printf("The %v built his House out of %v%v\n", GetPigName(pig), material, suffix)
}

func GetPigName(pig LittlePig) string {
	return utils.AmendString(reflect.TypeOf(pig).Name())
}

func RejectRequest(pig LittlePig) {
	reject := "Not by the hair of my chinny chin chin"
	fmt.Printf("%v: \"%v\"\n", GetPigName(pig), reject)
}
