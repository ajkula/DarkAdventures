package main

import (
	"math/rand"
)

var dragon *Walker

func CreateDragon() {
	HP := rand.Intn(50) + 100
	dragonStartPosition := getAvailableRoom(4, X-1)
	dragon = &Walker{
		Character: &Character{
			Name:            enemiesList.DRAGON,
			Npc:             true,
			Alive:           true,
			Icon:            dragonPos,
			CurrentLocation: dragonStartPosition,
			Evasion:         20,
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
					Evasion:  0,
					Skill:    2,
					Strength: 2,
				},
			},
			StatusEffects: &StatusEffectsBlueprint{
				AllStatus: []*Blueprint{},
			},
		},
		PreviousLocation: dragonStartPosition,
		Freeze:           false,
	}
	var loc *Location
	dragon.createEnemyInventory()
	dragon.setImage()
	dragon.Encounter = &Encounter{isFirst: true}
	pile.PushCharacters(dragon.Character)
	loc = dragon.SetPlayerRoom()
	loc.AddEphemeral(dragon)
}
