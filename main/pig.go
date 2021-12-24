package main

import (
	"fmt"
	"go-three-little-pigs/main/utils"
	"reflect"
)

type littlePig interface {
	getMaterial() string
}

func buildHouse(pig littlePig) house {
	material := pig.getMaterial()
	fmt.Printf("%v builds his house from %v\n", utils.AmendString(reflect.TypeOf(pig).Name()), material)
	//time.Sleep(5 * time.Second)
	return newHouse(material)
}

func rejectRequest(pig littlePig) {
	reject := "Not by the hair of my chinny chin chin"
	fmt.Printf("%v: \"%v\"\n", utils.AmendString(reflect.TypeOf(pig).Name()), reject)
}
