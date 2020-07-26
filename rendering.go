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
	Output("yellow", loc.Ephemeral)
	Output("red", loc.HasEnemy)
	Output("green", getTurns())
	Output("red", "dragon.Freeze "+strconv.FormatBool(dragon.Freeze))

	if loc.HasEnemy {
		if loc.Enemy.Alive {
			Output("red", DoubleTab+"There is "+Article(loc.Enemy.Name)+"ready to fight you!\n")
			p.showHP()
			loc.Enemy.showHP()

			showActions(p, battleCommands)
			showUniversalCmds()
			Battle(p, &loc.Enemy)
		}
	}

	if !loc.HasEnemy && !loc.HasSeller {
		showActions(p, worldCommands)
		showWhereCanGo(loc)
		showUniversalCmds()
		// move
		cmd := UserInputln()
		if ok := arrayIncludesCommand(worldCommands, cmd); ok {
			ProcessCommands(p, cmd)
		} else {
			Output("red", Tab+"You can't do that here...")
		}
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
		if ok := arrayIncludesCommand(sellerCommands, cmd); ok {
			ProcessCommands(p, cmd)
		} else {
			Output("red", Tab+"You can't do that here...")
		}
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

func extractCommand(input string) string {
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return ""
	}
	return strings.ToLower(tokens[0])
}

func arrayIncludesCommand(comArr []string, input string) bool {
	if str := extractCommand(input); str != "" {
		test := append(comArr[:], universalCommands[:]...)
		if ok := InitialsIndexOf(test, str); ok {
			return true
		}
	}
	return false
}
