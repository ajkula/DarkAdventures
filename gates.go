package main

import (
	"math/rand"
	"time"
)

type Gate struct {
	X      int
	Y      int
	Index  int
	TwoWay []*Gate
}

type Coords struct {
	Xa int
	Xb int
	Ya int
	Yb int
}

type SubMaps struct {
	NorthWest *Coords
	NorthEast *Coords
	SouthWest *Coords
	SouthEast *Coords
}

type SubMapNames struct {
	NorthWest string
	NorthEast string
	SouthWest string
	SouthEast string
}

var gateA *Gate = &Gate{}
var gateB *Gate = &Gate{}
var gatesTwoWayArray []*Gate

var subMapNames SubMapNames = SubMapNames{
	NorthEast: "NorthEast",
	NorthWest: "NorthWest",
	SouthEast: "SouthEast",
	SouthWest: "SouthWest",
}
var subMapsParts []string = []string{subMapNames.NorthEast, subMapNames.NorthWest, subMapNames.SouthEast, subMapNames.SouthWest}

var subMaps *SubMaps
var diagonals map[int][]string = map[int][]string{
	0: {subMapNames.NorthEast, subMapNames.SouthWest},
	1: {subMapNames.NorthWest, subMapNames.SouthEast},
}

func getDiagonal() []string {
	return diagonals[rand.Intn(2)]
}

func getRandomSubMap(part string) (x, y int) {
	switch part {
	case subMapNames.NorthWest:
		x = rand.Intn(subMaps.NorthWest.Xa) + subMaps.NorthWest.Xb
		y = rand.Intn(subMaps.NorthWest.Ya) + subMaps.NorthWest.Yb
	case subMapNames.NorthEast:
		x = rand.Intn(subMaps.NorthEast.Xa) + subMaps.NorthEast.Xb
		y = rand.Intn(subMaps.NorthEast.Ya) + subMaps.NorthEast.Yb
	case subMapNames.SouthWest:
		x = rand.Intn(subMaps.SouthWest.Xa) + subMaps.SouthWest.Xb
		y = rand.Intn(subMaps.SouthWest.Ya) + subMaps.SouthWest.Yb
	case subMapNames.SouthEast:
		x = rand.Intn(subMaps.SouthEast.Xa) + subMaps.SouthEast.Xb
		y = rand.Intn(subMaps.SouthEast.Ya) + subMaps.SouthEast.Yb
	}
	return x, y
}

func InitGates() {
	subMaps = &SubMaps{
		NorthWest: &Coords{
			Xa: X / 2,
			Xb: 0,
			Ya: Y / 2,
			Yb: 0,
		},
		NorthEast: &Coords{
			Xa: X / 2,
			Xb: X / 2,
			Ya: Y / 2,
			Yb: 0,
		},
		SouthWest: &Coords{
			Xa: X / 2,
			Xb: 0,
			Ya: (Y - 2) / 2,
			Yb: Y / 2,
		},
		SouthEast: &Coords{
			Xa: X / 2,
			Xb: X / 2,
			Ya: (Y - 2) / 2,
			Yb: Y / 2,
		},
	}

	pile.PushGates(gateA)
	pile.PushGates(gateB)
	diag := getDiagonal()
	for i, coord := range diag {
		x, y := getRandomSubMap(coord)
		pile.Gates[i].X = x
		pile.Gates[i].Y = y
		pile.Gates[i].Index = i
		pile.Gates[i].TwoWay = pile.Gates
		WorldMap[y][x].HasGate = true
		WorldMap[y][x].Gate = pile.Gates[i]
	}

	// fmt.Printf("gateA: %+v\n", gateA)
	// fmt.Printf("gateB: %+v\n", gateB)
}

func (gate *Gate) Warp(player *Character) {
	player.CurrentLocation = []int{gate.TwoWay[ToggleIndexs(gate.Index)].Y, gate.TwoWay[ToggleIndexs(gate.Index)].X}
	Output("green", warpText)
	time.Sleep(2 * time.Second)
	player.SetPlayerRoom()
}

func ToggleIndexs(i int) (index int) {
	index = 0
	if i == 0 {
		index = 1
	}
	return index
}
