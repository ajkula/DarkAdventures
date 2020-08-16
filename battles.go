package main

import (
	"os"
)

var boolean bool

func boolToInt8(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

func ResetTurns() {
	boolean = true
}

func getTurns() int8 {
	return boolToInt8(boolean)
}

// turns func => int8 returns 0 / 1 alternatively each time it's invoked
func turns() int8 {
	boolean = !boolean
	return boolToInt8(boolean)
}

func PrintArrayAligned(arr []string) string {
	return Tab + CalculateSpaceAlign(ArrayToString(arr))
}

func Battle(player, enemy *Character) {
	opponents := []*Character{player, enemy}
	if !player.isAlive() || !enemy.isAlive() {
		return
	}
	actualPlayer := opponents[turns()]
	if !actualPlayer.Npc {
		cmd := UserInputln()
		if ok := arrayIncludesCommand(battleCommands, cmd); ok {
			ProcessCommands(player, cmd, enemy)
		} else {
			Output("red", translate(cantDoThatTR))
			ResetTurns()
		}

	} else {
		Output("red", actualPlayer.Name)
		EnemyAction(player, enemy)
	}

	for _, opp := range opponents {
		if !opp.isAlive() {
			ResetTurns()
			switch opp.Npc {
			case true:
				Output("green", "\t", opp.Name, translate(hasBeenSlainTR))
				// you got :
				SCORE.scoreKills(opp.Name)
				EnemiesKilled++
				player.LevelUp.Exp += opp.ExpValue
				player.getEnemyItems(enemy)
				player.calcLVL()
				return
			case false:
				if !opp.hasItemInInventory(itemNames.Doll) || (opp.hasItemInInventory(itemNames.Doll) && opp.Inventory[itemNames.Doll].Quantity < 1) {
					Output("red", translate(youDiedTR))
					Output("red", GameOverAscii)
					SCORE.getSCORE()
					os.Exit(0)
				}
				if opp.hasItemInInventory(itemNames.Doll) && opp.Inventory[itemNames.Doll].Quantity >= 1 {
					enemy.LevelUp.Exp += player.ExpValue
					opp.useItem(itemNames.Doll)
					enemy.calcLVL()
				}
			}
		}
	}
}

// PercentChances(50) bool

func EnemyAction(p, e *Character) {
	if !e.isAlive() {
		return
	}

	MakeEnemyDecision(p, e)
}
