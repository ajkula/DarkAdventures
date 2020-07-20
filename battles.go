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

	// for {
	if !player.isAlive() || !enemy.isAlive() {
		// break
		return
	}

	actualPlayer := opponents[turns()]
	if !actualPlayer.Npc {
		// Output("white", "\tYou are fighting "+Article(enemy.Name)+"\n")
		// Output("white", Tab+CalculateSpaceAlign(player.Name+": ")+strconv.Itoa(player.Health)+" HP")
		// Output("yellow", Tab+CalculateSpaceAlign(enemy.Name+": ")+strconv.Itoa(enemy.Health)+" HP")
		// Output("white", Tab+CalculateSpaceAlign("You can: ")+ArrayToString(battleCommands))
		// Output("white", Tab+CalculateSpaceAlign("You can use: ")+actualPlayer.DisplayItems())
		// Output("white", Tab+CalculateSpaceAlign("...or: ")+ArrayToString(universalCommands))
		cmd := UserInputln()
		if ok := arrayIncludesCommand(battleCommands, cmd); ok {
			ProcessCommands(player, cmd, enemy)
		} else {
			Output("red", Tab+"You can't do that here...")
			ResetTurns()
		}

	} else {
		Output("red", actualPlayer.Name)
		EnemyAction(player, enemy)
	}
	// }

	for _, opp := range opponents {
		if !opp.isAlive() {
			ResetTurns()
			switch opp.Npc {
			case true:
				Output("green", "\t", opp.Name, " has been slain")
				// you got :
				player.getEnemyItems(enemy)
				return
			case false:
				if !opp.hasItemInInventory(itemNames.Doll) || (opp.hasItemInInventory(itemNames.Doll) && opp.Inventory[itemNames.Doll].Quantity < 1) {
					Output("red", "\tYou died\n")
					Output("red", "\tGAME OVER\n")
					os.Exit(0)
				}
				if opp.hasItemInInventory(itemNames.Doll) && opp.Inventory[itemNames.Doll].Quantity >= 1 {
					opp.useItem(itemNames.Doll)
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
