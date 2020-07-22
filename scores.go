package main

import "strconv"

type ScoreSchema struct {
	Enemies  map[string]*Calculus
	Items    map[string]*Calculus
	Dammages *DammageSchema
	Total    int
}

type DammageSchema struct {
	Dealt int
	Taken int
}

type Calculus struct {
	Quantity int
	Points   int
}

var SCORE *ScoreSchema = &ScoreSchema{
	Items: map[string]*Calculus{
		itemNames.Doll:      {Points: 50},
		itemNames.Moonstone: {Points: 45},
		itemNames.Scroll:    {Points: 30},
		itemNames.Potion:    {Points: 20},
		itemNames.Key:       {Points: 15},
		itemNames.Coins:     {Points: 10},
	},
	Enemies: map[string]*Calculus{
		enemiesList.GOBLIN:   {Points: 25},
		enemiesList.SKELETON: {Points: 30},
		enemiesList.SORCERER: {Points: 80},
		enemiesList.ORC:      {Points: 100},
		enemiesList.DRAGON:   {Points: 250},
	},
	Dammages: &DammageSchema{
		Dealt: 0,
		Taken: 0,
	},
	Total: 0,
}

func (score *ScoreSchema) scoreKills(name string) {
	score.Enemies[name].Quantity++
}

func (score *ScoreSchema) scoreItems(name string, amount int) {
	score.Items[name].Quantity += amount
}

func (score *ScoreSchema) scoreDammages(b bool, amount int) {
	if !b {
		score.Dammages.Dealt += amount
	} else {
		score.Dammages.Taken += amount
	}
}

func (score *ScoreSchema) getSCORE() {
	Output("white", DoubleTab+"SCORE")
	Output("white", DoubleTab+"Kills:")
	for _, name := range indexedEnemiesForRandomization {
		if SCORE.Enemies[name].Quantity > 0 {
			SCORE.Total += SCORE.Enemies[name].Points * SCORE.Enemies[name].Quantity
			Output("white", Tab+CalculateSpaceAlign(name+":")+strconv.Itoa(SCORE.Enemies[name].Quantity))
		}
	}
	Output("white", DoubleTab+"Loots:")
	for _, name := range ItemIndexList {
		if SCORE.Items[name].Quantity > 0 {
			SCORE.Total += SCORE.Items[name].Points * SCORE.Items[name].Quantity
			Output("white", Tab+CalculateSpaceAlign(name+":")+strconv.Itoa(SCORE.Items[name].Quantity))
		}
	}
	Output("white", Tab+CalculateSpaceAlign("DMG Dealt:")+strconv.Itoa(SCORE.Dammages.Dealt))
	Output("white", Tab+CalculateSpaceAlign("DMG Taken:")+strconv.Itoa(SCORE.Dammages.Taken))

	Output("white", Tab+CalculateSpaceAlign("TOTAL SCORE:")+strconv.Itoa(SCORE.Total))

} // SCORE.Enemies[name].Points*SCORE.Enemies[name].Quantity)
