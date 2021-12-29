package main

import "fmt"

type BigBadWolf struct{}

func (wolf BigBadWolf) RequestToEnter() {
	fmt.Println("Big Bad Wolf requested \"Let me in, let me in!\"")
}

func (wolf BigBadWolf) HuffAndPuff(m map[LittlePig]House, pig LittlePig) {

	mapSize := len(m)
	houseMaterial := m[pig].Material

	fmt.Println("Wolf huffed and puffed")

	if BRICKS != houseMaterial {
		delete(m, pig)
	}

	if mapSize > len(m) {
		fmt.Printf("And blew the %s house and ate the %s.\n", houseMaterial, GetPigName(pig))
	} else {
		fmt.Printf("Wolf is too weak to blow down the %s house. Wolf in risk to be eaten himself.\n", houseMaterial)
	}
}
