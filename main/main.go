package main

import "fmt"

func main()  {

	prologue := "Once upon a time there were three little pigs and the time came for them to leave home and seek their fortunes."
	fmt.Println(prologue)

	littlePigHouseMap := make(map[littlePig]house)
	littlePigs := []littlePig{firstLittlePig{}, secondLittlePig{}, thirdLittlePig{}}

	for _, littlePig := range littlePigs {
		littlePigHouseMap[littlePig] = buildHouse(littlePig)
	}

	bigBadWolf := bigBadWolf{}

	for  pig := range littlePigHouseMap {

		bigBadWolf.requestToEnter()
		rejectRequest(pig)
		bigBadWolf.huffAndPuff(littlePigHouseMap, pig)
	}
}
