package main

import (
	"strconv"
	"strings"
)

func PresentScene(p *Character) {
	loc := p.SetPlayerRoom()
	loc.RemoveBattle()

	// ICI
	Output("yellow", loc.Description)
	Output("red", loc.HasEnemy)
	Output("red", p.CurrentLocation)
	Output("green", getTurns())
	if loc.HasEnemy {
		if loc.Enemy.Alive {
			Output("red", DoubleTab+"There is "+Article(loc.Enemy.Name)+"ready to fight you!\n")
			p.showHP()
			loc.Enemy.showHP()

			// if getTurns() == 0 {
			showActions(p, battleCommands)
			showUniversalCmds()
			// }
			// loc.Description = "\tYou are fighting " + Article(loc.Enemy.Name) + "\n"
			Battle(p, &loc.Enemy)

		}
	}

	if !loc.HasEnemy && !loc.HasSeller {
		showActions(p, worldCommands)
		showWhereCanGo(loc)
		showUniversalCmds()
		// move
		cmd := UserInputln()
		ProcessCommands(p, cmd)
	}

	if loc.HasSeller {
		Output("yellow", loc.Seller)

		var concat string = ""
		for name, element := range loc.Item {
			concat += DoubleTab + strconv.Itoa(element.Quantity) + " - " + name + " for " + strconv.Itoa(element.Type.Price) + " coins." + "\n"
		}
		Output("yellow", Tab+"He's proposing:\n"+concat)
		showActions(p, sellerCommands)
		showWhereCanGo(loc)
		showUniversalCmds()
		// buy listener
		cmd := UserInputln()
		ProcessCommands(p, cmd)
	}

}

func showWhereCanGo(loc *Location) {
	Output("blue", Tab+CalculateSpaceAlign("You can Go:"), WhereCanYouGo(loc))
}

func showActions(p *Character, arr []string) {
	Output("white", Tab+CalculateSpaceAlign("You can: "), arr)
	Output("white", Tab+CalculateSpaceAlign("You can use: "), p.DisplayItems())
}

func showUniversalCmds() {
	Output("white", Tab+CalculateSpaceAlign("...or: ")+ArrayToString(universalCommands), "\n")
}

func CalculateSpaceAlign(str string) string {
	var length int = 1
	if Tabulation-len(str) > 1 {
		length = Tabulation - len(str)
	}
	spaces := strings.Repeat(" ", length)
	return str + spaces
}

func CustomSpaceAlign(str string, i int) string {
	var length int = 1
	if i-len(str) > 1 {
		length = i - len(str)
	}
	spaces := strings.Repeat(" ", length)
	return str + spaces
}

func Article(str string) string {
	vowels := []string{"a", "e", "i", "o"}
	if InitialsIndexOf(vowels, str) {
		return "an " + str + " "
	}
	return "a " + str + " "
}
