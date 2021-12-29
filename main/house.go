package main

const (
	STRAW  = "straw"
	STICKS = "sticks"
	BRICKS = "bricks"
)

type House struct {
	Material string
}

func NewHouse(material string) House {
	return House{Material: material}
}
