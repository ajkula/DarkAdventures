package main

import (
	"math/rand"
)

type Dragon struct {
	*Character
	PreviousLocation []int
	Freeze           bool
}

var dragon *Dragon

func (dragon *Dragon) dragonMoves() {
	// ICI
	// fmt.Printf("DRAGON SHOULD FREEZE: %+v\n", dragon.Freeze)
	if !dragon.Freeze {
		if ok := dragon.isAlive(); ok {
			loc := dragon.SetPlayerRoom()
			dragon.PreviousLocation = []int{loc.Y, loc.X}

			possibleWays := dragonPossibleWays()
			Output("red", "dragon can go: ", possibleWays)
			if len(possibleWays) > 0 {
				goTo := possibleWays[rand.Intn(len(possibleWays))]
				Output("red", "dragon moves to: ", goTo)
				loc.ClearEphemeral()
				dragon.MoveTo(goTo)
			}
			newLoc := dragon.SetPlayerRoom()
			// fmt.Printf("DRAGON: %+v\ndragon character: %+v\n", dragon, dragon.Character)
			// fmt.Printf("NewLOC: %+v\n", newLoc)
			// fmt.Printf("OLD LOC: Y: %+v X: %+v\n", loc.Y, loc.X)
			newLoc.AddEphemeral()
		}
	}
}

func (dragon *Dragon) shouldFreeze(str string) {
	dragon.Freeze = str == Initial(commands.Map)
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
	HP := rand.Intn(50) + 80
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
		Freeze:           false,
	}
	var loc *Location
	dragon.createEnemyInventory()
	loc = dragon.SetPlayerRoom()
	loc.AddEphemeral()
}

func dragonPossibleWays() []string {
	var youCanGo []string
	var yourPlace *Location
	yourPlace = dragon.SetPlayerRoom()

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
