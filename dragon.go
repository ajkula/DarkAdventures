package main

import "math/rand"

type Dragon struct {
	*Character
	PreviousLocation []int
}

var dragon *Dragon

func (dragon *Dragon) dragonMoves() {
	loc := dragon.SetPlayerRoom()
	dragon.PreviousLocation[0] = loc.Y
	dragon.PreviousLocation[1] = loc.X

	possibleWays := dragonPossibleWays()
	if len(possibleWays) > 0 {
		dragon.MoveTo(possibleWays[rand.Intn(len(possibleWays))])
	}
	loc = dragon.SetPlayerRoom()
	loc.AddEphemeral()
}

// Character{
// 	Health:     HP,
// 	BaseHealth: HP,
// 	Strength:   base.Strength,
// 	Evasion:    base.Evasion,
// 	Crit:       base.Crit,
// 	Name:       name,
// 	Alive:      true,
// 	Npc:        true,
// 	Inventory:  map[string]*ItemQuantity{},
// }

func CreateDragon() {
	HP := rand.Intn(20) + 80
	dragonStartPosition := getAvailableRoom()
	dragon = &Dragon{
		Character: &Character{
			Name:            enemiesList.DRAGON,
			Npc:             true,
			Alive:           true,
			CurrentLocation: dragonStartPosition,
			Evasion:         30,
			Health:          HP,
			BaseHealth:      HP,
			Skill:           3,
			Strength:        20,
			Crit:            30,
			Inventory:       map[string]*ItemQuantity{},
		},
		PreviousLocation: dragonStartPosition,
	}
	var loc *Location
	loc = dragon.SetPlayerRoom()
	loc.AddEphemeral()
}

func dragonPossibleWays() []string {
	var youCanGo []string
	var yourPlace *Location
	yourPlace = dragon.SetPlayerRoom()
	yourPlace.ClearEphemeral()

	if yourPlace.X >= 1 && canDragonMoveThatWay(yourPlace.X-1, yourPlace.Y) {
		youCanGo = append(youCanGo, directions.West)
	}
	if yourPlace.X <= 8 && canDragonMoveThatWay(yourPlace.X+1, yourPlace.Y) {
		youCanGo = append(youCanGo, directions.East)
	}

	if yourPlace.Y >= 1 && canDragonMoveThatWay(yourPlace.X, yourPlace.Y-1) {
		youCanGo = append(youCanGo, directions.North)
	}
	if yourPlace.Y <= 2 && canDragonMoveThatWay(yourPlace.X, yourPlace.Y+1) {
		youCanGo = append(youCanGo, directions.South)
	}
	return youCanGo
}

func canDragonMoveThatWay(y, x int) bool {
	return !(WorldMap[y][x].HasEnemy || WorldMap[y][x].HasSeller) && !(x == dragon.PreviousLocation[1] && y == dragon.PreviousLocation[0])
}

func getAvailableRoom() []int {
	dragonStartPosition := []int{rand.Intn(4), rand.Intn(9)}
	y := dragonStartPosition[0]
	x := dragonStartPosition[1]
	if !(WorldMap[y][x].HasEnemy || WorldMap[y][x].HasSeller) {
		return dragonStartPosition
	}
	return getAvailableRoom()
}
