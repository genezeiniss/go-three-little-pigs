package main

import "fmt"

type bigBadWolf struct{}

func (wolf bigBadWolf) requestToEnter() {
	fmt.Println("Big Bad Wolf: \"Let me in, let me in!\"")
}

func (wolf bigBadWolf) huffAndPuff(m map[littlePig]house, pig littlePig) {

	mapSize := len(m)
	houseMaterial := m[pig].material

	fmt.Println("Wolf huffs and puffs")

	if BRICKS != houseMaterial {
		delete(m, pig)
	}

	if mapSize > len(m) {
		fmt.Printf("Wolf blew %v house\n", houseMaterial)
	} else {
		fmt.Printf("Wolf could not blow down %v house\n", houseMaterial)
	}
}
