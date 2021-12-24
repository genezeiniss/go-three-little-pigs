package main

const (
	STRAW = "straw"
	STICKS = "sticks"
	BRICKS = "bricks"
)

type house struct {
	material string
}

func newHouse(material string) house {
	return house{material: material}
}
