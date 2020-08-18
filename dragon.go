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

type DragonLanding struct {
	isFirst bool
}

var dragonLanding = &DragonLanding{
	isFirst: true,
}

func (landing *DragonLanding) shouldSayIt() bool {
	return landing.isFirst
}
func (landing *DragonLanding) saidIt() {
	landing.isFirst = false
}
func (landing *DragonLanding) reset() {
	landing.isFirst = true
}

func (dragon *Dragon) dragonMoves() {
	dragonLanding.reset()
	// ICI
	// fmt.Printf("DRAGON SHOULD FREEZE: %+v\n", dragon.Freeze)
	if !dragon.Freeze {
		if ok := dragon.isAlive(); ok {
			loc := dragon.SetPlayerRoom()
			dragon.PreviousLocation = []int{loc.Y, loc.X}

			possibleWays := dragonPossibleWays()
			// Output("red", "dragon can go: ", possibleWays)
			if len(possibleWays) > 0 {
				goTo := possibleWays[rand.Intn(len(possibleWays))]
				// Output("red", "dragon moves to: ", goTo)
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

// TO DO: add player coins on presentation

func (dragon *Dragon) shouldFreeze(str string) {
	dragon.Freeze = InitialsIndexOf([]string{commands.Map}, str)
}

func CreateDragon() {
	HP := rand.Intn(50) + 80
	dragonStartPosition := getAvailableRoom(4, X-1)
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
			ExpValue:        50,
			Inventory:       map[string]*ItemQuantity{},
			LevelUp: &Leveling{
				NextRank:            5,
				NextBase:            5,
				Exp:                 rand.Intn(25) + rand.Intn(25),
				achievedLevelsChain: []int{},
				Rates: &Specifics{
					Health: func() int {
						return rand.Intn(10) + 1
					},
					Crit:     2,
					Evasion:  1,
					Skill:    2,
					Strength: 2,
				},
			},
		},
		PreviousLocation: dragonStartPosition,
		Freeze:           false,
	}
	var loc *Location
	dragon.createEnemyInventory()
	dragon.setImage()
	pile.PushCharacters(dragon.Character)
	loc = dragon.SetPlayerRoom()
	loc.AddEphemeral()
}

func dragonPossibleWays() []string {
	var youCanGo []string
	var yourPlace *Location
	yourPlace = dragon.SetPlayerRoom()

	if yourPlace.X >= 1 {
		if canDragonMoveThatWay(yourPlace.Y, yourPlace.X-1) {
			youCanGo = append(youCanGo, directions.West)
		}
	}
	if yourPlace.X <= 8 {
		if canDragonMoveThatWay(yourPlace.Y, yourPlace.X+1) {
			youCanGo = append(youCanGo, directions.East)
		}
	}

	if yourPlace.Y >= 1 {
		if canDragonMoveThatWay(yourPlace.Y-1, yourPlace.X) {
			youCanGo = append(youCanGo, directions.North)
		}
	}
	if yourPlace.Y <= 2 {
		if canDragonMoveThatWay(yourPlace.Y+1, yourPlace.X) {
			youCanGo = append(youCanGo, directions.South)
		}
	}
	return youCanGo
}

func canDragonMoveThatWay(y, x int) bool {
	var loc *Location = WorldMap[y][x]
	// ICI

	if loc.HasEnemy {
		// fmt.Println("can move that way: false", " enemy: ", loc.HasEnemy, " shop: ", loc.HasSeller)
		return false
	}
	if loc.HasSeller {
		// fmt.Println("can move that way: false", " enemy: ", loc.HasEnemy, " shop: ", loc.HasSeller)
		return false
	}
	// fmt.Println("can move that way: true", " enemy: ", loc.HasEnemy, " shop: ", loc.HasSeller)
	return true
}

func getAvailableRoom(yi, xi int) []int {
	dragonStartPosition := []int{rand.Intn(yi), rand.Intn(xi)}
	y := dragonStartPosition[0]
	x := dragonStartPosition[1]
	if !(WorldMap[y][x].HasEnemy || WorldMap[y][x].HasSeller) {
		return dragonStartPosition
	}
	return getAvailableRoom(yi, xi)
}
