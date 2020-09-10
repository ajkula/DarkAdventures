package main

import "math/rand"

var necromancer *Walker

func CreateNecromancer() {
	HP := rand.Intn(50) + 80
	necromancerStartPosition := getAvailableRoom(2, X-1)
	necromancer = &Walker{
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
					Evasion:  0,
					Skill:    1,
					Strength: 3,
				},
			},
			StatusEffects: &StatusEffectsBlueprint{
				AllStatus: []*Blueprint{},
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
