package main

import "math/rand"

type Necromancer struct {
	*Character
	PreviousLocation []int
	Freeze           bool
}

var necromancer *Necromancer

func (necromancer *Necromancer) shouldFreeze(str string) {
	necromancer.Freeze = InitialsIndexOf([]string{commands.Map}, str)
}

func CreateNecromancer() {
	HP := rand.Intn(50) + 80
	necromancerStartPosition := getAvailableRoom(2, X-1)
	necromancer = &Necromancer{
		Character: &Character{
			Name:            enemiesList.NECROMANCER,
			Npc:             true,
			Alive:           true,
			CurrentLocation: necromancerStartPosition,
			Evasion:         20,
			Health:          HP,
			BaseHealth:      HP,
			Skill:           5,
			Strength:        18,
			Crit:            35,
			ExpValue:        60,
			Inventory:       map[string]*ItemQuantity{},
			LevelUp: &Leveling{
				NextRank:            5,
				NextBase:            5,
				Exp:                 rand.Intn(35) + rand.Intn(35),
				achievedLevelsChain: []int{},
				Rates: &Specifics{
					Health: func() int {
						return rand.Intn(10) + 1
					},
					Crit:     3,
					Evasion:  1,
					Skill:    1,
					Strength: 3,
				},
			},
		},
		PreviousLocation: necromancerStartPosition,
		Freeze:           false,
	}
	// var loc *Location
	necromancer.createEnemyInventory()
	necromancer.setImage()
	pile.PushCharacters(necromancer.Character)
	necromancer.SetPlayerRoom()
	// loc = necromancer.SetPlayerRoom()
	// loc.AddEphemeral()
}

func (necromancer *Necromancer) necromancerMoves() {
	// dragonLanding.reset()

	if !necromancer.Freeze {
		if ok := necromancer.isAlive(); ok {
			loc := necromancer.SetPlayerRoom()

			possibleWays := necromancerPossibleWays()
			if len(possibleWays) > 0 {
				goTo := possibleWays[rand.Intn(len(possibleWays))]
				loc.ClearEphemeral()
				necromancer.MoveTo(goTo)
			}
			necromancer.PreviousLocation = []int{loc.Y, loc.X}
			newLoc := necromancer.SetPlayerRoom()
			newLoc.AddEphemeral()
		}
	}
}

func necromancerPossibleWays() []string {
	var youCanGo []string
	var yourPlace *Location
	yourPlace = necromancer.SetPlayerRoom()

	if yourPlace.X >= 1 {
		if canNecromancerMoveThatWay(yourPlace.Y, yourPlace.X-1) {
			youCanGo = append(youCanGo, directions.West)
		}
	}
	if yourPlace.X <= 8 {
		if canNecromancerMoveThatWay(yourPlace.Y, yourPlace.X+1) {
			youCanGo = append(youCanGo, directions.East)
		}
	}
	if yourPlace.Y >= 1 {
		if canNecromancerMoveThatWay(yourPlace.Y-1, yourPlace.X) {
			youCanGo = append(youCanGo, directions.North)
		}
	}
	if yourPlace.Y <= 2 {
		if canNecromancerMoveThatWay(yourPlace.Y+1, yourPlace.X) {
			youCanGo = append(youCanGo, directions.South)
		}
	}
	return youCanGo
}

func canNecromancerMoveThatWay(y, x int) bool {
	var loc *Location = WorldMap[y][x]
	// ICI
	if (necromancer.PreviousLocation[0] == y) && (necromancer.PreviousLocation[1] == x) {
		return false
	}
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

func (necromancer *Necromancer) deadScan() {

}
