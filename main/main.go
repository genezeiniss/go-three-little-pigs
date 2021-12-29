package main

import (
	"fmt"
)

type ChannelResult struct {
	Pig   LittlePig
	House House
}

func main() {

	fmt.Println("Once upon a time there were three little pigs and the time came for them to leave home and seek their fortunes.")

	littlePigs := []LittlePig{FirstLittlePig{}, SecondLittlePig{}, ThirdLittlePig{}}

	channel := make(chan ChannelResult)

	for _, pig := range littlePigs {
		go BuildHouse(pig, channel)
	}

	pigHouseMap := make(map[LittlePig]House)
	for i := 0; i < len(littlePigs); i++ {
		channelResult := <-channel
		pigHouseMap[channelResult.Pig] = channelResult.House
	}

	fmt.Println("   ***   ")
	fmt.Println("One night the Big Bad Wolf, who dearly loved to eat fat little piggies, came along.")

	BigBadWolf := BigBadWolf{}

	for pig, house := range pigHouseMap {

		fmt.Printf("Big Bad Wolf saw the little pig in his house of %s.\n", house.Material)

		BigBadWolf.RequestToEnter()
		RejectRequest(pig)
		BigBadWolf.HuffAndPuff(pigHouseMap, pig)
		fmt.Println("   ***   ")
	}
}
