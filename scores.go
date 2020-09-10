package main

import (
	"strconv"
)

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
		enemiesList.GOBLIN:      {Points: 25},
		enemiesList.SKELETON:    {Points: 30},
		enemiesList.SORCERER:    {Points: 80},
		enemiesList.ORC:         {Points: 100},
		enemiesList.DRAGON:      {Points: 250},
		enemiesList.NIGHTWALKER: {Points: 250},
		enemiesList.NECROMANCER: {Points: 150},
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
	kills := ""
	loot := ""
	Output("white", DoubleTab+"SCORE")
	for name, obj := range SCORE.Enemies {
		if obj.Quantity > 0 {
			SCORE.Total += obj.Points * obj.Quantity
			enemy := strconv.Itoa(obj.Quantity)
			kills += Tab + CustomSpaceAlign(name+":", 22-len(enemy)) + enemy + "\n"
		}
	}
	if kills != "" {
		Output("white", DoubleTab+"Kills:")
		Output("white", kills)
	}
	for _, name := range ItemIndexList {
		if SCORE.Items[name].Quantity > 0 {
			SCORE.Total += SCORE.Items[name].Points * SCORE.Items[name].Quantity
			item := strconv.Itoa(SCORE.Items[name].Quantity)
			loot += Tab + CustomSpaceAlign(name+":", 22-len(item)) + item + "\n"
		}
	}
	if loot != "" {
		Output("white", DoubleTab+"LOOT:")
		Output("white", loot)
	}
	if SCORE.Dammages.Dealt > 0 || SCORE.Dammages.Taken > 0 {
		Output("white", DoubleTab+"POINTS:")
	}
	dealt := strconv.Itoa(SCORE.Dammages.Dealt)
	taken := strconv.Itoa(SCORE.Dammages.Taken)
	Output("white", Tab+CustomSpaceAlign("DMG Dealt:", 22-len(dealt))+dealt)
	Output("white", Tab+CustomSpaceAlign("DMG Taken:", 22-len(taken))+taken)

	total := strconv.Itoa(SCORE.Total)
	Output("white", Tab+CustomSpaceAlign("TOTAL SCORE:", 22-len(total))+total)
}

func (score *ScoreSchema) removeBaseInventory(inv map[string]*ItemQuantity) {
	for name, itemQ := range inv {
		SCORE.Items[name].Quantity -= itemQ.Quantity
	}
}
