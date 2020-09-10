package main

import (
	"fmt"
	"math/rand"
)

var nightWalkerA *Walker
var nightWalkerB *Walker

func CreateNightWalker() *Walker {
	var nightWalker *Walker
	HP := rand.Intn(80) + 80
	startPosition := getAvailableRoom(8, X-1)
	nightWalker = &Walker{
		Character: &Character{
			Name:            enemiesList.NIGHTWALKER,
			Npc:             true,
			Alive:           true,
			Icon:            nightWalkerPos,
			CurrentLocation: startPosition,
			Evasion:         15,
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
				Exp:                 rand.Intn(50) + rand.Intn(25),
				achievedLevelsChain: []int{},
				Rates: &Specifics{
					Health: func() int {
						return rand.Intn(10) + 2
					},
					Crit:     2,
					Evasion:  0,
					Skill:    1,
					Strength: 2,
				},
			},
			StatusEffects: &StatusEffectsBlueprint{
				AllStatus: []*Blueprint{},
			},
		},
		PreviousLocation: startPosition,
		Freeze:           false,
	}
	var loc *Location
	nightWalker.createEnemyInventory()
	nightWalker.setImage()
	nightWalker.Encounter = &Encounter{isFirst: true}
	pile.PushCharacters(nightWalker.Character)
	loc = nightWalker.SetPlayerRoom()
	loc.AddEphemeral(nightWalker)
	fmt.Printf("\n%+v\n", nightWalker)
	return nightWalker
}
